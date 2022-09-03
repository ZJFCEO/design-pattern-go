package main

// 节点，返回一个数据
type Node interface {
	Interpret() int
}
type ValNode struct {
	val int
}

func (valn *ValNode) Interpret() int {
	return valn.val
}
