package main

import "fmt"

// MakeFormulaStack returns the stack of formula so that
// the smallest subformula becomes top.
func MakeFormulaStack(f Formula) []Formula {
	var queue []Formula
	visited := make(map[Formula]bool)
	var stack []Formula

	queue = append(queue, f)
	visited[f] = true
	stack = append(stack, f)

	for len(queue) > 0 {
		// Dequeue.
		v := queue[0]
		if len(queue) == 1 {
			queue = []Formula{}
		} else {
			queue = queue[1:]
		}

		// Store nodes to stack with BFS.
		switch v := v.(type) {
		case ConjFormula:
			if _, ok := visited[v.f1]; !ok {
				visited[v.f1] = true
				queue = append(queue, v.f1)
				stack = append(stack, v.f1)
			}
			if _, ok := visited[v.f2]; !ok {
				visited[v.f2] = true
				queue = append(queue, v.f2)
				stack = append(stack, v.f2)
			}
		case NegFormula:
			if _, ok := visited[v.f]; !ok {
				visited[v.f] = true
				queue = append(queue, v.f)
				stack = append(stack, v.f)
			}
		case AFFormula:
			if _, ok := visited[v.f1]; !ok {
				visited[v.f1] = true
				queue = append(queue, v.f1)
				stack = append(stack, v.f1)
			}
		case EUFormula:
			if _, ok := visited[v.f1]; !ok {
				visited[v.f1] = true
				queue = append(queue, v.f1)
				stack = append(stack, v.f1)
			}
			if _, ok := visited[v.f2]; !ok {
				visited[v.f2] = true
				queue = append(queue, v.f2)
				stack = append(stack, v.f2)
			}
		case EXFormula:
			if _, ok := visited[v.f1]; !ok {
				visited[v.f1] = true
				queue = append(queue, v.f1)
				stack = append(stack, v.f1)
			}
		}
	}
	return stack
}

// LabelContains returns true if given labels contains given formula.
func LabelContains(labels []Formula, f Formula) bool {
	for _, l := range labels {
		if l.Equal(f) {
			return true
		}
	}
	return false
}

// Next returns next states of given number as int array.
func (ts TransitionSystem) Next(from int) []int {
	var nextStates []int
	for _, t := range ts.transitions {
		if t.from == from {
			nextStates = append(nextStates, t.to)
		}
	}
	return nextStates
}

// Labeling returns states that satisfy given formula checked with the labeling algorithm.
func Labeling(ts TransitionSystem, f Formula) map[int]*State {
	// Label all the states with true.
	for _, s := range ts.states {
		s.labels = append(s.labels, TrueFormula{})
	}

	// Stores visited nodes with BFS as stack.
	stack := MakeFormulaStack(f)
	for {
		// Do until stack becomes empty.
		if len(stack) == 0 {
			break
		}

		// Pop node from stack.
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch cur := cur.(type) {
		case ConjFormula: // if f contains Conjformula c then
			// For each state, if it is labeled with both c.f1 and c.f2,
			// then label it with c.
			for _, s := range ts.states {
				if LabelContains(s.labels, cur.f1) && LabelContains(s.labels, cur.f2) {
					s.labels = append(s.labels, cur)
				}
			}
		case NegFormula: // if f contains NegFormula n then
			// For each state, if it is NOT labeled with n.f
			// then label it with n.
			for _, s := range ts.states {
				if !LabelContains(s.labels, cur.f) {
					s.labels = append(s.labels, cur)
				}
			}
		case AFFormula: // if f contains AFFormula af then
			// For each state, if it is labeled with af.f1
			// then label it with af.
			for _, s := range ts.states {
				if LabelContains(s.labels, cur.f1) {
					s.labels = append(s.labels, cur)
				}
			}
			// Repeat: for each state, if all the next states are labeled with af
			// then label it with af, too.
			for {
				unchanged := true
				for i := range ts.states {
					if LabelContains(ts.states[i].labels, cur) {
						continue
					}
					for _, j := range ts.Next(i) {
						if !LabelContains(ts.states[j].labels, cur) {
							break
						}
						ts.states[i].labels = append(ts.states[i].labels, cur)
						unchanged = false
					}
				}
				if unchanged {
					break
				}
			}
		case EUFormula: // if f contains EUFormula eu then
			// For each state, if it is labeled with eu.f2
			// then label it with eu.
			for _, s := range ts.states {
				if LabelContains(s.labels, cur.f2) {
					s.labels = append(s.labels, cur)
				}
			}
			// Repeat: for each state, if it is labeled with eu.f1 and
			// at least one of the next state is labeled with eu
			// then label it with eu, too.
			for {
				unchanged := true
				for i := range ts.states {
					if LabelContains(ts.states[i].labels, cur) {
						continue
					}
					for _, j := range ts.Next(i) {
						if LabelContains(ts.states[i].labels, cur.f1) && LabelContains(ts.states[j].labels, cur) {
							ts.states[i].labels = append(ts.states[i].labels, cur)
							unchanged = false
							break
						}
					}
				}
				if unchanged {
					break
				}
			}
		case EXFormula: // if f contains EXFormula ex then
			// For each state, if one of the next state is labeled with ex.f1
			// then label it with ex.
			for i := range ts.states {
				for _, j := range ts.Next(i) {
					if LabelContains(ts.states[j].labels, cur.f1) {
						ts.states[i].labels = append(ts.states[i].labels, cur)
						break
					}
				}
			}
		}
	}

	// Return the set of states that satisfy f.
	satisfyStates := make(map[int]*State)
	for i, s := range ts.states {
		if LabelContains(s.labels, f) {
			satisfyStates[i] = s
		}
	}

	return satisfyStates
}

func main() {
	// Below represents a formula E[¬c2Uc1] and a transition system.
	// Its answer is [0, 1, 2, 3, 4].
	f := EUFormula{
		f1: NegFormula{
			f: AtomicFormula{
				a: "c2",
			},
		},
		f2: AtomicFormula{
			a: "c1",
		},
	}
	fmt.Print("CTL formula: ")
	f.Print()
	fmt.Println()

	var ts TransitionSystem
	ts.N = 9
	ts.transitions = []Transition{
		{0, 1}, {1, 2}, {2, 3}, {1, 4}, {4, 3}, {4, 0}, {3, 5},
		{0, 5}, {5, 6}, {6, 7}, {5, 8}, {8, 7}, {8, 0}, {7, 1},
	}
	ts.states = map[int]*State{
		0: {labels: []Formula{AP("n1"), AP("n2")}},
		1: {labels: []Formula{AP("t1"), AP("n2")}},
		2: {labels: []Formula{AP("t1"), AP("t2")}},
		3: {labels: []Formula{AP("c1"), AP("t2")}},
		4: {labels: []Formula{AP("c1"), AP("n2")}},
		5: {labels: []Formula{AP("n1"), AP("t2")}},
		6: {labels: []Formula{AP("t1"), AP("t2")}},
		7: {labels: []Formula{AP("t1"), AP("c2")}},
		8: {labels: []Formula{AP("n1"), AP("c2")}},
	}

	// Solve with the labeling algorithm.
	satisfyStates := Labeling(ts, f)

	// Output result and labels for each states.
	fmt.Println("== RESULT ==")
	for i := 0; i < ts.N; i++ {
		fmt.Print(i)
		if _, ok := satisfyStates[i]; ok {
			fmt.Print(" [SATISFY]:     ")
		} else {
			fmt.Print(" [NOT SATISFY]: ")
		}

		// Output labels.
		for _, l := range ts.states[i].labels {
			l.Print()
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
