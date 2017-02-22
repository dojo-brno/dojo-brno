package gameoflife

type GenerationGrid [][]bool

func (g GenerationGrid) Next() GenerationGrid {
	if len(g) == 2 {
		return GenerationGrid{
			{CellNext(0, 0, g), CellNext(0, 1, g)},
			{CellNext(1, 0, g), CellNext(1, 1, g)},
		}
	}
	for i := range g {
		for j := range g[i] {
			g[i][j] = false
		}
	}
	return g
}

func CellNext(x, y int, g GenerationGrid) bool {
	alive := NumberOfAliveNeighbours(x, y, g)
	if alive == 3 {
		return true
	}
	return g[0][1]
}

func NumberOfAliveNeighbours(x, y int, g GenerationGrid) int {
	return 3

}
