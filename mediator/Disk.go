package main

import "fmt"

type Disk struct {
	data string
}

func (d *Disk) Store(data string) {
	fmt.Println("disk is running ", data)
}
