package bowling

import "testing"

func TestBowlingGame(t *testing.T) {
	tests := []struct {
		game Game
		want int
	}{
		{
			game: Game{},
			want: 0,
		},
		{
			game: Game{{1}},
			want: 1,
		},
		{
			game: Game{{1}, {1}},
			want: 2,
		},
		{
			game: Game{{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
			want: 10,
		},
		{
			game: Game{{1, 1}},
			want: 2,
		},
		{
			game: Game{{1, 9}, {1}},
			want: 12,
		},
		{
			game: Game{{}, {}, {}, {}, {}, {}, {}, {}, {}, {1, 9}, {1}},
			want: 11,
		},
		{
			game: Game{{10}, {1, 1}},
			want: 14,
		},
		{
			game: Game{{}, {}, {}, {}, {}, {}, {}, {}, {}, {10}, {1, 2}},
			want: 13,
		},
		{
			game: Game{{10}, {10}, {1, 2}},
			want: 21 + 13 + 3,
		},
		{
			game: Game{{}, {}, {}, {}, {}, {}, {}, {}, {10}, {10}, {10, 10}},
			want: 30 + 30,
		},
		{
			game: Game{{10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10}, {10, 10}},
			want: 300,
		},
		{
			game: Game{{5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5}},
			want: 150,
		},
		{
			game: Game{{0, 10}, {10}, {1, 2}},
			want: 20 + 13 + 3,
		},
	}
	for _, tt := range tests {
		if got := Score(tt.game); got != tt.want {
			t.Errorf("Score(%v) = %v want %v", tt.game, got, tt.want)
		}
	}
}
