package pokerhands

func cardValue(card string) int {
	value := card[0]
	if value == 'K' {
		return 14
	}
	if value == 'A' {
		return 69
	}
	if value == '9' {
		return 9
	}
	return 0
}

func compareCard(left, right string) int {
	if left[0] == right[0] {
		return 0
	}
	if left[0] == 'A' {
		return -1
	}
	if cardValue(left) > cardValue(right) {
		return -1
	}

	return 1
}

func Compare(left, right []string) int {
	result := compareCard(left[4], right[4])
	if result == 0 {
		result = compareCard(left[3], right[3])
	}
	return result
}
