package main

type Operator interface {
	SetLeft(left int)
	SetRight(right int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}
