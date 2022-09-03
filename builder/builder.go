package main

type Builder interface {
	Part1()
	Part2()
	Part3()
}
type Director struct {
	builder Builder //建造者的接口
}

// 创造接口
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) MakeData() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
