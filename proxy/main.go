package main

import "fmt"

func main() {
	var sub Subject
	sub = &Proxy{money: -1000}
	fmt.Println(sub.Do())
}
