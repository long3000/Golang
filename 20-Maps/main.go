package main

import (
	"fmt"
)

func main() {
	// Declare a MAP
	// SIMPLE WAY
	colors := map[string]string{
		"red":         "#CC2A36",
		"green":       "#36CC2A",
		"blue":        "#2A36CC",
		"dark orchid": "B23AEE",
	}

	// VARIABLE
	var colorsVar map[string]string

	// VARIBALE Assignment
	colorsMake := make(map[string]string)

	//Assign key-value
	colorsMake["white"] = "#FFFFFF"
	colorsMake["black"] = "#000000"

	//Remove/Delete pair member in MAP
	delete(colorsMake, "white")

	fmt.Println(colors)
	fmt.Println(colorsVar)
	fmt.Println(colorsMake)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color '%s' has hex: %s \n", color, hex)
	}
}
