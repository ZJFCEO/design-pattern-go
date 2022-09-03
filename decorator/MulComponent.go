package main

type MulComponent struct {
	Component
	num int
}

func WarpMulComponent(c Component, num int) Component {
	return &MulComponent{c, num}
}

// 乘法
func (mul *MulComponent) Calc() int {
	return mul.Component.Calc() * mul.num
}
