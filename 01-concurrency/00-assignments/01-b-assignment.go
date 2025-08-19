package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		// executing the "sequential" function as a concurrent operation
		go func() {
			defer wg.Done()
			PrintIfPrime(no)
		}()
	}
	wg.Wait()

}

// designed to be executed sequentially
func PrintIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
