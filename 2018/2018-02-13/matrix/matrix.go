package matrix

/*
 1 2 3
 4 5 6
 7 8 9
*/

type corner struct {
	row int
	col int
}

func Spiral(m [][]int) []int {
	if len(m) == 0 {
		return []int{}
	}

	out := make([]int, 0)

	rows := len(m)
	cols := len(m[0])

	leftUp := corner{0, 0}
	rightUp := corner{0, cols - 1}
	rightDown := corner{rows - 1, cols - 1}
	leftDown := corner{rows - 1, 0}

	for leftUp.col <= rightUp.col && leftUp.row <= leftDown.row {
		//move right
		for ci := leftUp.col; ci <= rightUp.col; ci++ {
			out = append(out, m[leftUp.row][ci])
		}
		leftUp.row++
		rightUp.row++
		if leftUp.row > leftDown.row {
			break
		}

		//move down
		for ri := rightUp.row; ri <= rightDown.row; ri++ {
			out = append(out, m[ri][rightUp.col])
		}
		rightUp.col--
		rightDown.col--
		if rightUp.col < leftUp.col {
			break
		}

		//move left
		for ci := rightDown.col; ci >= leftDown.col; ci-- {
			out = append(out, m[rightDown.row][ci])
		}
		rightDown.row--
		leftDown.row--

		//move up
		for ri := leftDown.row; ri >= leftUp.row; ri-- {
			out = append(out, m[ri][leftDown.col])
		}
		leftUp.col++
		leftDown.col++
	}

	return out
}
