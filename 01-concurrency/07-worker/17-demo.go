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
	primesCh := generatePrimes(start, end, 3)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func produceNumbers(start, end int) <-chan int {
	noCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			noCh <- no
		}
		close(noCh)
	}()
	return noCh
}

func generatePrimes(start, end int, workerCount int) <-chan int {
	primesCh := make(chan int)
	noCh := produceNumbers(start, end)

	wg := &sync.WaitGroup{}
	for id := range workerCount {
		wg.Add(1)
		fmt.Printf("[generatePrimes] Worker [%d] starting...\n", id)
		go processNo(id, noCh, primesCh, wg)
	}

	go func() {
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func processNo(id int, noCh <-chan int, primesCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := range noCh {
		fmt.Printf("[processNo](%d) processing %d\n", id, no)
		if isPrime(no) {
			primesCh <- no
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
