package main

import (
	"os"
	"testing"
)

// testing the newDeck() function - ensures that a new deck is created with the correct cards
func TestNewDeck(t *testing.T) {
	d := newDeck()
	length := len(d)
	firstCard := d[0]
	lastCard := d[length-1]

	if length != 16 {
		t.Errorf("Expected deck length of 16, but got %v", length)
	}

	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", firstCard)
	}

	if lastCard != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", lastCard)
	}
}

// testing the saveToDeck() and newDeckFromFile() functions - ensures that the deck is saved to the user's hard drive and loaded without any issues
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting.txt")

	// create a new deck and save it to the user's hard drive
	d := newDeck()
	d.saveToFile("_deckTesting.txt")

	loadedDeck := newDeckFromFile("_deckTesting.txt")
	lengthOfLoadedDeck := len(loadedDeck)

	if lengthOfLoadedDeck != 16 {
		t.Errorf("Expected deck length of 16, but got %v", lengthOfLoadedDeck)
	}

	os.Remove("_deckTesting.txt")
}
