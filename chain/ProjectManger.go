package main

import "fmt"

type ProjectManager struct {
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{}
}
func NewProjectManagerChain() *RequestChain {
	return &RequestChain{&ProjectManager{}, nil}
}

func (pm *ProjectManager) HaveRight(money int) bool {
	return money < 1000
}
func (pm *ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "hupeng" {
		fmt.Printf("ProjectManager 授权%s %d\n", name, money)
		return true
	}
	fmt.Printf("ProjectManager not 授权%s %d\n", name, money)
	return false
}
