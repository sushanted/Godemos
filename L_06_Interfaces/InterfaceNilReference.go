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

	if personInfo != nil {
		fmt.Println(personInfo.String())
	}

}
