package romans_test

import (
	"testing"

	"github.com/dojo-brno/dojo-brno/2016-12-07/romans"
)

func TestToIntValidInput(t *testing.T) {
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
		if got, err := romans.ToInt(tt.roman); err != nil || got != tt.want {
			t.Errorf("ToInt(%v) got (%v, %v) want (%v, nil)", tt.roman, got, err, tt.want)
		}
	}
}

func TestToIntInvalidInput(t *testing.T) {
	badInputs := []string{
		"a",
	}
	want := romans.Invalid
	for _, in := range badInputs {
		if _, err := romans.ToInt(in); err != want {
			t.Errorf("ToInt(%v) got (_, %v) want (_, %v)", in, err, want)
		}
	}
}
