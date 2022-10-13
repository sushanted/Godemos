package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Enter anything to abort:")

	fibCh := make(chan int)
	stdin := make(chan int)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		stdin <- -1
	}()

	go fibGenerator(fibCh)

	for {
		select {
		case fib := <-fibCh:
			fmt.Println(fib)
		case <-stdin:
			fmt.Println("Aborting")
			return
		}
	}
}

func fibGenerator(fibCh chan<- int) int {
	for i := 0; ; i++ {
		fibCh <- fib(i)
	}
}

func fib(n int) int {
	time.Sleep(10 * time.Millisecond)
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
