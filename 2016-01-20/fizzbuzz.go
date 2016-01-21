package fizzbuzz

import "fmt"

func fizzbuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	}
	if number == 5 {
		return "Buzz"
	}
	return fmt.Sprint(number)
}
