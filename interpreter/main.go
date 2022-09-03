package main

import "fmt"

func main() {
	//1+2
	//1+2+3
	//1+2+3+5-4
	p := &Parser{}
	fmt.Print("start\n")
	p.Parse("1 + 3 - 2 + 5 - 6 + 7 - 10")
	fmt.Println(p.Result().Interpret())
}
