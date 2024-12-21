package main

import "fmt"

type person struct {
	name string
	age  int
}

func newperson(name_ string, age_ int) *person {
	return &person{
		name: name_,
		age:  age_,
	}
}

func main() {
	p1 := newperson("小米", 12)
	p2 := newperson("红米", 11)
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)
}
