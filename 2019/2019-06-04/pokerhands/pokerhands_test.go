package pokerhands

import "testing"

func TestCompare(t *testing.T) {
	highCardLowest := Hand{{Value: Two, Colour: Clubs}, {Value: Four, Colour: Clubs}, {Value: Five, Colour: Clubs}, {Value: Six, Colour: Clubs}, {Value: Seven, Colour: Diamonds}}
	highCardMedium := Hand{{Value: Two, Colour: Diamonds}, {Value: Four, Colour: Diamonds}, {Value: Five, Colour: Diamonds}, {Value: Six, Colour: Diamonds}, {Value: Ten, Colour: Clubs}}
	highCardHighest := Hand{{Value: Three, Colour: Hearts}, {Value: Five, Colour: Hearts}, {Value: Seven, Colour: Hearts}, {Value: Ten, Colour: Hearts}, {Value: Ace, Colour: Spades}}
	onePairHand := Hand{{Value: Two, Colour: Clubs}, {Value: Two, Colour: Diamonds}, {Value: Five, Colour: Clubs}, {Value: Six, Colour: Clubs}, {Value: Seven, Colour: Diamonds}}
	lowerPairHand := Hand{{Value: Two, Colour: Clubs}, {Value: Two, Colour: Diamonds}, {Value: Five, Colour: Clubs}, {Value: Six, Colour: Clubs}, {Value: Seven, Colour: Diamonds}}
	higherPairHand := Hand{{Value: Three, Colour: Clubs}, {Value: Three, Colour: Diamonds}, {Value: Five, Colour: Clubs}, {Value: Six, Colour: Clubs}, {Value: Seven, Colour: Diamonds}}
	tests := []struct {
		handA, handB Hand
		want         Result
	}{
		{handA: highCardLowest, handB: highCardMedium, want: BWon},
		{handA: highCardHighest, handB: highCardMedium, want: AWon},
		{handA: highCardMedium, handB: highCardMedium, want: Tie},
		{handA: onePairHand, handB: highCardMedium, want: AWon},
		{handA: highCardMedium, handB: onePairHand, want: BWon},
		{handA: onePairHand, handB: onePairHand, want: Tie},
		{handA: lowerPairHand, handB: higherPairHand, want: BWon},
	}

	for _, tt := range tests {
		if got := Compare(tt.handA, tt.handB); tt.want != got {
			t.Errorf("Compare(%v, %v) = (%v) WANT (%v)", tt.handA, tt.handB, got, tt.want)
		}
	}
}
