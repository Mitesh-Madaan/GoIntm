package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	stopCh := timeout(5 * time.Second)
	ch := genNos(stopCh)

	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

}

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}

// producer
func genNos(stopCh <-chan time.Time) <-chan int {
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
