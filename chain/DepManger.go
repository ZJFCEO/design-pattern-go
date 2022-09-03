package main

import "fmt"

type DepManager struct {
}

func NewDepManager() *DepManager {
	return &DepManager{}
}
func NewDepManagerChain() *RequestChain {
	return &RequestChain{&DepManager{}, nil}
}

func (dm *DepManager) HaveRight(money int) bool {
	return money < 5000
}
func (dm *DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "lixiang" {
		fmt.Printf("DepManager  授权%s %d\n", name, money)
		return true
	}
	fmt.Printf("DepManager not 授权%s %d\n", name, money)
	return false
}
