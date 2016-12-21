package pokerhands_test

import (
	"testing"

	"github.com/dojo-brno/dojo-brno/2016/2016-04-27/pokerhands"
)

func TestPokerHandsEqual(t *testing.T) {
	hand1 := pokerhands.Hand{"2S", "3S", "4S", "5S", "7H"}
	hand2 := pokerhands.Hand{"2H", "3H", "4H", "5H", "7S"}
	want := 0
	if got := pokerhands.Compare(hand1, hand2); got != want {
		t.Errorf("Compare(%v, %v) = %v, want %v", hand1, hand2, got, want)
	}
}

func TestPokerHandsGreater(t *testing.T) {
	hand1 := pokerhands.Hand{"2H", "3H", "4H", "5H", "8S"}
	hand2 := pokerhands.Hand{"2S", "3S", "4S", "5S", "7H"}
	want := 1
	if got := pokerhands.Compare(hand1, hand2); got != want {
		t.Errorf("Compare(%v, %v) = %v, want %v", hand1, hand2, got, want)
	}
}

func TestPokerHandsLesser(t *testing.T) {
	tests := []struct {
		left, right pokerhands.Hand
	}{
		{
			pokerhands.Hand{"2S", "3S", "4S", "5S", "7H"}, pokerhands.Hand{"2H", "3H", "4H", "5H", "8S"},
		},
		// swaping 4th and 5th on right hand
		{
			pokerhands.Hand{"2S", "3S", "4S", "5S", "7H"}, pokerhands.Hand{"2H", "3H", "4H", "8S", "5H"},
		},
	}
	for _, hands := range tests {
		if got := pokerhands.Compare(hands.left, hands.right); got != -1 {
			t.Errorf("Compare(%v, %v) = %v, want -1", hands.left, hands.right, got)
		}

	}
}
