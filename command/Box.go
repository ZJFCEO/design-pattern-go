package main

import "fmt"

type Box struct {
	Washclothes Command
	WarmBed     Command
}

func NewBox(Washclothes, WarmBed Command) *Box {
	return &Box{Washclothes, WarmBed}
}

func (b *Box) GoWashclothes() {
	fmt.Println("给妹子买了一束花")
	b.Washclothes.Execute()
}
func (b *Box) GoWarmBed() {
	fmt.Println("给妹子买了一个iphone max")
	b.WarmBed.Execute()
}
