package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Define array
	fmt.Println("Define Array")
	var numbers [5]int
	var cities [5]string
	var matrix [3][3]int // array 2D

	// Insert data
	fmt.Println(">>>>>> insert matrix data")
	for i := 0; i < 5; i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}

	// Insert data matrix
	fmt.Println(">>>>>> insert matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// Display data
	fmt.Println(">>>>>> Display data array")
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
		fmt.Println(cities[j])
	}

	// Display data matrix
	fmt.Println(">>>>>> Display data array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
