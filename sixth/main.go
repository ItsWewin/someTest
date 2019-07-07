package main

import (
	"fmt"
)

func main() {
	myChan := make(chan int)
	go func() {
		myChan <- 1
		fmt.Printf("goroutine0")
		// fmt.Printf("goroutine1")
		// myChan <- 2
		// myChan <- 1
		// myChan <- 2
		// fmt.Printf("goroutine2")
		// myChan <- 1
		// myChan <- 2
	}()
	// time.Sleep(2 * time.Second)
	<-myChan
}
