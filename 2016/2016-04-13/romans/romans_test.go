package romans_test

import (
	"testing"

	"github.com/dojo-brno/dojo-brno/2016/2016-04-13/romans"
)

func TestRomans(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{1, "I"},
		{2, "II"},
		{5, "V"},
		{10, "X"},
		{15, "XV"},
		{6, "VI"},
		{4, "IV"},
		{9, "IX"},
		{20, "XX"},
		{50, "L"},
		{100, "C"},
		{150, "CL"},
		{60, "LX"},
		{40, "XL"},
		{90, "XC"},
		{200, "CC"},
		{500, "D"},
	}
	for _, tt := range tests {
		if got := romans.ToRoman(tt.in); got != tt.want {
			t.Errorf("ToRoman(%v) = '%v', want '%v'", tt.in, got, tt.want)
		}
	}

}
