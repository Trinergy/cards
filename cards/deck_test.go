package cards

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Run("Cards are shuffled", func(t *testing.T) {
		d := NewDeck()
		d.Shuffle()

		if reflect.DeepEqual(startDeck.Cards, d.Cards) {
			t.Errorf("Cards are not shuffled")
		}
	})
}

func TestDeal(t *testing.T) {
	t.Run("Cards are dealt", func(t *testing.T) {
		number := 5
		d := NewDeck()

		initialDeckLength := len(d.Cards)

		hand := d.Deal(number)
		if len(d.Cards) != (initialDeckLength - number) {
			t.Errorf("Dealt cards were not removed from deck")
		}
		if len(hand.Cards) != number {
			t.Errorf("Cards were not dealt")
		}
	})
}
