package fizzbuzz

import (
	"strconv"
	"testing"
)

func fizzbuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{99, "Fizz"},
		{100, "Buzz"},
	}
	for _, tt := range tests {
		if got := fizzbuzz(tt.in); got != tt.out {
			t.Errorf("fizzbuzz(%d) == %q, want %q", tt.in, got, tt.out)
		}
	}
}
