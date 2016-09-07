package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{5, "buzz"},
	}
	for _, tt := range tests {
		got := FizzBuzz(tt.in)
		if got != tt.want {
			t.Errorf("FizzBuzz(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
