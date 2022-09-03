package main

import "fmt"

// 适配的目标接口
type Target interface {
	Request(int, int) string
}

type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

func (adap *adapter) Request(a, b int) string {
	return adap.SpecificRequest(a, b)
}

// 被适配
type Adaptee interface {
	SpecificRequest(int, int) string
}

// 载体，被适配的目标类
type adapteeImpl struct {
}

// 工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// 实际函数
func (ada *adapteeImpl) SpecificRequest(a, b int) string {
	fmt.Println(a, b)
	return "SpecificRequest你妹"
}
