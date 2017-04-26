package tictactoe

import "testing"

func TestEmptyGame(t *testing.T) {
	game := Game{}
	want := false
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestRowWinnerX(t *testing.T) {
	game := Game{Rows: [][]Value{
		{Empty, Empty, Empty},
		{X, X, X},
		{O, O, Empty},
	}}
	want := true
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestRowWinnerO(t *testing.T) {
	game := Game{Rows: [][]Value{
		{Empty, X, Empty},
		{O, O, O},
		{X, X, Empty},
	}}
	want := true
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestRowWinnerX_0(t *testing.T) {
	game := Game{Rows: [][]Value{
		{X, X, X},
		{O, Empty, Empty},
		{Empty, O, Empty},
	}}
	want := true
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestRowNoWinner(t *testing.T) {
	game := Game{Rows: [][]Value{
		{Empty, Empty, Empty},
		{X, O, X},
		{O, O, Empty},
	}}
	want := false
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestColumnWinner(t *testing.T) {
	game := Game{Rows: [][]Value{
		{X, Empty, Empty},
		{X, O, Empty},
		{X, O, Empty},
	}}
	want := true
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestColumnWinner_1(t *testing.T) {
	game := Game{Rows: [][]Value{
		{Empty, O, Empty},
		{X, O, Empty},
		{X, O, Empty},
	}}
	want := true
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}

func TestColumnNoWinner(t *testing.T) {
	game := Game{Rows: [][]Value{
		{Empty, X, Empty},
		{Empty, O, X},
		{Empty, O, X},
	}}
	want := false
	if got := game.HasWinner(); got != want {
		t.Errorf("game.HasWinner() = %v, want %v", got, want)
	}
}
