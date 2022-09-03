package main

import "fmt"

func main() {
	builder := &StringBuilder{}
	dict := NewDirector(builder)
	dict.MakeData()
	fmt.Println(builder.GetResult())
}
