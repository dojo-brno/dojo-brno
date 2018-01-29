package nim

import "fmt"

type Game struct {
	Heap [3]int
}

type Move struct {
	Heap   int
	Amount int
}

type InvalidMoveError Move

func (e InvalidMoveError) Error() string {
	return fmt.Sprintf("move(HeapId %v, Amount %v)", e.Heap, e.Amount)
}

func MakeMove(g Game, m Move) (Game, error) {
	var game Game
	if m.Amount == 0 {
		return g, InvalidMoveError(m)
	}
	if m.Amount > g.Heap[m.Heap] {
		return g, InvalidMoveError(m)
	}
	copy(game.Heap, g.Heap)
	game.Heap[m.Heap] = g.Heap[m.Heap] - m.Amount
	return game, nil
}
