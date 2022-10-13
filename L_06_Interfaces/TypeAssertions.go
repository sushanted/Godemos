package main

import "fmt"

type Animal interface {
	Sound() string
}

type Herbivore interface {
	Animal
	Herbs() string
}

type MilkProducer interface { // NOTE this doesn't have Sound method
	Milk() string
}

type Goat struct {
	Name string
}

func (goat Goat) Sound() string {
	return "ba"
}

func (goat Goat) Herbs() string {
	return "hay"
}

func (goat Goat) Milk() string {
	return "protein"
}

type Sheep struct {
	Name string
}

func (sheep Sheep) Sound() string {
	return "mehemehe"
}

func main() {

	// NOTE : .(Goat) and .(\*Goat) are  different types and need to be asserted differently

	//The interface can  point to any of them

	assertSimpleType()

	assertPointerType()

	assertInterfaceType()

	assertNonHierarchyInterfaceType()

	assertTypeWithCheck()

	assertTypeWithPointerCheck()

}

func assertSimpleType() {

	var animal Animal = Goat{"aba"}

	goat := animal.(Goat)

	fmt.Println(goat.Name) // now can directly use .Name

}

func assertPointerType() {
	var animalPointer Animal = &Goat{"ababa"} // The interface can point to a pointer as well

	goatPointer := animalPointer.(*Goat) // The conversion here should be with a pointer

	fmt.Println(goatPointer.Name)
}

func assertInterfaceType() {

	var animal Animal = Goat{"bababa"}

	herbAnimal := animal.(Herbivore)

	fmt.Println(herbAnimal.Herbs()) // now Herbs is accessible
}

func assertNonHierarchyInterfaceType() {
	var animal Animal = Goat{"maaa"}

	herbAnimal := animal.(MilkProducer)

	fmt.Println(herbAnimal.Milk()) // now Milk is accessible
}

func assertTypeWithCheck() {

	var animal Animal = Sheep{"shepa"}

	if sheep, ok := animal.(Sheep); ok {
		fmt.Println("This is a sheep: ", sheep.Name)
	} else {
		fmt.Println("Not a sheep")
	}

	if _, ok := animal.(Goat); !ok {
		fmt.Println("Not a goat : ", animal)
	}

	// reuse of name
	if animal, ok := animal.(Sheep); ok {
		// NOTE : animal is a new variable in the scope of if block
		fmt.Println("This is a sheep: ", animal.Name)
	} else {
		fmt.Println("Not a sheep")
	}

	// fmt.Println(animal.Name) // here the animal is a variable in outer scope, not asserted as a Sheep

}

func assertTypeWithPointerCheck() {
	var animal Animal = &Goat{"gotta"}

	if _, ok := animal.(Goat); !ok {
		fmt.Println("Animal is not a goat, but maybe a pointer to goat")
	}

	if goat, ok := animal.(*Goat); ok {
		fmt.Println("Animal is a pointer to goat :", goat)
	}

}
