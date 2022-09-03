package main

type RequestChain struct {
	Manger
	successor *RequestChain //处理成功
}

// 判断责任链在哪里结束
func (rc *RequestChain) SetSuceessor(endrc *RequestChain) {
	rc.successor = endrc
}

// 权力判断
func (rc *RequestChain) HaveRight(money int) bool {
	return true
}

// 向后传递,有权利当场处理，没有权利放弃，向后传递
func (rc *RequestChain) HandleFeeRequest(name string, money int) bool {
	if rc.Manger.HaveRight(money) {
		return rc.Manger.HandleFeeRequest(name, money)
	}
	if rc.successor != nil {
		return rc.successor.HandleFeeRequest(name, money)
	}
	return false
}
