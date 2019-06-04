package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Define Slice
	fmt.Println("Define Slices")
	var numbers []int
	numbers = make([]int, 5)
	matrix := make([][]int, 3*3)

	// Insert data
	fmt.Println(">>>>> Insert slice data")
	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100) // Random number
	}

	// Insert matrix data
	fmt.Println(">>>>> Insert slice matrix data")
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	// Display Data
	fmt.Println(">>>>> Display Slice data")
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
	}

	// Display Matrix data
	fmt.Println(">>>>> Display Slice 2D data")
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
		// for j := 0; j < 3; j++ {
		// 	fmt.Println(matrix[i][j])
		// }
	}
}
