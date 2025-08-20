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

	// context to carry data
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	// cancellable context by time
	ctx, cancel := context.WithTimeout(valCtx, 5*time.Second)

	ch := genNos(ctx)

	fmt.Println("Hit ENTER to stop...auto stops after 5 secs....")
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
	fmt.Println("[genNos] root-key :", ctx.Value("root-key"))
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				if ctx.Err() == context.Canceled {
					fmt.Println("[genNos] programmatic cancellation signal received")
				}
				if ctx.Err() == context.DeadlineExceeded {
					fmt.Println("[genNos] timeout cancellation signal received")
				}
				break LOOP
			case ch <- (i + 1) * 10:
			}
		}
		close(ch)
	}()
	return ch
}
