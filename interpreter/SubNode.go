package main

type SubNode struct {
	left, right Node
}

func (a *SubNode) Interpret() int {
	return a.left.Interpret() - a.right.Interpret()
}
