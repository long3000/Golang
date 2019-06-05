package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Define array
	fmt.Println("Define map")
	products := make(map[string]int) // make - built-in function to create dynamic array
	customers := make(map[string]int)

	// Insert data
	fmt.Println(">>>> Insert Map data")
	products["product_1"] = rand.Intn(100)
	products["product_2"] = rand.Intn(100)

	customers["customer_1"] = rand.Intn(100)
	customers["customer_2"] = rand.Intn(100)

	// Display Data
	fmt.Println(">>>>> Display Data")
	fmt.Println(products["product_1"])
	fmt.Println(products["product_2"])

	fmt.Println(customers["customer_1"])
	fmt.Println(customers["customer_2"])
}
