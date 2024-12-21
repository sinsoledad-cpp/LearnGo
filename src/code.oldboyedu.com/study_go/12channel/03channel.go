package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChane = make(chan *result, 100)
var wg sync.WaitGroup

func zhoulin(zl chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func baodelu(zl <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-zl
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go zhoulin(jobChan)
	wg.Add(24)
	for i:=0;i<24;i++{
		go baodelu(jobChan,resultChane)
	}
	for result:=range resultChane{
		fmt.Printf("value:%d sum:%d\n",result.job.value,result.result)
	}
}
