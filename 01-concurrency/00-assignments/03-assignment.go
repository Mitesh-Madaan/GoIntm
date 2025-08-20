/*
Replace "Communicate by sharing memory" with "Share memory by communicating", namely "use channels"
*/

package main

import (
	"fmt"
	"sync"
)

// Communicate by sharing memory
var primes []int
var mutex sync.Mutex

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

	// wait for the goroutines to complete
	wg.Wait()

	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}

}

// designed to be executed sequentially
func PrintIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}
