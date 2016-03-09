package fizzbuzz

import "testing"

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
	}

	for _, test := range tests {
		got := fizzbuzz(test.in)
		if got != test.out {
			t.Errorf("fizzbuzz(%v) = %v, want %v", test.in, got, test.out)
		}
	}
}
