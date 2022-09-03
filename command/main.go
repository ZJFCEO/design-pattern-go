package main

func main() {
	mb := &MotherBoard{}
	cmd1 := NewMMCommand1(mb)
	cmd2 := NewMMCommand2(mb)

	box1 := NewBox(cmd1, cmd2)
	box1.GoWarmBed()
	box1.GoWashclothes()

	//box2:=NewBox(cmd1, cmd2)
	//box2.GoWarmBed()
	//box2.GoWashclothes()
}
