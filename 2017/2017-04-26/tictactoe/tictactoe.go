package tictactoe

type Game struct {
	Rows [][]Value
}

type Value uint8

const (
	Empty Value = iota
	X
	O
)

func (g Game) HasWinner() bool {
	return g.hasRowWinner() || g.hasColumnWinner()
}

func (g Game) hasRowWinner() bool {
	for _, row := range g.Rows {
		if isRowWinner(row) {
			return true
		}
	}
	return false
}

func (g Game) hasColumnWinner() bool {
	if len(g.Rows) == 0 {
		return false
	}
	for column := 0; column < len(g.Rows[0]); column++ {
		if g.Rows[0][column] != Empty && g.Rows[1][column] == g.Rows[0][column] && g.Rows[2][column] == g.Rows[0][column] {
			return true
		}
	}
	return false
}

func isRowWinner(row []Value) bool {
	for _, item := range row {
		if item != row[0] || item == Empty {
			return false
		}
	}
	return true
}
