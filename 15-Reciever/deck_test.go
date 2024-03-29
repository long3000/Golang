package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Check deck size
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	// Check 1st and Last combination
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades', but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected 'Five of Clubs', but got %v", d[len(d)-1])
	}
}

// Test file IO
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards, but bot %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
