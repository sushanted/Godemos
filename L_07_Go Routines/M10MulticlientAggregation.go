package main

import (
	"fmt"
	"time"
)

func main() {

	clients := make([]chan int, 10)
	aggregate := make(chan int, 10)

	for i := 0; i < len(clients); i++ {
		clients[i] = make(chan int)
	}

	go func() {
		for i := 0; i < 10; i++ {
			clients[i] <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {

		for i := 0; i < len(clients); i++ {
			aggregate <- <-clients[i]
		}

	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Received: ", <-aggregate)
	}

}
