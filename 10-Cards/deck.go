package main

import "fmt"

// Create a new type of 'Deck'
// which is a slice of strings
type deck []string

// Create a function that print values in Type Deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
