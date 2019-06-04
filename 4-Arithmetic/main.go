package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare variables
	var a, b int
	var aa float64

	// Assign values
	a = 5
	b = 10
	aa = 2.4

	// Arithmetic operation
	// Addition
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)
	// Subtraction
	d := a - b
	fmt.Printf("%d - %d = %d \n", a, b, d)
	// Division
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n", a, b, e)
	// Multiplication
	f := a * b
	fmt.Printf("%d x %d = %d \n", a, b, f)
	// Power
	g := math.Pow(aa, 2)
	fmt.Printf("%.2f^%d = %.2f \n", aa, 2, g)
	// Sin
	h := math.Sin(aa)
	fmt.Printf("Sin(%.2f) = %.2f \n", aa, h)
	// Square Root
	j := math.Sqrt(aa)
	fmt.Printf("Sqrt(%.2f) = %.2f \n", aa, j)
}
