package main

import "fmt"

func main() {

	done := make(chan int) // To wait for the consumer to finish

	sharedChannel := make(chan int)

	go producer(sharedChannel)
	go consumer(sharedChannel, done)

	<-done // wait for the consumer to complete
}

func producer(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	// NOTE:

	// var bidir chan int = in // cannot assign a send-only channel to a bidirectional channel
	// <-in // not allowed on a send-only channel
}

func consumer(out <-chan int, done chan<- int) {
	for value := range out {
		fmt.Println(value)
	}
	close(done) // This also works to unblock the receiver

	// NOTE:

	// var bidir chan int = out // cannot assign a receive-only channel to a bidirectional channel
	// out <- 2 // Not allowed on a receive-only channel
	// close(out) // Not allowed on a receive-only channel

}
