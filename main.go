package main

import (
	"fmt"

	"github.com/Trinergy/cards/cards"
)

func main() {
	deck := cards.NewDeck()
	cards.PrettyPrintDeck(deck)
	cards.Shuffle(deck)
	fmt.Println("===Shuffled Deck===")
	cards.PrettyPrintDeck(deck)
	fmt.Println("===Grabbing Deck 2===")
	deck2 := cards.NewDeck()
	cards.PrettyPrintDeck(deck2)
}
