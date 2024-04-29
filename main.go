package main

import "fmt"

func main() {
	cards := newDeck()

	// shuffle the card
	cards.shuffle()

	// deal the cards to the first and second player
	firstPlayer, remainingCards := deal(cards, 5)
	secondPlayer, _ := deal(remainingCards, 5)

	// print the cards
	cards.print()

	fmt.Printf("\nFirst player: %v \nSecond player: %v", firstPlayer, secondPlayer)
}
