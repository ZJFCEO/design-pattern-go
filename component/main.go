package main

func main() {
	root := NewComponent(CompositeNode, "root")
	r1 := NewComponent(CompositeNode, "r1")
	r2 := NewComponent(CompositeNode, "r2")
	r3 := NewComponent(CompositeNode, "r3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(r1)
	root.AddChild(r2)
	r1.AddChild(r3)

	r1.AddChild(l1)

	r2.AddChild(l2)
	r2.AddChild(l3)
	root.Print("")
}
