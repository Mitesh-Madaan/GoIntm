package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// state
var count atomic.Int64

func main() {
	wg := &sync.WaitGroup{}
	for range 300 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("count :", count.Load())
}

// behavior
func increment() {
	count.Add(1)
}
