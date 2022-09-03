package main

import "fmt"

func main() {
	adapee := NewAdaptee()       //适配器
	target := NewAdapter(adapee) //传递进入
	res := target.Request(1, 2)
	fmt.Println(res)

}
