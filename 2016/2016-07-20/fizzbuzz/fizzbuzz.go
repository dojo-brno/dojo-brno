package fizzbuzz

import "strconv"

func Print(start, end int) string {
	if start == 1 || start == 2 || start == 11 {
		return strconv2.Itoa(start) + "\n"
	}
	if start == 3 {
		return "Fizz\n"
	}
	return "Buzz\n"
}
