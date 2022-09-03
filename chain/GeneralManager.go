package main

import "fmt"

type GeneralManager struct {
}

func NewGeneralManager() *GeneralManager {
	return &GeneralManager{}
}
func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{&GeneralManager{}, nil}
}

func (pm *GeneralManager) HaveRight(money int) bool {
	return true
}
func (pm *GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "weishangyin" {
		fmt.Printf("GeneralManager  授权%s %d\n", name, money)
		return true
	}
	fmt.Printf("GeneralManager not 授权%s %d\n", name, money)
	return false
}
