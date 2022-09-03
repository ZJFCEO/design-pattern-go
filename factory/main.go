package main

import "fmt"

func main() {
	fac := PlusOperatorFactory{}
	op := fac.Create()
	op.SetLeft(2)
	op.SetRight(5)
	fmt.Println(op.Result())

	fac2 := PlusOperatorFactory{}
	op2 := fac2.Create()
	op2.SetLeft(9)
	op2.SetRight(5)
	fmt.Println(op2.Result())

	fac3 := SubOperatorFactory{}
	op3 := fac3.Create()
	op3.SetLeft(9)
	op3.SetRight(5)
	fmt.Println(op3.Result())
}
