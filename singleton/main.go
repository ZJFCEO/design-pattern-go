package main

import "fmt"

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1.data)
	fmt.Println(s2.data)
	fmt.Println(s1 == s2)
}
