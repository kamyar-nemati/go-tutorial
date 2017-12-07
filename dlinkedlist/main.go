package main

import (
	"fmt"
)

//a doubly linked list is made of connected nodes
type node struct {
	last *node       //to point to previous node
	data interface{} //to hold data
	next *node       //to point to next node
}

//node's constructor
func (n *node) construct() {
	n.last = nil
	n.data = nil
	n.next = nil
}

//available behaviours of our doubly linked list
type iList interface {
	init(v interface{}) *node     //to initialize the list
	goFirst() *node               //to move the pointer to the beginning of the list
	goLast() *node                //to move the pointer to the end of the list
	moveFront() *node             //to move the pointer to the next node
	moveBack() *node              //to move the pointer to the previous node
	addFront(v interface{}) *node //to add an item to the front of the pointer
	addBack(v interface{}) *node  //to add an item to the back of the pointer
}

//the doubly linked list itself
type list struct {
	count int   //to hold the number of nodes in the list
	ptr   *node //to point to a specific node in the list
	first *node //to point to the beginning of the list
	last  *node //to point to the end of the list
}

func (l *list) init(v interface{}) *node {
	//abort if it's not the first initiation
	if l.count != 0 {
		return nil
	}

	//create a node
	newNode := new(node)

	//initialize the node
	newNode.construct()
	newNode.data = v

	//as the first node in the list
	l.ptr = newNode
	l.first = newNode
	l.last = newNode

	l.count = 1

	return newNode
}

func (l *list) goFirst() *node {
	l.ptr = l.first
	return l.ptr
}

func (l *list) goLast() *node {
	l.ptr = l.last
	return l.ptr
}

func (l *list) moveFront() *node {
	//move forward if available
	if l.ptr.next != nil {
		l.ptr = l.ptr.next
	}
	return l.ptr
}

func (l *list) moveBack() *node {
	//move backward if available
	if l.ptr.last != nil {
		l.ptr = l.ptr.last
	}
	return l.ptr
}

func (l *list) addFront(v interface{}) *node {
	//reject if list is not initialized
	if l.count == 0 {
		return nil
	}

	//create and initialize a node
	newNode := new(node)
	newNode.construct()

	newNode.data = v
	newNode.last = l.ptr

	//put it into the list
	l.ptr.next = newNode

	//update the pointer
	l.ptr = newNode

	//increase the counter
	l.count++

	return newNode
}

func (l *list) addBack(v interface{}) *node {
	//reject if list is not initialized
	if l.count == 0 {
		return nil
	}

	//create and initialize a node
	newnode := new(node)
	newnode.construct()

	newnode.data = v
	newnode.next = l.ptr

	//put it into the list
	l.ptr.last = newnode

	//update the pointer
	l.ptr = newnode

	//increase the counter
	l.count++

	return newnode
}

func main() {
	//instantiate the list object
	lst := list{}
	//initialize the list
	lst.init(0)

	//add some values
	for i := 1; i < 10; i++ {
		lst.addFront(i)
	}

	//set the pointer to the beginning
	lst.goFirst()

	//read through
	for i := 1; i < lst.count; i++ {
		lst.moveFront() //let's just ignore the first value which is zero(0)

		x := (*lst.ptr).data
		fmt.Println(x)
	}
}
