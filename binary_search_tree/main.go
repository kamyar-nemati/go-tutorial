package main

type node struct {
	data  interface{}
	left  *node
	right *node
}

func (n *node) construct() {
	n.data = nil
	n.left = nil
	n.right = nil
}

type bst struct {
	tree *node
}

type iBst interface {
	init()
	add(v interface{})
}

func main() {
}
