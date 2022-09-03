package main

func main() {
	c := CustomerCol{}
	c.Add(NewEnterpriseCustomer("Microsoft"))

	c.Add(NewIndividualCustomer("billgates"))
	c.Add(NewEnterpriseCustomer("Google"))
	//c.Accept(&SerivceRequesttVisitor{})
	c.Accept(&AnalysisVisitor{})
}
