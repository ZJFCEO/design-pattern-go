package main

type CustomerCol struct {
	customers []Customer //接待者集合
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer) //叠加
}
func (c *CustomerCol) Accept(visitor Visitor) { //每个服务者，接受访问者
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}
