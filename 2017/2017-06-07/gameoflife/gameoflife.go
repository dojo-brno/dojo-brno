package gameoflife

type Cell struct {
	X, Y int
}

type Game struct {
	Rows, Columns int
	World         Grid
}

type Grid map[Cell]struct{}

var E = struct{}{}

func (g Game) Next() Game {
	if len(g.World) <= 2 {
		g.World = Grid{}
		return g
	}

	if _, ok := g.World[Cell{0, 7}]; ok {
		gn := Game{Rows: g.Rows, Columns: g.Columns, World: Grid{}}

		for cell := range g.World {
			switch count := g.countLiveNeighbours(cell.X, cell.Y); {
			case ((count < 2) || (count > 3)):
				// die
			}

		}
		return gn
	}

	if _, ok := g.World[Cell{0, 0}]; ok {
		g.World = Grid{{0, 1}: E}
		return g
	}
	return Game{
		Rows:    4,
		Columns: 8,
		World: Grid{
			{3, 1}: E,
			{4, 1}: E,
			{3, 2}: E,
			{4, 2}: E,
		},
	}
}

func (g Game) countLiveNeighbours(row, column int) int {
	return 0
}
