package main

import "fmt"

func main() {
	// Declare variables
	var str string
	var n, m int
	var mn float32

	// Assign values
	str = "Hello"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("Value of 'Str' = ", str)
	fmt.Println("Value of 'n' = ", n)
	fmt.Println("Value of 'm' = ", m)
	fmt.Println("Value of 'mn' = ", mn)

	// Declare and assign values to variables
	var city string = "London"
	var x int = 100

	fmt.Println("Value of city = ", city)
	fmt.Println("Value of n = ", x)

	// Declare variable with deifninng its type
	country := "DE"
	val := 15
	fmt.Println("Value of 'country' = ", country)
	fmt.Println("Value of 'val' = ", val)

	// Define multiple variables
	var (
		name  string
		email string
		age   int
	)
	name = "John"
	email = "john@email.com"
	age = 25

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
