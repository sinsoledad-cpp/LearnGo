package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

func f1() {
	for i := 0; i < 100; i++ {
		go hello(i)
	}
}

func f2() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	// f1()
	f2()
	fmt.Println("main")
	time.Sleep(time.Second)
}
