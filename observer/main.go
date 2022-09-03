package main

func main() {
	subject := NewSubject()
	r1 := NewReader("lixiang")
	r2 := NewReader("hupeng")
	r3 := NewReader("zhangbo")
	subject.Attch(r1)
	subject.Attch(r2)
	subject.Attch(r3)
	subject.UpdateContext("妹子来了")
	r4 := NewReader("yangjie")
	subject.Attch(r4)
	subject.UpdateContext("漂亮妹子来了")
}
