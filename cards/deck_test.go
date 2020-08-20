package cards

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if !reflect.DeepEqual(startDeck.Cards, d.Cards) {
		t.Errorf("New Deck does not have correct starting cards")
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	d.Shuffle()

	if reflect.DeepEqual(startDeck.Cards, d.Cards) {
		t.Errorf("Cards are not shuffled")
	}
}

func TestDeal(t *testing.T) {
	handSize := 5
	d := NewDeck()

	initialDeckLength := len(d.Cards)
	hand := d.Deal(handSize)

	if len(d.Cards) != (initialDeckLength - handSize) {
		t.Errorf("Dealt cards were not removed from deck")
	}
	if len(hand.Cards) != handSize {
		t.Errorf("Cards were not dealt")
	}
}
