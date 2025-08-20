package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch, count := genNos()

	for range count {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}

}

// producer
func genNos() (<-chan int, int) {
	ch := make(chan int)
	count := rand.Intn(20)
	fmt.Printf("[genNos] count = %d\n", count)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
	}()
	return ch, count
}
