package main

import "fmt"

// Create a new type of 'dec;'
// Which is a slice (mutable array) of string

type deck []string

func newDeck() deck {
	cards := deck{} // Type deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards // Return cards 'slice'
}

// Reciever -> ()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
