package fizzbuzz

import "fmt"

func fizzbuzz(n int) string {
	var result string
	switch {
	case n%3 == 0:
		result += "Fizz"
		fallthrough
	case n%5 == 0:
		result += "Buzz"
	default:
		result = fmt.Sprint(n)
	}
	return result
}
