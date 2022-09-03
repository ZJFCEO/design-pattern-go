package main

import "fmt"

func main() {
	mgr := NewPrototypeManager()
	t1 := &Type1{name: "type1"}
	mgr.Set("t1", t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone() //复制
	if t11 == t22 {
		fmt.Println("浅复制")
	} else {
		fmt.Println("深复制")
	}

}
