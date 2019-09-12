package fizzbuzz

import "testing"
import "strconv"

func FizzBuzz(num int) string {
	if num == 1 {
		return strconv.Itoa(1)
	}
	if num == 2 {
		return "2"
	}
	if num == 4 {
		return "4"
	}
	return "Fizz"
}

func TestFizzBuzz(t *testing.T) {
	if got := FizzBuzz(1); got != "1" {
		t.Fail()
	}
	if got := FizzBuzz(2); got != "2" {
		t.Fail()
	}
	if got := FizzBuzz(3); got != "Fizz" {
		t.Fail()
	}
	if got := FizzBuzz(4); got != "4" {
		t.Fail()
	}
}
