package main

// 简单工厂设计模式
// 解决软件的可拓展性问题

type API interface {
	Say(name string) string
}

func NewAPI(str string) API {
	if str == "zh" {
		return &Chinese{}
	} else if str == "en" {
		return &English{}
	} else {
		return nil
	}
}

type Chinese struct{}

func (c *Chinese) Say(name string) string {
	return "你好，" + name
}

type English struct{}

func (e *English) Say(name string) string {
	return "hello, " + name
}
