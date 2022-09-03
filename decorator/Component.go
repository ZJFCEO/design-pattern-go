package main

type Component interface {
	Calc() int
}
type ConcreateComponent struct {
}

func (*ConcreateComponent) Calc() int {
	return 0
}
