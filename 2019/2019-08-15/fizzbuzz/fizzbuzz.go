package fizzbuzz

import "strconv"

// Production Code
func isNotInRange(number int) bool {
	return number < 1 || number >= 100
}

func isFizz(number int) bool {
	return number%3 == 0
}

func Fizzbuzz(number int) string {
	if isNotInRange(number) == true {
		return "0"
	}
	if isFizz(number) {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	value := strconv.Itoa(number)
	return value
}
