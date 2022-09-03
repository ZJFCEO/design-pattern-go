package main

import "fmt"

type MMfeel struct {
	Tall     int
	Rich     int
	Handsome int
}

var States []MMfeel = make([]MMfeel, 0)
var Stateslength int = 0

func (mmf *MMfeel) A第一次见面(tall int, rich int, handsome int) {

	mmf.Tall = tall
	mmf.Rich = rich
	mmf.Handsome = handsome
	fmt.Println("当前状态", mmf.Tall, mmf.Rich, mmf.Handsome)
	States = append(States, *mmf)
}
func (mmf *MMfeel) A去韩国() {
	mmf.Handsome = mmf.Handsome + 30
	States = append(States, *mmf)
}
func (mmf *MMfeel) A中彩票() {
	mmf.Rich += 5000000
	States = append(States, *mmf)
}
func (mmf *MMfeel) A断腿骨() {
	mmf.Tall += 10
	States = append(States, *mmf)
}
func (mmf *MMfeel) A妹子的分数() {
	fmt.Println("当前状态", mmf.Tall, mmf.Rich, mmf.Handsome)
}
