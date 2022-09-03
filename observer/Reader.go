package main

import "fmt"

// 读取类
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}

func (r *Reader) Update(subject *Subject) {
	fmt.Printf("%s 收到 %s\n", r.name, subject.context)
}
