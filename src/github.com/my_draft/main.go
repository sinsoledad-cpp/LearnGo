package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Move(name string) {
	d.Name = name
}

func (d *Dog) Speak() {
	fmt.Println("Woof! I'm", d.Name)
}

func main() {
	var a Animal

	// 修改接口字段的实际值
	dog := &Dog{Name: "Charlie"}
	a = dog // a now points to a new *Dog

	a.Speak() // 输出: Woof! I'm Charlie

	// 修改接口字段的实际值
	dog.Move("Bob")
	a.Speak() // 输出: Woof! I'm Bob
}
