package main

import (
	"fmt"
	"sync"
)

// state
var count int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()
}
