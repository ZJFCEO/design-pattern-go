package main

import "fmt"

func main() {
	var c Component = &ConcreateComponent{}
	//c=WarpAddComponent(c,10)
	c = WarpMulComponent(c, 10)
	fmt.Println(c.Calc())
}
