package nim

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	var game Game
	game.Heap = [3]int{2, 0, 0}
	move := Move{0, 0}
	wantError := InvalidMoveError(move)
	_, gotError := MakeMove(game, move)
	if gotError != wantError {
		t.Errorf("MakeMove(%v, %v) error expected %v, got %v)", game, move, gotError, wantError)
	}
}
func TestMakeFirstMove(t *testing.T) {
	var game, wantGame Game
	game.Heap = [3]int{2, 0, 0}
	move := Move{0, 1}
	wantGame.Heap = [3]int{1, 0, 0}

	gotGame, err := MakeMove(game, move)
	if err != nil {
		t.Errorf("MakeMove(%v, %v) expected no error but was: %v", game, move, err)
	}
	if wantGame != gotGame {
		t.Errorf("MakeMove(%v, %v): %v, WANT: %v", game, move, gotGame, wantGame)
	}
}

func TestMakeErrorMove(t *testing.T) {
	var game, wantGame Game
	game.Heap = [3]int{2, 0, 0}
	move := Move{0, 3}
	wantGame.Heap = [3]int{2, 0, 0}
	_, err := MakeMove(game, move)
	wantError := InvalidMoveError(move)
	if err != wantError {
		t.Errorf("MakeMove(%v, %v) = (_, %v) WANT (_,%v)", game, move, err, wantError)
	}
}

func TestMakeSecondMove(t *testing.T) {
	var game, wantGame Game
	game.Heap = [3]int{2, 5, 3}
	move := Move{1, 5}
	wantGame.Heap = [3]int{2, 0, 3}

	gotGame, err := MakeMove(game, move)
	if err != nil {
		t.Errorf("MakeMove(%v, %v) expected no error but was: %v", game, move, err)
	}
	if wantGame != gotGame {
		t.Errorf("MakeMove(%v, %v): %v, WANT: %v", game, move, gotGame, wantGame)
	}
}
