package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "# of gorountines to spin up!")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... Hit ENTER to start..\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1)
		go fn(id+1, wg)
	}
	wg.Wait()
	fmt.Println("Done!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn [%d] - started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn [%d] - completed\n", id)
}
