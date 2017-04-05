package tictactoe

// < 0 means x
// > 0 means o
// = 0 means empty
type Game [][]int

const (
	X = -1
	O = 1
	E = 0
)

func (g Game) HasWinner() bool {
	return g.isWinner(X) || g.isWinner(O)
}

func (g Game) isWinner(value int) bool {
	return g.isDiagonalEqualTo(value) || g.isAntidiagonalEqualTo(value) || g.isAnyRowFilledWith(value) || g.isAnyColumnFilledWith(value)
}

func (s Game) isDiagonalEqualTo(value int) bool {
	return s != nil && s[0][0] == value && s[1][1] == value && s[2][2] == value
}

func (s Game) isAntidiagonalEqualTo(value int) bool {
	return s != nil && s[0][2] == value && s[1][1] == value && s[2][0] == value
}

func (s Game) isAnyRowFilledWith(value int) bool {
	for i := 0; i < len(s); i++ {
		if s.isRowEqualTo(i, value) {
			return true
		}
	}
	return false
}

func (s Game) isRowEqualTo(rowIndex, value int) bool {
	return s != nil && s[rowIndex][0] == value && s[rowIndex][1] == value && s[rowIndex][2] == value
}

func (g Game) isAnyColumnFilledWith(value int) bool {
	for i := 0; i < len(g); i++ {
		if g.isColumnEqualTo(i, value) {
			return true
		}
	}
	return false
}

func (g Game) isColumnEqualTo(columnIndex, value int) bool {
	for i := 0; i < len(g); i++ {
		if g[i][columnIndex] != value {
			return false
		}
	}
	return true
}
