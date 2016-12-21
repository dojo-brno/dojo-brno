package fizzbuzz

import "strconv"

func FizzBuzz(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	}

	if num%3 == 0 {
		return "Fizz"
	}

	if num%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(num)
}
