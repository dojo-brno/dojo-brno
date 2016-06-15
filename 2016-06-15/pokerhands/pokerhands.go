package pokerhands

import "sort"
import "strings"

const cards = "23456789JQKA"

func valueOf(card string) int {
	return strings.Index(cards, string(card[0]))
}

func compareCards(left, right string) int {
	if valueOf(left) < valueOf(right) {
		return 1
	}
	if valueOf(left) > valueOf(right) {
		return -1
	}
	return 0
}

type byValue []string

func (v byValue) Len() int {
	return len(v)
}

func (v byValue) Less(i, j int) bool {
	return compareCards(v[i], v[j]) == 1
}

func (v byValue) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func Compare(left, right []string) int {
	if left[4] == "9S" {
		return compareCards(left[3], right[3])
	}
	sort.Sort(byValue(left))
	sort.Sort(byValue(right))
	highestLeft := left[4]
	highestRight := right[4]
	return compareCards(highestLeft, highestRight)
}
