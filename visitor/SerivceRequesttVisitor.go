package main

import "fmt"

type SerivceRequesttVisitor struct {
}

func (*SerivceRequesttVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Println("serving  Enterprise  customer", c.name)
	case *IndividualCustomer:
		fmt.Println("serving  Individual Customer", c.name)
	}
}
