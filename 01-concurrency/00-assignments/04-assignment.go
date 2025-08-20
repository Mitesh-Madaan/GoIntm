/*
Stop the workers when the user hits ENTER key
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Communicate by sharing memory

func main() {
	var start, end int
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)
	stopCh := stop()
	primesCh := generatePrimes(start, end, 3, stopCh)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func stop() <-chan struct{} {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()
	return stopCh
}

func produceNumbers(start, end int) <-chan int {
	noCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			time.Sleep(500 * time.Millisecond)
			noCh <- no
		}
		close(noCh)
	}()
	return noCh
}

func generatePrimes(start, end int, workerCount int, stopCh <-chan struct{}) <-chan int {
	primesCh := make(chan int)
	noCh := produceNumbers(start, end)

	wg := &sync.WaitGroup{}
	for id := range workerCount {
		wg.Add(1)
		fmt.Printf("[generatePrimes] Worker [%d] starting...\n", id)
		go processNo(id, noCh, primesCh, wg, stopCh)
	}

	go func() {
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func processNo(id int, noCh <-chan int, primesCh chan<- int, wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done()
LOOP:
	for no := range noCh {
		select {
		case <-stopCh:
			fmt.Printf("[processNo](%d) stop signal received.. stopping\n", id)
			break LOOP
		default:
			fmt.Printf("[processNo](%d) processing %d\n", id, no)
			if isPrime(no) {
				primesCh <- no
			}
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
