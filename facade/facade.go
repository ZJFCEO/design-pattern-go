package main

import "fmt"

// 测试
type TAPI interface {
	Test() string
}
type APICall struct {
	a AmouleAPI
	b BmouleAPI
}

func (api *APICall) Test() string {
	return fmt.Sprintf("%s\n%s", api.a.TestA(), api.b.TestB())
}
func NewTAPI() TAPI {
	return &APICall{NewAmouleAPI(), NewBmouleAPI()}
}

// A测试
type AmouleAPI interface {
	TestA() string
}
type aMoubleImpl struct{}

func (*aMoubleImpl) TestA() string {
	return "主网已经运行"
}
func NewAmouleAPI() AmouleAPI {
	return &aMoubleImpl{}
}

// B测试
type BmouleAPI interface {
	TestB() string
}
type bMoubleImpl struct{}

func (*bMoubleImpl) TestB() string {
	return "测试网已经运行"
}
func NewBmouleAPI() BmouleAPI {
	return &bMoubleImpl{}
}
