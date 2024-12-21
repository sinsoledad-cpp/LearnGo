package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var y int64

var wg sync.WaitGroup

func add() {
	x++
	atomic.AddInt64(&y, 1)
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(y)

	z := int64(2)
	ok := atomic.CompareAndSwapInt64(&z, 3, 4)
	fmt.Println(ok, z)
	ok = atomic.CompareAndSwapInt64(&z, 2, 4) 
	fmt.Println(ok, z)
}
