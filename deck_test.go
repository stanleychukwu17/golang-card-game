package main

import "testing"

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
