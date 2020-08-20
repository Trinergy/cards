package cards

import (
	"fmt"
	"math/rand"
	"time"
)

var startDeck = []Card{
	{1, "Spades"},
	{2, "Spades"},
	{3, "Spades"},
	{4, "Spades"},
	{5, "Spades"},
	{6, "Spades"},
	{7, "Spades"},
	{8, "Spades"},
	{9, "Spades"},
	{10, "Spades"},
	{11, "Spades"},
	{12, "Spades"},
	{13, "Spades"},
	{1, "Hearts"},
	{2, "Hearts"},
	{3, "Hearts"},
	{4, "Hearts"},
	{5, "Hearts"},
	{6, "Hearts"},
	{7, "Hearts"},
	{8, "Hearts"},
	{9, "Hearts"},
	{10, "Hearts"},
	{11, "Hearts"},
	{12, "Hearts"},
	{13, "Hearts"},
	{1, "Clubs"},
	{2, "Clubs"},
	{3, "Clubs"},
	{4, "Clubs"},
	{5, "Clubs"},
	{6, "Clubs"},
	{7, "Clubs"},
	{8, "Clubs"},
	{9, "Clubs"},
	{10, "Clubs"},
	{11, "Clubs"},
	{12, "Clubs"},
	{13, "Clubs"},
	{1, "Diamonds"},
	{2, "Diamonds"},
	{3, "Diamonds"},
	{4, "Diamonds"},
	{5, "Diamonds"},
	{6, "Diamonds"},
	{7, "Diamonds"},
	{8, "Diamonds"},
	{9, "Diamonds"},
	{10, "Diamonds"},
	{11, "Diamonds"},
	{12, "Diamonds"},
	{13, "Diamonds"},
}

// CurrentDeck being used
var CurrentDeck []Card

// NewDeck resets the CurrentDeck to a new copy of the starting deck
func NewDeck() []Card {
	newDeck := make([]Card, len(startDeck))
	copy(newDeck, startDeck)

	CurrentDeck = newDeck

	return CurrentDeck
}

//PrintCards logs all cards in deck parameter
func PrintCards(deck []Card) {
	for _, card := range deck {
		fmt.Printf("%v of %v\n", CardTranslation[card.value], card.suite)
	}
}

// PrintDeck logs all cards in the current deck
func PrintDeck() {
	PrintCards(CurrentDeck)
}

// Shuffle will shuffle the deck of cards
// reference: https://programming.guide/go/shuffle-slice-array.html
func Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(CurrentDeck), func(i, j int) { CurrentDeck[i], CurrentDeck[j] = CurrentDeck[j], CurrentDeck[i] })
}

// Deal gives you # cards from the top of the deck
func Deal(number int) []Card {
	if number <= 0 || number > 52 {
		panic(fmt.Sprintf("Can not deal %v card(s)", number))
	}
	hand := CurrentDeck[:number]
	CurrentDeck = CurrentDeck[number:]
	return hand
}
