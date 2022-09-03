package main

type OperatorBase struct {
	left, right int
}

func (op *OperatorBase) SetLeft(left int) {
	op.left = left
}

func (op *OperatorBase) SetRight(right int) {
	op.right = right
}
