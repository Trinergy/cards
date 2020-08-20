package main

import (
	"fmt"

	"github.com/Trinergy/cards/cards"
)

func main() {
	d := cards.NewDeck()
	d.Print()
	d.Shuffle()
	fmt.Println("===Shuffled Deck===")
	d.Print()
	fmt.Println("===BRAND NEW------------------------------------===")

	cards.NewDeck()
	d.Print()
	// fmt.Println("===Dealt Hand===")
	// hand := cards.Deal(3)
	// cards.PrintCards(hand)
	// fmt.Println("===Dealt Deck===")
	// d.Print()
}
