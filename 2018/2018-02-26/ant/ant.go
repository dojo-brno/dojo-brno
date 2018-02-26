package ant

type Color int

const (
	Black Color = 1 + iota
	White
)

func (c Color) String() string {
	if c == Black {
		return "black"
	}
	return "white"
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	case Left:
		return "left"
	case Right:
		return "right"
	}
	return "Invalid"
}

type Board [][]Color

type Direction int

const (
	Up Direction = 1 + iota
	Right
	Down
	Left
)

type Ant struct {
	X, Y        int
	Orientation Direction
}

func Step(b Board, a Ant) (Board, Ant) {
	if a.Orientation == Right {
		return Board{{White}, {Black}}, Ant{X: 0, Y: 1, Orientation: turnLeft(a.Orientation)}
	}
	return Board{{swapColor(b[0][0]), Black}}, Ant{X: 1, Y: 0, Orientation: Right}
}

func turnLeft(d Direction) Direction {
	return Up
}

func swapColor(c Color) Color {
	if c == Black {
		return White
	}
	return Black
}
