package main

// Transition represents transition between states specified with number.
type Transition struct {
	from int
	to   int
}

// State represents a state labeled with formulas.
type State struct {
	labels []Formula
}

// TransitionSystem represents a transition system.
type TransitionSystem struct {
	N           int
	transitions []Transition
	states      map[int]*State
}
