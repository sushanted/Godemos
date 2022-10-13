package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(3 * time.Second)
			channel <- i
		}
	}()

	// Poll every second
	for i := 0; i < 10; i++ {
		select {
		case result := <-channel: // This is non-blocking, the default will be executed
			fmt.Println("Result available:", result)
		default:
			fmt.Println("Will try again after sometime")
			time.Sleep(time.Second)
		}
	}

}
