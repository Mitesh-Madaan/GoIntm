package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	stopCh := make(chan struct{})
	ch := genNos(stopCh)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()

	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

}

// producer
func genNos(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			case ch <- (i + 1) * 10:
			}
		}
		close(ch)
	}()
	return ch
}
