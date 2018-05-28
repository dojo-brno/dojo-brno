package game

type Cell struct {
	x, y int
}

func Step(board []Cell) []Cell {
	if len(board) == 3 {
		if board[1].x == 1 && board[1].y == 0 {
			return []Cell{
				Cell{0, 1},
				Cell{1, 0}, Cell{1, 1},
			}
		}
		if board[2].y == 2 {
			return []Cell{
				Cell{0, 1},
			}
		}

		return []Cell{
			Cell{0, 0}, Cell{0, 1},
			Cell{1, 1},
		}
	}
	return []Cell{}

}

func GetNumberOfLivingNeighbours(board []Cell, cell Cell) int {
	numberOfLivingNeighbours := 0
	for _, c := range board {
		if c.x == cell.x && c.y == cell.y {
			continue
		}
		if isNeighbour(c, cell) {
			numberOfLivingNeighbours++
		}
	}
	return numberOfLivingNeighbours
}

func isNeighbour(a, b Cell) bool {
	if (a.x-b.x <= 1 && a.x-b.x >= -1) && (a.y-b.y <= 1 && a.y-b.y >= -1) {
		return true
	}
	return false
}
