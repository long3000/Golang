package main

func main() {
	// REPLACED WITH newDeck()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("cards_deck")
	cards := newDeckFromFile("cards_deck")
	cards.print()

}
