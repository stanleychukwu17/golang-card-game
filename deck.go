package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// creates a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// prints the deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns a new hand of cards from the deck of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// a receiver function that converts the deck to strings so that they can be saved in a file
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves the deck to a file to the user hard drive
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// reading a deck of file saved on the hard drive
func newDeckFromFile(filename string) deck {
	// where bs = []byte (i.e a slice of bytes)
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // means there was an error during execution
	}

	// string(bs) converts the byte slice to a string
	// then strings.Split converts the string to a slice of strings
	s := strings.Split(string(bs), ",")

	// deck(s) converts the []string to a *type deck*
	// converting the []string into a *type deck* allows us to be able to use all the methods on the *type deck*
	return deck(s)
}

// a function that shuffles the deck of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // time.Now().UnixNano() returns the current time in nanoseconds - this is used to seed the random number generator
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // generates a random int64 number btw 0 and the (len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // swap positions in the deck
	}
}
