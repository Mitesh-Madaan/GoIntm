package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("App Panicked, err :", e)
		}
		fmt.Println("Thank you!")
	}()
	// async

	// the user chooses to look for the error and handle it
	/*
		resultCh, errCh := divideAsync(100, 0)
		select {
		case result := <-resultCh:
			fmt.Println("Quotient :", result)
		case e := <-errCh:
			fmt.Println("DivideAsync Error :", e)
		}
	*/

	// the user chooses NOT to check for error
	resultCh, _ := divideAsync(100, 0)
	result := <-resultCh
	fmt.Println("Quotient :", result)

	// Sync
	/*
		quotient := divideSync(100, 0)
		fmt.Println("Quotient :", quotient)
	*/

}

func divideSync(multipler, divisor int) int {

	result := multipler / divisor
	return result

}

func divideAsync(multipler, divisor int) (<-chan int, <-chan error) {
	resultCh := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				errCh <- e.(error)
			}
		}()
		result := multipler / divisor
		resultCh <- result
	}()
	return resultCh, errCh
}
