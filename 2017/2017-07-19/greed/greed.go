package greed

import "sort"

var rules = []func(...int) (int, bool, bool){
	straight,
	threePairs,
	tripleOnes,
	ordinaryTriple,
	singleOne,
	singleFive,
}

func Score(values ...int) int {
	sort.Ints(values)
	total := 0
	for _, rule := range rules {
		score, overrule, ok := rule(values...)
		if ok {
			total += score
			if overrule {
				return score
			}
		}
	}
	return total
}

func straight(values ...int) (int, bool, bool) {
	return 1200, true, len(values) == 6 && values[0] == 1 && values[1] == 2 && values[2] == 3 && values[3] == 4 && values[4] == 5 && values[5] == 6
}

func threePairs(values ...int) (int, bool, bool) {
	return 800, true, len(values) == 6 && values[0] == values[1] && values[1] != values[2] && values[2] == values[3] && values[3] != values[4] && values[4] == values[5]
}

func tripleOnes(values ...int) (int, bool, bool) {
	return 1000, false, len(values) == 3 && values[0] == 1 && values[1] == 1 && values[2] == 1
}

func ordinaryTriple(values ...int) (int, bool, bool) {
	if len(values) < 3 {
		return 0, false, false
	}
	count = countNumbers(3)
	return 300, false, count == 3

}

func count(number int, values ...int) int {

}

func singleOne(values ...int) (int, bool, bool) {
	count := 0
	for _, value := range values {
		if value == 1 {
			count++
		}
	}
	return 100, false, count == 1
}

func singleFive(values ...int) (int, bool, bool) {
	count := 0
	for _, value := range values {
		if value == 5 {
			count++
		}
	}
	return 50, false, count == 1
}
