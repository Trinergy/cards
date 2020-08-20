package main

import (
	"github.com/Trinergy/cards/cards"
)

func main() {
	d := cards.NewDeck()
	d.Shuffle()
	d.Print()
	d.SaveToFile("test")
}
