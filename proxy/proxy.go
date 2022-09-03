package main

type Subject interface {
	Do() string //实际业务，  业务系统，检查是否欠费，检查密码是否正确
}
type RealSubject struct {
}

func (sb RealSubject) Do() string {
	return "以太坊执行智能合约"
}

type Proxy struct {
	real  RealSubject
	money int
}

func (p Proxy) Do() string {
	if p.money > 0 {
		return p.real.Do()
	} else {
		return "请充值"
	}

}
