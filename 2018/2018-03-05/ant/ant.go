package ant

type World struct {
	Ant   Ant
	Board Board
}

type Board [][]Color

func (w *World) Step() error {
	newAnt := w.Ant.GetNext(w.Board[w.Ant.Position[0]][w.Ant.Position[1]])
	w.Board.switchColor(w.Ant.Position[0], w.Ant.Position[1])
	w.Ant = newAnt
	return nil
}

type Ant struct {
	Position  [2]int // index 0: X, index 1: Y
	Direction int    // 0 up, 3 right, 6 down, 9 left
}

func (a Ant) GetNext(c Color) Ant {
	newAnt := a
	newAnt.updateDirection(c)
	index, value := newAnt.getIncDec()
	newAnt.Position[index] += value
	return newAnt
}

type Color int

const (
	White Color = iota // 90 degrees turn right
	Black              // 90 degrees turn left
)

func (c Color) String() string {
	if c == White {
		return "White"
	}

	return "Black"
}

func (b Board) switchColor(x, y int) {
	switch b[x][y] {
	case White:
		b[x][y] = Black
	case Black:
		b[x][y] = White
	}
}

func (a *Ant) updateDirection(c Color) {
	switch c {
	case White:
		a.Direction += 3
	case Black:
		a.Direction += 9
	}
	a.Direction %= 12
}

func (a Ant) getIncDec() (int, int) {
	switch a.Direction {
	case 0:
		return 1, 1
	case 3:
		return 0, 1
	case 6:
		return 1, -1
	case 9:
		return 0, -1
	}
	return 0, 0
}
