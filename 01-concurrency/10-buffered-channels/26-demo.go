package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// consumer
func main() {
	fmt.Println("Process Id :", os.Getpid())
	stopCh := make(chan struct{})
	ch := genNos(stopCh)

	go func() {
		interruptCh := make(chan os.Signal, 1)
		signal.Notify(interruptCh, os.Interrupt)
		for {
			<-interruptCh
			close(stopCh)
		}
	}()

	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done!")

}

// producer
func genNos(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				fmt.Printf("Stop signal received.. ")
				break LOOP
			default:
				ch <- (i + 1) * 10
			}
		}
		close(ch)
	}()
	return ch
}
