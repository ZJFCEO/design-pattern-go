package main

import "fmt"

func main() {
	v1 := NewImageViewer("1.jpg")
	v1.Display()
	v2 := NewImageViewer("1.jpg")
	v2.Display()
	if v1.ImageFlyWeight == v2.ImageFlyWeight {
		fmt.Println("节约内存")
	} else {
		fmt.Println("langfei内存")
	}
}
