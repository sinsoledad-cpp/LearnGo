package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine通道b中取到了", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChanenel() {
	fmt.Println(b)
	b = make(chan int, 1)
	b <- 10
	fmt.Println("10发送到通道b中了...")
	x := <-b
	fmt.Println("通道b中取到了", x)
	close(b)
}

func main() {
	// noBufChannel()
	bufChanenel()
}
