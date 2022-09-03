package main

import "fmt"

type MotherBoard struct {
}

func (*MotherBoard) WashClothes() {
	fmt.Println("委派妹子去洗衣服了")
}
func (*MotherBoard) Warmbed() {
	fmt.Println("委派妹子去暖床了")
}
