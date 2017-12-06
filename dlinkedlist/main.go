package main

type node struct {
	last *node
	data interface{}
	next *node
}

func (n *node) construct() {
	(*n).last = nil
	(*n).data = nil
	(*n).next = nil
}

type iList interface {
	init(v interface{}) *node
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

func (l *list) init(v interface{}) {
	if (*l).count != 0 {
		return
	}

	newList := new(node)

	(*newList).construct()
	(*newList).data = v

	(*l).ptr = newList
	(*l).first = newList
	(*l).last = newList
}

func (l *list) goFirst() {
	(*l).ptr = (*l).first
}

func (l *list) goLast() {
	(*l).ptr = (*l).last
}

func (l *list) moveFront() {
	if (*(*l).ptr).next == nil {
		return
	}

	(*l).ptr = (*(*l).ptr).next
}

func (l *list) moveBack() {
	if (*(*l).ptr).last == nil {
		return
	}

	(*l).ptr = (*(*l).ptr).last
}

func main() {
}
