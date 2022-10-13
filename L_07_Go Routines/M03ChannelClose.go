package main

import "fmt"

func main() {

	numbers := createNumbersChannel()

	// when numbers is closed, ok becomes false
	for i, ok := <-numbers; ok; i, ok = <-numbers {
		fmt.Println("Received: ", i, ok)
	}

	// Follows a simple variant of for loop for a channel, breaking when the channel closes

	numbers = createNumbersChannel()

	// when numbers is closed, the loop is also terminated
	for i := range numbers {
		fmt.Println("Received: ", i)
	}

}

func createNumbersChannel() (numbers chan int) {

	numbers = make(chan int)

	go generateNumbers(numbers)

	return numbers
}

func generateNumbers(numbers chan int) {
	defer close(numbers) // Once channel is closed, retriever keep getting the 0 value
	for i := 0; i < 10; i++ {
		fmt.Println("Sending ", i)
		numbers <- i
	}
	fmt.Println("Done with generation")
}
