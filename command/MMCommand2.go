package main

type MMCommand2 struct {
	mb *MotherBoard
}

func NewMMCommand2(mb *MotherBoard) *MMCommand2 {
	return &MMCommand2{mb}
}

func (mmc1 MMCommand2) Execute() {
	mmc1.mb.Warmbed() //洗衣服
}
