package cards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

var startDeck = Deck{
	Cards: []Card{
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
	},
}

// Deck of cards
type Deck struct {
	Cards []Card
}

// NewDeck resets the deck to a new copy of the starting deck
// reference: https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
func NewDeck() Deck {
	newDeck := make([]Card, len(startDeck.Cards))
	copy(newDeck, startDeck.Cards)

	return Deck{Cards: newDeck}
}

// Print logs all cards in deck
func (d *Deck) Print() {
	for _, card := range d.Cards {
		card.print()
	}
}

// Shuffle will shuffle the deck of cards
// reference: https://programming.guide/go/shuffle-slice-array.html
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Deal gives you # cards from the top of the deck
func (d *Deck) Deal(handSize int) Deck {
	if handSize <= 0 || handSize > 52 {
		panic(fmt.Sprintf("Can not deal %v card(s)", handSize))
	}
	hand := Deck{Cards: d.Cards[:handSize]}
	d.Cards = d.Cards[handSize:]
	return hand
}

// SaveToJSONFile converts the deck's cards to a byte slice and writes it to a json file
func (d *Deck) SaveToJSONFile(filename string) {
	file := d.toJSON()
	err := ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		panic(fmt.Sprintln("Could not write to file"))
	}
}

func (d *Deck) toJSON() []byte {
	file, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		panic(fmt.Sprintln("Could not covert to JSON"))
	}
	return file
}

// DeckFromJSONFile reads a json file and converts it into a deck
func DeckFromJSONFile(filename string) Deck {
	jsonBlob, errRead := ioutil.ReadFile(filename)

	if errRead != nil {
		panic(fmt.Sprintln("Could not read JSON file"))
	}

	var deck Deck
	err := json.Unmarshal(jsonBlob, &deck)

	if err != nil {
		panic(fmt.Sprintln("Could not read JSON file"))
	}
	return deck
}
