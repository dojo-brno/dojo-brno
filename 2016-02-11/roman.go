package roman

import "strings"

func ToRoman(number int) string {
	// Weird cases
	if number == 4 {
		return "IV"
	}
	if number == 5 {
		return "V"
	}
	// Other cases
	var letter string
	if number < 10 {
		letter = "I"
	} else if number < 100 {
		letter = "X"
		number /= 10
	} else if number < 1000 {
		letter = "C"
		number /= 100
	} else {
		letter = "M"
		number /= 1000
	}
	return strings.Repeat(letter, number)
}
