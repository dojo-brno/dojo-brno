package diamond

type Diamond string

func GetDiamond(letter byte) Diamond {
	numberOfCycles := int(letter - 'A' + 1)
	var upper, bottom Diamond
	for i := 0; i < numberOfCycles; i++ {
		currentRow := getRow(i, numberOfCycles)
		if numberOfCycles != i+1 {
			bottom = currentRow + bottom
		}
		upper += currentRow
	}
	return upper + bottom
}

func getRow(row, numberOfCycles int) Diamond {
	spaces := getFilling(' ', numberOfCycles-1)
	if row == 0 {
		return spaces + Diamond('A') + spaces + "\n"
	}
	outsideSpaces := getFilling(' ', numberOfCycles-row-1)
	insideSpaces := getFilling(' ', row)
	return outsideSpaces + Diamond('A'+row) + insideSpaces + Diamond('A'+row) + outsideSpaces + "\n"
}

func getFilling(letter byte, number int) Diamond {
	fill := ""
	for i := 1; i <= number; i++ {
		fill += string(letter)
	}
	return Diamond(fill)
}
