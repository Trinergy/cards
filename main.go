package main

import (
	"fmt"

	"github.com/Trinergy/cards/cards"
)

func main() {
	deck := cards.NewDeck()
	cards.PrettyPrintCards(deck)
	cards.Shuffle(deck)
	fmt.Println("===Shuffled Deck===")
	cards.PrettyPrintCards(deck)
	fmt.Println("===Grabbing Deck 2===")
	deck2 := cards.NewDeck()
	cards.PrettyPrintCards(deck2)
	fmt.Println("===Dealt Hand===")
	hand := cards.Deal(3, deck2)
	cards.PrettyPrintCards(hand)
	fmt.Println("===Dealt Deck===")
	cards.PrettyPrintCards(deck2)
}
