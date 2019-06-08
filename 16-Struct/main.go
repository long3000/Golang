package main

import "fmt"

//__________________________
//| firstName --> <string> |
//| lastName  --> <string> |
//|________________________|
//|__________________________|
//|  fistName --> "Alex"     |
//|  lastName --> "Anderson" |
//|--------------------------|

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName:  "Wick",
		contactInfo: contactInfo{
			email:   "jwick@email.com",
			zipCode: 59124,
		},
	}
	fmt.Println(john)
	john.print()

	john.updateName("James")
	john.print()
	// fmt.Printf("%+v \n", john)

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var john person
	// fmt.Println(john)
	// fmt.Printf("%+v \n", john)
}

// Reciever for 'person' type
func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
