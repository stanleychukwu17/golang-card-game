package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	// hands, remainingCards := deal(cards, 5)
}
