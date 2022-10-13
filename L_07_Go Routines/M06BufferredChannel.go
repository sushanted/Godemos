package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("\nWith unbuffered channel")
	// Here the worker is blocked till the there is someone available to collect the results
	// it cannot move ahead (or finish) till then.
	// Observe : written result is printed only after the time.Sleep is complete
	writeResultsOnChannel(make(chan int))

	fmt.Println("\nWith buffered channel")
	// Here the worker is free to move ahead (or finish) for next execution after writing the results
	// As the results would be buffered and would be available for collection at any later point in time
	// Observe : written result is printed as soon as results are written
	writeResultsOnChannel(make(chan int, 1))

}

func writeResultsOnChannel(ubc chan int) {

	go complete(func() int {
		return 5
	}, ubc)

	fmt.Println("Waiting for results")
	time.Sleep(2 * time.Second)
	fmt.Println("Retrieving results")
	fmt.Println("Result", <-ubc)
}

func complete(task func() int, result chan<- int) {
	fmt.Println("Executing task")
	result <- task()
	fmt.Println("Written results")
}
