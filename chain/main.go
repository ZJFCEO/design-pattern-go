package main

func main() {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuceessor(c2)
	c2.SetSuceessor(c3)

	var c Manger = c1
	c.HandleFeeRequest("hupeng", 1500)
	c.HandleFeeRequest("hupeng", 500)
	c.HandleFeeRequest("lixiang", 4500)
	c.HandleFeeRequest("lixiang", 11500)
	c.HandleFeeRequest("weishangyin", 4500)
	c.HandleFeeRequest("weishangyin", 11500)
}
