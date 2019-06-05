package main

import (
	"fmt"
	"math"
)

func main() {
	// var radius = 2.53
	radius := 2.55
	circleArea(radius)

	calculate(2.3, 2.7, 3.4)

	aResult := advancedCalculate(4.2, 1.2, 3.12)
	fmt.Printf("Advanced function result: %.2f \n", aResult)

	prodResult, prodName := compute(2.3, 4.1, 1.21, "Product")
	fmt.Printf("Result via output 'float': %.2f \n", prodResult)
	fmt.Printf("%s \n", prodName)

	multResult := add(1, 5, 3, 2, 5, 6, 1)
	fmt.Printf("Result of multiple params addition: %d \n", multResult)

	closureFunc("BaBuun")
}

func circleArea(r float64) {
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle area (r=%.2f) = %.2f \n", r, area)
}

func calculate(a, b, c float64) {
	result := a * b * c
	fmt.Printf("a = %.2f \nb = %.2f \nc = %.2f \nResult = %.2f \n", a, b, c, result)
}

func advancedCalculate(a, b, c float64) float64 {
	result := a * b * c
	return result
}

func compute(a, b, c float64, name string) (float64, string) {
	result := a * b * c
	newStr := fmt.Sprintf("%s value = %.2f", name, result)
	return result, newStr
}

// Multiple Parameters (... can be used to declare zero or more parameters)
func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Closure Function
func closureFunc(name string) {
	add := func(a, b int) {
		result := a + b
		fmt.Printf("Addition result: %d \n", result)
	}

	sub := func(a, b int) {
		result := a - b
		fmt.Printf("Subtraction result: %d \n", result)
	}

	prod := func(a, b int) int {
		return a * b
	}

	fmt.Printf("Closure function '%s' \n", name)
	add(1, 2)
	sub(4, 2)
	val := prod(4, 2)
	fmt.Printf("Val from 'Product' operation: %d \n", val)
}
