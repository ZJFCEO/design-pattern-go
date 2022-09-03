package main

func main() {
	//m:=Bridge.NewCommonMessage(Bridge.ViaSMS())
	//m:=Bridge.NewComomnMessage(Bridge.ViaEmail())
	//m.SendMessage("baBy 你好","nimei")
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("baBy 你好", "nimei")
}
