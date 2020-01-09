package fizzbuzz

import (
	"strconv"
	"strings"
)

func isFizz(number int) bool {
	return number%3 == 0
}

func containFizz(number int) string {
	convertedNumber := strconv.Itoa(number)
	retval := ""
	if isFizz(number) {
		retval += "Fizz"
	}

	if !strings.Contains(retval, "Fizz") {
		for _, c := range convertedNumber {
			if c == '3' {
				retval += "Fizz"
			}
		}
	}
	return retval
}

func containBuzz(number int) string {
	convertedNumber := strconv.Itoa(number)
	retval := ""
	if isBuzz(number) {
		retval += "Buzz"
	}

	if !strings.Contains(retval, "Buzz") {
		for _, c := range convertedNumber {
			if c == '5' {
				retval += "Buzz"
			}
		}
	}
	return retval
}

func isBuzz(number int) bool {
	return number%5 == 0
}

func FizzBuzz(number int) string {
	retval := containFizz(number)
	retval += containBuzz(number)

	if retval == "" {
		retval = strconv.Itoa(number)
	}

	return retval
}
