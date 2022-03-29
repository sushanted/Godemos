package main

import "fmt"

type Dog string

func (dog *Dog) sound() {
	fmt.Println("Bowbow")
}

func main() {
	sound := (*Dog).sound
	dog := Dog("bruno")
	sound(&dog)

	dog.sound()
}
