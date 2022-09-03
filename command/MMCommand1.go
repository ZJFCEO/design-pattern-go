package main

type MMCommand1 struct {
	mb *MotherBoard
}

func NewMMCommand1(mb *MotherBoard) *MMCommand1 {
	return &MMCommand1{mb}
}

func (mmc1 MMCommand1) Execute() {
	mmc1.mb.WashClothes() //洗衣服
}
