package romans

import (
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		roman string
		want  int
	}{
		// all basic numerals
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},

		// additive repetitions of a numeral
		{"MMMCCCXXXIII", 3333},

		// additive case of different numerals
		{"MDCLXVI", 1666},

		// subtractive case
		{"CDXLIV", 444},
		{"CMXCIX", 999},
	}
	for _, tt := range tests {
		got, err := ToInt(tt.roman)
		if err != nil {
			t.Errorf("ToInt(%q): err = %v, want nil", tt.roman, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ToInt(%q) = %v, want %v", tt.roman, got, tt.want)
		}
	}
}

func TestToIntBadInput(t *testing.T) {
	badInput := []string{
		"",
		"a",
		"MIM",
	}
	for _, in := range badInput {
		_, err := ToInt(in)
		if err != ErrInvalid {
			t.Errorf("ToInt(%q): err = %v, want %v", in, err, ErrInvalid)
		}
	}
}
