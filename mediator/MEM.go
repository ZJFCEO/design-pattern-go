package main

import "fmt"

type MEM struct {
	data string
}

func (m *MEM) Dump(data string) {
	fmt.Println("MEM is running ", data)
}
