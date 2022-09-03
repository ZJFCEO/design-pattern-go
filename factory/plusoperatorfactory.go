package main

// 操作的抽象
type PlusOperatorFactory struct{}

// 操作类中包含操作数
type PlusOperator struct {
	*OperatorBase
}

// 实际运行
func (p PlusOperator) Result() int {
	return p.left + p.right
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{&OperatorBase{}}
}
