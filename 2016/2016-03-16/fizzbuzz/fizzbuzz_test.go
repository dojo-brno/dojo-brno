package fizzbuzz

import "fmt"
import "testing"

func fizzbuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(number)
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{4, "4"},
		{3, "Fizz"},
		{6, "Fizz"},
		{9, "Fizz"},
		{5, "Buzz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{60, "FizzBuzz"},
	}
	for _, tt := range tests {
		if got := fizzbuzz(tt.in); got != tt.out {
			t.Errorf("fizzbuzz(%v) = %q, want %q", tt.in, got, tt.out)
		}
	}
}
