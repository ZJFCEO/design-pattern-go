package main

import "fmt"

func main() {
	meditor := GetMediatorInstace()
	fmt.Println(meditor)
	meditor.Ccpu.Process("hello")
	meditor.Cgpu.Display("hello")
	meditor.Cdisk.Store("hello")
	meditor.Cmem.Dump("gogogo")
	meditor.Changed(meditor.Ccpu)
	meditor.Changed(meditor.Cmem)
}
