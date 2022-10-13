package main

import "fmt"

func main() {

	printEven := make(chan int) // capacity not specified, only used for synchronization
	printOdd := make(chan int)

	// Even printer
	go func() {
		for i := 0; i < 10; i += 2 {
			fmt.Println("waiting to print:", i)
			<-printEven
			fmt.Println(i)
			printOdd <- 1
		}
		// Finally, wait for the Odd printer to finish, receive something to unblock the send in odd printer
		<-printEven
	}()

	// Odd printer
	printEven <- 1
	for i := 1; i < 10; i += 2 {
		fmt.Println("waiting to print:", i)
		<-printOdd
		fmt.Println(i)
		printEven <- 1
	}

}
