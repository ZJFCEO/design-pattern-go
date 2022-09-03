package main

type CommonMessage struct {
	method MessageImlementer
}

func NewCommonMessage(method MessageImlementer) *CommonMessage {
	return &CommonMessage{method: method}
}
func (com *CommonMessage) SendMessage(text, to string) {
	com.method.Send(text, to)
}
