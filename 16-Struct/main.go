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

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
