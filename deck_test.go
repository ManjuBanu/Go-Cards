package main

import (
	"os"
	"testing"
)

//plan
// [newDeck()] => 1. Code to make sure that a deck is created with x number of cards
// 			   2. to make sure that the first card is an Ace of spades
// 			   3. and last card is a four of clubs

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//16 => cardSuits * cardValue in newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck  length of 16 , but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//deleting file before test
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck , but got %v", len(loadedDeck))
	}

	//deleting file after test
	os.Remove("_decktesting")
}
