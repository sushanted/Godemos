package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person *Person) String() string {
	return fmt.Sprintf("%s has age %d", person.Name, person.Age)
}

func main() {

	var person *Person

	fmt.Printf("%T %v\n", person, person)

	var personInfo fmt.Stringer

	personInfo = person

	fmt.Printf("%T %v\n", personInfo, personInfo)

	if person != nil {
		fmt.Println(person.String())
	}

	fmt.Println("Trying with interface")

	fmt.Printf("Type: %T\n", personInfo) // This is not nil

	fmt.Printf("Value: %v\n", personInfo) // This would be nil

	fmt.Printf("Is nil: %v\n", personInfo == nil) // This is not overall nil as it contains the type as Person

	// This would cause a fault
	if personInfo != nil {
		fmt.Println(personInfo.String())
	}

}
