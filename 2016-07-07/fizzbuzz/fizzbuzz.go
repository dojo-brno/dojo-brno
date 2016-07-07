package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
        // if n == 4 {
        //         return "Fizz"
        // }
        if n%15 == 0 {
                return "FizzBuzz"
        }
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
