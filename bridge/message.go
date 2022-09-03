package main

// SMS短信
// Email

type AbstractMessage interface {
	SendMessage(text, to string) //发送快，发送普通速速
}

type MessageImlementer interface {
	Send(text, to string) //短信，邮件
}
