package main

import (
	"context"
	"fmt"
	"time"
)

// consumer
func main() {
	// top most context
	rootCtx := context.Background()

	// cancellable context
	ctx, cancel := context.WithCancel(rootCtx)

	ch := genNos(ctx)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel() //send the cancellation signal through the context
	}()

	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}

}

// producer
func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				break LOOP
			case ch <- (i + 1) * 10:
			}
		}
		close(ch)
	}()
	return ch
}
