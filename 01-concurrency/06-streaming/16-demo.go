package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()

	for data := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	fmt.Println("Done!")

}

// producer
func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	fmt.Printf("[genNos] count = %d\n", count)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
		close(ch)
	}()
	return ch
}
