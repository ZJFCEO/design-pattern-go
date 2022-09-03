package main

type Cloneable interface {
	Clone() Cloneable
}

// 原型对象的类
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{make(map[string]Cloneable)}
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	return t
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t   //开辟内存新建变量，存储指针指向的内容
	return &tc //返回地址
}
