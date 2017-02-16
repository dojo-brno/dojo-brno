package fizzbuzz

import "strconv"

func CountDigits(x int) int {
	y := strconv.Itoa(x)
	sum := 0
	for _, tt := range y {
		val, _ := strconv.Atoi(string(tt))
		sum += val
	}
	return sum
}

func FizzBuzz(x int) string {
	if CountDigits(x) == 7 {
		return "Rohlik"
	}
	if x%15 == 0 {
		return "fizzbuzz"
	}
	if x%3 == 0 {
		return "fizz"
	}
	if x%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(x)
}
