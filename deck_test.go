package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// t here is test handler
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected King of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
