package main

import (
	"github.com/Trinergy/cards/cards"
)

func main() {
	d := cards.NewDeck()
	d.SaveToJSONFile("test.json")
	readDeck := cards.DeckFromJSONFile("test.json")
	readDeck.Print()
}
