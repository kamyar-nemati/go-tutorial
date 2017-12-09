package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) construct(v *int) {
	n.data = *v
	n.left = nil
	n.right = nil
}

type bst struct {
	root  *node
	count int
}

type iBst interface {
	init(v int)
	add(v int)
	find(v int)
	traversePreorder()
	traverseInorder()
	traversePostorder()
}

func (b *bst) init(v int) {
	if b.count != 0 {
		return
	}

	newNode := new(node)

	newNode.construct(&v)

	b.root = newNode

	b.count++
}

func insert(n *node, v *int) {
	if *v <= n.data && n.left != nil {
		insert(n.left, v)
		return
	}

	if *v > n.data && n.right != nil {
		insert(n.right, v)
		return
	}

	newNode := new(node)
	newNode.construct(v)

	if *v <= n.data {
		n.left = newNode
	} else {
		n.right = newNode
	}
}

func (b *bst) add(v int) {
	if b.count == 0 {
		return
	}

	insert(b.root, &v)

	b.count++
}

func search(n *node, v *int) {
	if n.data == *v {
		fmt.Println(n.data)
		return
	}

	if *v <= n.data {
		fmt.Println("Going left")
		search(n.left, v)
	} else {
		fmt.Println("Going right")
		search(n.right, v)
	}
}

func (b *bst) find(v int) {
	search(b.root, &v)
}

func preorder(n *node) {
	fmt.Println(n.data)

	if n.left != nil {
		preorder(n.left)
	}

	if n.right != nil {
		preorder(n.right)
	}
}

func inorder(n *node) {
	if n.left != nil {
		inorder(n.left)
	}

	fmt.Println(n.data)

	if n.right != nil {
		inorder(n.right)
	}
}

func postorder(n *node) {
	if n.left != nil {
		postorder(n.left)
	}

	if n.right != nil {
		postorder(n.right)
	}

	fmt.Println(n.data)
}

func (b *bst) traversePreorder() {
	preorder(b.root)
}

func (b *bst) traverseInorder() {
	inorder(b.root)
}

func (b *bst) traversePostorder() {
	postorder(b.root)
}

func main() {
	t := bst{}
	t.init(5)

	t.add(4)
	t.add(6)
	t.add(2)
	t.add(3)
	t.add(7)
	t.add(8)

	fmt.Println("Preorder traverse:")
	t.traversePreorder()

	fmt.Println("Inorder traverse: (sorted)")
	t.traverseInorder()

	fmt.Println("Postorder traverse:")
	t.traversePostorder()

	fmt.Println("Finding 3:")
	t.find(3)
}
