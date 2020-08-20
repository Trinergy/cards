package cards

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Run("Cards are shuffled", func(t *testing.T) {
		NewDeck()
		Shuffle()

		if reflect.DeepEqual(startDeck, CurrentDeck) {
			t.Errorf("Cards are not shuffled")
		}
	})
}

func TestDeal(t *testing.T) {
	t.Run("Cards are dealt", func(t *testing.T) {
		number := 5
		NewDeck()

		initialDeckLength := len(CurrentDeck)
		hand := Deal(number)
		if len(CurrentDeck) != (initialDeckLength - number) {
			t.Errorf("Dealt cards were not removed from deck")
		}
		if len(hand) != number {
			t.Errorf("Cards were not dealt")
		}
	})
}
