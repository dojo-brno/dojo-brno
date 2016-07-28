package romans

import "testing"

func TestRomanConversion(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},

		{"II", 2},
		{"IV", 4},
	}
	for _, testCase := range tests {
		got := Roman2Arabic(testCase.input)
		if got != testCase.want {
			t.Errorf("Roman2Arabic(%v) = %v, want %v", testCase.input, got, testCase.want)
		}
	}

}
