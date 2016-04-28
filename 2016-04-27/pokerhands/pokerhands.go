package pokerhands

import "strings"

type Hand [5]string

func Compare(left, right Hand) int {
	highestLeft := highest(left)
	highestRight := highest(right)
	return strings.Compare(highestLeft, highestRight)
}

func highest(h Hand) string {
	n := h[0][:1]
	for _, x := range h[1:] {
		if x[:1] > n {
			n = x[:1]
		}
	}
	return n
}
