package search

func SearchSorted(n int, alon []int) bool {
	if len(alon) == 0 {
		return false
	}
	switch m := alon[len(alon)/2]; {
	case m == n:
		return true
	case m > n:
		return SearchSorted(n, alon[:len(alon)/2])
	default:
		return SearchSorted(n, alon[len(alon)/2+1:])
	}
}
