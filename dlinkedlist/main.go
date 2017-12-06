package main

type node struct {
	last *node
	data interface{}
	next *node
}

type iList interface {
	init() *node
	goFirst() *node
	goLast() *node
	moveFront() *node
	moveBack() *node
	addFront() *node
	addBack() *node
}

type list struct {
	count int
	ptr   *node
	first *node
	last  *node
}

func main() {
}
