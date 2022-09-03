package main

type Customer interface { //接待者
	Accept(Visitor)
}
type Visitor interface { //访问者
	Visit(Customer)
}
