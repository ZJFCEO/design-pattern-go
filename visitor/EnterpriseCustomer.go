package main

type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{name: name}
}
func (ec *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(ec)
}
