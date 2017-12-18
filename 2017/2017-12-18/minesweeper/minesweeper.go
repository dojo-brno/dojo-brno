package minesweeper

import "strconv"

type MineField struct {
	Rows, Columns int
	Field         [][]string
}

const (
	emptyField = "."
	mine       = "*"
)

func Evaluate(mineField MineField) MineField {
	if !isValid(mineField) {
		return mineField
	}
	for i, row := range mineField.Field {
		for j, field := range row {
			if field == emptyField {
				mineField.Field[i][j] = sumUpMines(mineField, i, j)
			}
		}
	}
	return mineField
}

func sumUpMines(mineField MineField, row, column int) string {
	dir := []struct {
		x, y int
	}{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	sum := 0
	for j, _ := range dir {
		tempx := row + dir[j].x
		tempy := column + dir[j].y
		if isItIn(mineField, tempx, tempy) && mineField.Field[tempx][tempy] == mine {
			sum++
		}

	}
	return strconv.Itoa(sum)
}

func isItIn(mineField MineField, row, column int) bool {
	return !(row < 0 || column < 0 || row >= mineField.Rows || column >= mineField.Columns)
}

func isValid(mineField MineField) bool {
	if mineField.Rows != len(mineField.Field) {
		return false
	}
	for _, row := range mineField.Field {
		if mineField.Columns != len(row) {
			return false
		}
	}
	return true
}
