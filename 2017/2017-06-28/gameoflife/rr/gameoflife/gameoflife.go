package gameoflife

type Cell struct {
	x, y int
}

type Game struct {
	cells      []Cell
	rows, cols int
}

func NewGame(x, y int) Game {
	return Game{}
}

func (g Game) Add(x, y int) Game {
	g.cell.append(Cell{x, y})
	return g
}

func (g Game) Next() Game {
	return g
}

func (g Game) LiveNeighbours(x, y int) int {
	return 0
}
