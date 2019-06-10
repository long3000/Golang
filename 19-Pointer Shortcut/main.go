package main

import (
	"fmt"
)

type contact struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	john := person{
		firstName: "Jon",
		lastName:  "Wick",
		contact: contact{
			email:   "johnw@email.com",
			zipcode: 91213,
		},
	}
	// fmt.Println(john)
	john.print()

	// jonPointer := &john
	// jonPointer.updateName("John")

	// STRUCTS are reference type thus reciever pointer type
	// can be referenced back to 'person'
	john.updateName("John")
	john.print()
}
