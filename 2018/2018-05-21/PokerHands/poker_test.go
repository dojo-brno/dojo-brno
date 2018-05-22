package poker

import "testing"

func TestIsStrainghtFlush(t *testing.T) {
	tests := []struct {
		hand Hand
		want bool
	}{
		{
			hand: Hand{
				Card{Value: Two, Suit: Spades},
				Card{Value: Two, Suit: Hearts},
				Card{Value: Two, Suit: Clubs},
				Card{Value: Three, Suit: Spades},
				Card{Value: Three, Suit: Diamonds},
			},
			want: false,
		},
		{
			hand: Hand{
				Card{Value: Two, Suit: Spades},
				Card{Value: Three, Suit: Spades},
				Card{Value: Four, Suit: Spades},
				Card{Value: Five, Suit: Spades},
				Card{Value: Six, Suit: Spades},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		if got := IsStrainghtFlush(tt.hand); got != tt.want {
			t.Errorf("IsStrainghtFlush(%v) = %v, WANT = %v", tt.hand, got, tt.want)
		}
	}

}
