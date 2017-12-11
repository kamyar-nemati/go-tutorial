package main

import (
	"fmt"
)

//the basic structure of each individual node
type node struct {
	last *node       //to point to previous node
	data interface{} //to hold data
	next *node       //to point to next node
}

//node's constructor
func (n *node) construct(v *interface{}) {
	n.last = nil
	n.data = *v
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
	remove() *node                //to delete the node to which it's been pointed
}

//the doubly linked list itself
type list struct {
	count int   //to hold the number of nodes in the list
	ptr   *node //to point to a specific node in the list
	first *node //to point to the beginning of the list
	last  *node //to point to the end of the list
}

func (l *list) init(v interface{}) *node {
	//abort if it's already initiated
	if l.count != 0 {
		return nil
	}

	//create a node
	newNode := new(node)

	//initialize the node
	newNode.construct(&v)

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
	//reject if the list is not yet initialized
	if l.count == 0 {
		return nil
	}

	//create and initialize a node
	newNode := new(node)
	newNode.construct(&v)

	/* squeeze the node into the list */
	newNode.next = l.ptr.next
	if l.ptr.next != nil {
		l.ptr.next.last = newNode
	} else {
		l.last = newNode
	}
	newNode.last = l.ptr
	l.ptr.next = newNode
	/* - */

	//update the pointer
	l.ptr = newNode

	//increase the counter
	l.count++

	return newNode
}

func (l *list) addBack(v interface{}) *node {
	//reject if the list is not yet initialized
	if l.count == 0 {
		return nil
	}

	//create and initialize a node
	newnode := new(node)
	newnode.construct(&v)

	/* squeeze the node into the list */
	newnode.last = l.ptr.last
	if l.ptr.last != nil {
		l.ptr.last.next = newnode
	} else {
		l.first = newnode
	}
	newnode.next = l.ptr
	l.ptr.last = newnode
	/* - */

	//update the pointer
	l.ptr = newnode

	//increase the counter
	l.count++

	return newnode
}

func (l *list) remove() *node {
	//reject if the list is not yet initialized
	if l.count == 0 {
		return nil
	}

	//keep a temporary pointer to the current node
	y := l.ptr

	if l.count == 1 { //list has only one item
		l.first = nil
		l.ptr = nil
		l.last = nil
	} else { //list has multiple items
		if l.ptr.last == nil { //delete the first item
			l.ptr = l.ptr.next
			l.first = l.ptr
			//remove the pointer to the deleted node
			l.ptr.last = nil
		} else if l.ptr.next == nil { //delete the last item
			l.ptr = l.ptr.last
			l.last = l.ptr
			//remove the pointer to the deleted node
			l.ptr.next = nil
		} else { //delete an item from the middle of the list
			//detaching the node to be deleted
			l.ptr.last.next = l.ptr.next
			l.ptr.next.last = l.ptr.last
			//update the pointer
			l.ptr = l.ptr.next
		}
	}

	//perhaps we should detache the dangling node from the list
	y.last = nil
	y.next = nil

	//decrease the counter
	l.count--

	return l.ptr
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

	//list is: 0-1-2-3-4-5-6-7-8-9

	lst.goFirst()   //we're at 0
	lst.moveFront() //move to 1
	lst.moveFront() //move to 2
	lst.remove()    //delete 2 and move to 3
	lst.moveFront() //mote to 4
	lst.moveFront() //mote to 5
	lst.moveFront() //mote to 6
	lst.remove()    //delete 6 and move to 7
	lst.remove()    //delete 7 and move to 8
	lst.moveFront() //move to 9
	lst.remove()    //delete 9 and move to 8
	lst.remove()    //delete 8 and move to 5

	//list is now: 0-1-3-4-5

	//set the pointer to the beginning
	lst.goLast()

	//read through in reverse order
	for i := 1; i < lst.count; i++ {
		x := (*lst.ptr).data
		fmt.Println(x)

		lst.moveBack() //let's just ignore the first value which is zero(0)
	}
}
