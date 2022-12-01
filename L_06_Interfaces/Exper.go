package main

import "fmt"

func main() {
	colorMap := map[string]int{
		"red":    1,
		"green":  2,
		"blue":   3,
		"yellow": 4,
	}

	// Simple iteration, we get all values
	for key := range colorMap {
		fmt.Printf("Color is %s\n", key)
	}
}
