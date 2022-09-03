package main

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{name: name}
}
func (ec *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(ec)
}
