package main

import (
	"fmt"
)

func main() {
	var x int

	x = 10
	fmt.Println(x)  // Print value of 'x'
	fmt.Println(&x) // Print address of 'x'

	// Declare Pointer
	var num *int
	val := new(int)

	num = new(int)
	*num = x // Set value
	val = &x // Set address

	fmt.Println("==== Pointer Num ====")
	fmt.Println(*num) // Print value of 'x'
	fmt.Println(num)  // Print address of 'x'

	fmt.Println("==== Pointer Val ====")
	fmt.Println(*val) // Print value of 'x'
	fmt.Println(val)  // Print address of 'x'
}
