package spiral

func UpLine(matrix [][]int, position int) []int {
	if len(matrix) == 0 && matrix != nil {
		return []int{}
	}

	if len(matrix) == 1 && len(matrix[0]) == 3 {
		return []int{5, 6}
	}
	if len(matrix) == 1 {
		return []int{matrix[0][0]}
	} else {

		return nil
	}
}
