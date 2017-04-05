package tictactoe

import "testing"

func TestHasWinner(t *testing.T) {
	tests := []struct {
		game Game
		want bool
	}{
		{nil, false},
		{
			Game{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			false,
		},
		{
			Game{
				[]int{X, O, O},
				[]int{E, X, E},
				[]int{E, E, X},
			},
			true,
		},
		{
			Game{
				[]int{E, O, X},
				[]int{O, X, E},
				[]int{X, E, E},
			},
			true,
		},
		{
			Game{
				[]int{X, X, X},
				[]int{O, O, E},
				[]int{E, E, E},
			},
			true,
		},
		{
			Game{
				[]int{O, O, E},
				[]int{X, X, X},
				[]int{E, E, E},
			},
			true,
		},
		{
			Game{
				[]int{X, O, E},
				[]int{X, E, O},
				[]int{X, E, E},
			},
			true,
		},
		{
			Game{
				[]int{X, O, E},
				[]int{X, O, X},
				[]int{E, O, E},
			},
			true,
		},
		{
			Game{
				[]int{X},
			},
			true,
		},
	}
	for _, tt := range tests {
		if got := tt.game.HasWinner(); got != tt.want {
			t.Errorf("%v HasWinner() = %v, want %v", tt.game, got, tt.want)
		}
	}
}
