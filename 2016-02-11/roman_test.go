package roman

import "testing"

type testCase struct {
	in   int
	want string
}

func TestRoman(t *testing.T) {
	tests := []testCase{
		// Single, repeatable digits
		{1, "I"},
		{10, "X"},
		{100, "C"},
		{1000, "M"},
		// Single, non-repeatable digits
		{5, "V"},
		// Repetitions of single digits
		{2, "II"},
		{3, "III"},
		{20, "XX"},
		{200, "CC"},
		{2000, "MM"},
		// Weird cases
		{4, "IV"},
	}
	for _, tt := range tests {
		if got := ToRoman(tt.in); got != tt.want {
			t.Errorf("ToRoman(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
