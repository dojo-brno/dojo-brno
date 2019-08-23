package fizzbuzz

import "strconv"
import "strings"

func FizzBuzz(number int) string {
	if isFizz(number) && isBuzz(number) {
		return "FizzBuzz"
	}
	if isFizz(number) {
		return "Fizz"
	}
	if isBuzz(number) {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func isFizz(number int) bool {
	return number%3 == 0 || isContaining(number, "3")
}

func isBuzz(number int) bool {
	return number%5 == 0 || isContaining(number, "5")
}

func isContaining(number int, containingNumber string) bool {
	return strings.Contains(strconv.Itoa(number), containingNumber)

	// for _, c := range strconv.Itoa(number) {
	// 	if c == containingNumber {
	// 		return true
	// 	}
	// }
	// return false
}
