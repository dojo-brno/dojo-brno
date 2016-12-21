package minesweeper

import "testing"

func TestMinesweeper(t *testing.T) {
	field := "*"
	want := "*"
	if got := revealAdjacentMines(field); got != want {
		t.Errorf("revealAdjacentMines(%v) = %v, want %v", got, field, want)
	}
}
