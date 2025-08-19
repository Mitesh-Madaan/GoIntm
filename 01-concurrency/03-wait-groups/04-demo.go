package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //increment the counter by 1
	go f1()   // scheduling the execution of f1() through the scheduler to be executed in future
	f2()

	// block so that the scheduler picks f1() for execution

	wg.Wait() // block until the wg counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
