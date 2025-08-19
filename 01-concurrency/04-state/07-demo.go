package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex //struct composition
	// state
	count int
}

func (sf *SafeCounter) Add(delta int) {
	sf.Lock()
	{
		sf.count++
	}
	sf.Unlock()
}

var sf SafeCounter

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
	fmt.Println("count :", sf.count)
}

// behavior
func increment() {
	sf.Add(1)
}
