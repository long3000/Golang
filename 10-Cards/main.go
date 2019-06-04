package main

func main() {
	cards := deck{newCard(), newCard()}    // Declare SLICE of type 'String' elements
	cards = append(cards, "Six of Spades") // Add element(s) to SLICE (append DOES NOT modify the 'card' Slice but rather creating a new one)

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
