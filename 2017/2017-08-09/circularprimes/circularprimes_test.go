package circularprimes

import "testing"

func TestIsCircularPrime(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{6, false},
		{7, true},
		{9, false},
		{25, false},
	}
	for _, tt := range tests {
		if got := IsCircularPrime(tt.input); got != tt.want {
			t.Errorf("IsCircularPrime(%v) = %v, want %v",
				tt.input, got, tt.want)
		}
	}
}
