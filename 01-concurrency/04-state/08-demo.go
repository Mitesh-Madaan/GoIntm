package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// state
var count int64

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
	fmt.Println("count :", count)
}

// behavior
func increment() {
	atomic.AddInt64(&count, 1)
}
