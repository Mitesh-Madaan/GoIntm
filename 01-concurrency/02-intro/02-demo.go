package main

import (
	"fmt"
)

func main() {
	go f1() // scheduling the execution of f1() through the scheduler to be executed in future
	f2()

	// block so that the scheduler picks f1() for execution

	// Poor man's synchronization techniques (DO NOT USE THESE)
	// time.Sleep(1 * time.Millisecond)
	// fmt.Scanln()

}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
