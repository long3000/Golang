package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// Deal function
func deal(d deck, handSize int) (deck, deck) /* Return two values of 'type' deck */ {
	return d[:handSize], d[handSize:]
}

// Convert to String function
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // Return a slice of strings then joined
}

// Save output to local storage
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // WriteFile (filename, []bytes, permission type)
}

// Read from file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil { // Error Handling
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // This package Exit funciton will quit the program any number != 0
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
