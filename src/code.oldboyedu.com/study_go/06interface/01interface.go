package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//使用值接收者实现了接口的所有方法

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("🐱吃%s\n", food)
}

//使用指针接收者实现了接口的所有方法

// func (c *cat) move() {
// 	fmt.Println("走猫步")
// }

// func (c *cat) eat(food string) {
// 	fmt.Printf("🐱吃%s\n", food)
// }

func main() {
	var a1 animal
	fmt.Println(a1)

	c1 := cat{"tom", 4}
	c2 := &cat{"java", 4}

	// //使用值接收者实现了接口的所有方法
	a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

	//	使用指针接收者实现了接口的所有方法
	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
