package romans

import (
	"testing"
)

func TestRomans2Dec(t *testing.T) {
	tests := []struct {
		test string
		want int
	}{
		{"I", 1},
		{"II", 2},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"a", -1},
	}
	for _, tt := range tests {
		if got := Roman2Dec(tt.test); got != tt.want {
			t.Errorf("Roman2Dec(%v) got (%v) want (%v)", tt.test, got, tt.want)
		}
	}
}
