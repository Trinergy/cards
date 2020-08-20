package main

import (
	"fmt"

	"github.com/Trinergy/cards/cards"
)

func main() {
	cards.NewDeck()
	cards.PrintDeck()
	cards.Shuffle()
	fmt.Println("===Shuffled Deck===")
	cards.PrintDeck()
	fmt.Println("===BRAND NEW------------------------------------===")

	cards.NewDeck()
	cards.PrintDeck()
	// fmt.Println("===Dealt Hand===")
	// hand := cards.Deal(3)
	// cards.PrintCards(hand)
	// fmt.Println("===Dealt Deck===")
	// cards.PrintDeck()
}
