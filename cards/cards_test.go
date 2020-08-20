package cards

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Run("Cards are shuffled", func(t *testing.T) {
		d := NewDeck()
		d.shuffle()

		if reflect.DeepEqual(startDeck.cards, d.cards) {
			t.Errorf("Cards are not shuffled")
		}
	})
}

func TestDeal(t *testing.T) {
	t.Run("Cards are dealt", func(t *testing.T) {
		number := 5
		d := NewDeck()

		initialDeckLength := len(d.cards)
		fmt.Println(initialDeckLength)

		hand := d.deal(number)
		fmt.Println(len(d.cards))
		if len(d.cards) != (initialDeckLength - number) {
			t.Errorf("Dealt cards were not removed from deck")
		}
		if len(hand.cards) != number {
			t.Errorf("Cards were not dealt")
		}
	})
}
