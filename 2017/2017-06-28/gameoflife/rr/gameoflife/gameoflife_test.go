package gameoflife

import "testing"

func TestGameNext(t *testing.T) {
	tests := []struct {
		name string
		game Game
		want Game
	}{
		{
			name: "empty world stays empty",
			game: NewGame(0, 0),
			want: NewGame(0, 0),
		},
		{
			name: "empty world stays empty no matter the size",
			game: NewGame(3, 3),
			want: NewGame(3, 3),
		},
		{
			name: "[1] live cell dies with less than 2 live neighbours",
			game: NewGame(3, 3).Add(2, 2),
			want: NewGame(3, 3),
		},
		{
			name: "[2] live cell dies with more than 3 live neighbours",
			game: NewGame(3, 3).Add(0, 1).Add(1, 0).Add(1, 1).Add(1, 2).Add(2, 1),
			want: NewGame(3, 3).Add(0, 1).Add(1, 0).Add(1, 2).Add(2, 1),
		},
		{
			name: "[3] live cell with 3 live neighbours stays alive",
			game: NewGame(2, 2).Add(0, 0).Add(0, 1).Add(1, 0).Add(1, 1),
			want: NewGame(2, 2).Add(0, 0).Add(0, 1).Add(1, 0).Add(1, 1),
		},
		{
			name: "[4] dead cell with 3 live neighbours becomes alive",
			game: NewGame(2, 2).Add(0, 0).Add(0, 1).Add(1, 0),
			want: NewGame(2, 2).Add(0, 0).Add(0, 1).Add(1, 0).Add(1, 1),
		},
	}
	for _, tt := range tests {
		if got := tt.game.Next(); got != tt.want {
			t.Errorf("game.Next() = %v, want %v", got, tt.want)
		}
	}
}

func TestNumberOfLiveNeighbours(t *testing.T) {
	g := NewGame(3, 5).Add(0, 1).Add(1, 0).Add(1, 1).Add(1, 2).Add(2, 1)
	tests := []struct {
		x, y int
		want int
	}{
		{0, 0, 3},
		{1, 0, 3},
		{1, 1, 4},
		{2, 3, 1},
		{2, 4, 0},
	}
	for _, tt := range tests {
		if got := g.LiveNeighbours(tt.x, tt.y); got != tt.want {
			t.Errorf("g.LiveNeighbours(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
		}
	}
}
