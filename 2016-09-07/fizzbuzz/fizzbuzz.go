package fizzbuzz

import "strconv"

func FizzBuzz(i int) string {
	if i == 3 {
		return "fizz"
	}
	if i == 5 {
		return "buzz"
	}
	return strconv.Itoa(i)
}
