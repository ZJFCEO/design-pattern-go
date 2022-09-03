package main

type AddComponent struct {
	Component
	num int
}

func WarpAddComponent(c Component, num int) Component {
	return &AddComponent{c, num}
}

// +
func (mul *AddComponent) Calc() int {
	return mul.Component.Calc() + mul.num
}
