package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200
	data := <-ch
	fmt.Println(data)
	data = <-ch
	fmt.Println(data)
}
