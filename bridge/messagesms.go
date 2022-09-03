package main

import "fmt"

type MessageSMS struct {
}

func ViaSMS() MessageImlementer {
	return &MessageSMS{}
}

func (msms *MessageSMS) Send(text, to string) { //短信发送
	fmt.Printf("send %s to %s via SMS", text, to)
}
