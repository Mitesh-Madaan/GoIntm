package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	wg := &sync.WaitGroup{}
	fmt.Println("Enter the start & end :")
	fmt.Scanln(&start, &end)

	for no := start; no <= end; no++ {
		wg.Add(1)
		go PrintIfPrime(no, wg)
	}
	wg.Wait()
}

func PrintIfPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
}
