package main

import "fmt"

// Formula represents CTL formula as an interface.
type Formula interface {
	Print()
	Equal(Formula) bool
}

type AP string

func (p AP) Print() {
	fmt.Print(p)
}

func (p AP) Equal(q Formula) bool {
	switch q := q.(type) {
	case AP:
		return p == q
	case AtomicFormula:
		return p == q.a
	default:
		return false
	}
}

type TrueFormula struct{}

func (p TrueFormula) Print() {
	fmt.Print("true")
}

func (p TrueFormula) Equal(q Formula) bool {
	switch q.(type) {
	case TrueFormula:
		return true
	default:
		return false
	}
}

type AtomicFormula struct {
	a AP
}

func (p AtomicFormula) Print() {
	fmt.Print(p.a)
}

func (p AtomicFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case AtomicFormula:
		return p.a == q.a
	case AP:
		return p.a == q
	default:
		return false
	}
}

type ConjFormula struct {
	f1 Formula
	f2 Formula
}

func (p ConjFormula) Print() {
	p.f1.Print()
	fmt.Print("∧")
	p.f2.Print()
}

func (p ConjFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case ConjFormula:
		return p.f1.Equal(q.f1) && p.f2.Equal(q.f2)
	default:
		return false
	}
}

type NegFormula struct {
	f Formula
}

func (p NegFormula) Print() {
	fmt.Print("¬")
	p.f.Print()
}

func (p NegFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case NegFormula:
		return p.f.Equal(q.f)
	default:
		return false
	}
}

type AFFormula struct {
	f1 Formula
}

func (p AFFormula) Print() {
	fmt.Print("AF")
	p.f1.Print()
}

func (p AFFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case AFFormula:
		return p.f1.Equal(q.f1)
	default:
		return false
	}
}

type EUFormula struct {
	f1 Formula
	f2 Formula
}

func (p EUFormula) Print() {
	fmt.Print("E[")
	p.f1.Print()
	fmt.Print("U")
	p.f2.Print()
	fmt.Print("]")
}

func (p EUFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case EUFormula:
		return p.f1.Equal(q.f1) && p.f2.Equal(q.f2)
	default:
		return false
	}
}

type EXFormula struct {
	f1 Formula
}

func (p EXFormula) Print() {
	fmt.Print("EX")
	p.f1.Print()
}

func (p EXFormula) Equal(q Formula) bool {
	switch q := q.(type) {
	case EXFormula:
		return p.f1.Equal(q.f1)
	default:
		return false
	}
}
