package gameoflife

type Board [4][8]bool

func printBoard(b Board) string {
	value := ""
	for _, row := range b {
		for _, cell := range row {
			if cell {
				value += "*"
			} else {
				value += "-"
			}
		}
		value += "\n"
	}
	return value
}

func Survive(nLivingNeighbours int) bool {
	if nLivingNeighbours == 2 || nLivingNeighbours == 3 {
		return true
	}
	return false
}

func Born(nLivingNeighbours int) bool {
	if nLivingNeighbours == 3 {
		return true
	}
	return false
}

func GameOfLife(board Board) Board {
	return Board{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}
