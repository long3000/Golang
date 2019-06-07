package main

func main() {
	// REPLACED WITH newDeck()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// Replaced this Loop with Reciever
	cards.print()
}

// Reciever function

// REPLACED WITH newDeck()
// func newCard() string {
// 	return "Five of Diamonds"
// }
