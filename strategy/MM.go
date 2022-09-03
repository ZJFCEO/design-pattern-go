package main

import "fmt"

type MMContext struct {
	Name        string
	Age         int
	paoStrategy MMStrategy
}

func NewMMContext(name string, age int, mms MMStrategy) *MMContext {
	return &MMContext{name, age, mms}
}
func (mm *MMContext) Pao() {
	mm.paoStrategy.Pao(mm)
}

type MMStrategy interface {
	Pao(*MMContext)
}

type Girl struct {
}

func (*Girl) Pao(ctx *MMContext) {
	fmt.Println("带妹子去看遍世间繁华", ctx.Name)
}

type Women struct {
}

func (*Women) Pao(ctx *MMContext) {
	fmt.Println("带大妈取做旋转木马", ctx.Name)
}
