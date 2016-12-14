package romans

import (
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		roman string
		want  int
	}{
		{"I", 1},
		{"II", 2},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"III", 3},
		{"IV", 4},
		{"VI", 6},
		{"CCXI", 211},
		{"CCIX", 209},
		{"MCDXLVI", 1446},
		{"MIM", 1999},
		{"MMXIV", 2014},
	}
	for _, tt := range tests {
		got, err := ToInt(tt.roman)
		if err != nil {
			t.Errorf("ToInt(%q): err = %v, want nil", tt.roman, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ToInt(%v) = (%v) want (%v)", tt.roman, got, tt.want)
		}
	}
}

func TestToIntBadInput(t *testing.T) {
	badInput := []string{
		"",
		"a",
	}
	for _, in := range badInput {
		_, err := ToInt(in)
		if err != ErrInvalid {
			t.Errorf("ToInt(%q): err = %v, want %v", in, err, ErrInvalid)
		}
	}
}
