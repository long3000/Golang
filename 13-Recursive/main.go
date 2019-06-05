package main

import (
	"fmt"
)

func main() {
	fibonacci(6)
}

func fibonacci(n int) int {
	fmt.Printf("Current number: %d \n", n)
	if n == 0 {
		fmt.Println("Zero.")
		return 0
	} else if n == 1 {
		fmt.Println("Operation reached.")
		return 1
	}
	return (fibonacci(n - 1))
}
