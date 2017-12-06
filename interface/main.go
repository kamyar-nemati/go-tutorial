package main

import "fmt"

type iActions interface {
	walk() string
	talk() string
}

type human struct {
	name string
	age  int
}

func (h human) walk() {
	fmt.Printf("%s is WALKING for %d year(s).\n", h.name, h.age)
}

func (h human) talk() {
	fmt.Printf("%s is TALKING for %d year(s).\n", h.name, h.age)
}

func main() {
	k := human{name: "Kamyar", age: 30}
	k.walk()
	k.talk()
}
