package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		number int
		want   string
	}{
		{1, "1"},
		{3, "fizz"},
		{5, "buzz"},
		{14, "14"},
		{15, "fizzbuzz"},
		{30, "fizzbuzz"},
		{7, "Rohlik"},
		{34, "Rohlik"},
		{61, "Rohlik"},
	}
	for _, tt := range tests {
		if got := FizzBuzz(tt.number); got != tt.want {
			t.Errorf("FizzBuzz(%v) = %q want %q", tt.number, got, tt.want)
		}
	}
}

func TestFizzBuzzEnterprise(t *testing.T) {
	tests := []struct {
		rules []Rule
		number int
		want   string
	}{
		{
			[]Rule{
				func (n int) (string, bool) {
					return "nádraží", true
				},
			},
			1,
			"nádraží",
		},
	}
	for _, tt := range tests {
		if got := rules.FizzBuzz(tt.number); got != tt.want {
			t.Errorf("rules.FizzBuzz(%v) = %q want %q", tt.number, got, tt.want)
		}
	}
}
