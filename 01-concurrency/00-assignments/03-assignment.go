/*
Replace "Communicate by sharing memory" with "Share memory by communicating", namely "use channels"
*/

package main

import (
	"fmt"
	"sync"
)

// Communicate by sharing memory

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	primesCh := generatePrimes(start, end)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func generatePrimes(start, end int) <-chan int {
	primesCh := make(chan int)

	var wg sync.WaitGroup
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				primesCh <- no
			}
		}()
	}
	go func() {
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
