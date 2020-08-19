package cards

import (
	"fmt"
)

const CardTranslation := map[int]string {
	1: "ace",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
	10: "ten",
	11: "jack",
	12: "queen",
	13: "king",
}

const Deck := []struct {
	value int
	suite string
}{
	{1, "spades"},
	{2, "spades"},
	{3, "spades"},
	{4, "spades"},
	{5, "spades"},
	{6, "spades"},
	{7, "spades"},
	{8, "spades"},
	{9, "spades"},
	{10, "spades"},
	{11, "spades"},
	{12, "spades"},
	{1, "hearts"},
	{2, "hearts"},
	{3, "hearts"},
	{4, "hearts"},
	{5, "hearts"},
	{6, "hearts"},
	{7, "hearts"},
	{8, "hearts"},
	{9, "hearts"},
	{10, "hearts"},
	{11, "hearts"},
	{12, "hearts"},
	{1, "clubs"},
	{2, "clubs"},
	{3, "clubs"},
	{4, "clubs"},
	{5, "clubs"},
	{6, "clubs"},
	{7, "clubs"},
	{8, "clubs"},
	{9, "clubs"},
	{10, "clubs"},
	{11, "clubs"},
	{12, "clubs"},
	{1, "diamonds"},
	{2, "diamonds"},
	{3, "diamonds"},
	{4, "diamonds"},
	{5, "diamonds"},
	{6, "diamonds"},
	{7, "diamonds"},
	{8, "diamonds"},
	{9, "diamonds"},
	{10, "diamonds"},
	{11, "diamonds"},
	{12, "diamonds"},
}

func NewDeck() struct {
	newDeck := Deck
	return newDeck
}