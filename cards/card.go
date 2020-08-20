package cards

import "fmt"

// CardTranslation is a table to map number to string
var CardTranslation = map[int]string{
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

// Card for playing
type Card struct {
	Value int
	Suit  string
}

func (c *Card) print() {
	fmt.Printf("%v of %v\n", CardTranslation[c.Value], c.Suit)
}
