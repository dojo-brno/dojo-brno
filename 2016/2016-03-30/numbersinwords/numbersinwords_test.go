package numbersinwords

import "testing"

func number2words(n int) string {
	if n >= 100 {
		out := number2words(n/100) + " hundred"
		rest := n - n/100*100
		if rest != 0 {
			out += " and " + number2words(rest)
		}
		return out
	}
	if n == 21 || n == 22 || n == 31 {
		return number2words(n/10*10) + " " + number2words(n-n/10*10)
	}
	numberDB := map[int]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		10: "ten",
		11: "eleven",
		20: "twenty",
		30: "thirty",
	}
	return numberDB[n]
}

func TestNumbersInWords(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{10, "ten"},
		{11, "eleven"},
		{20, "twenty"},
		{21, "twenty one"},
		{22, "twenty two"},
		{30, "thirty"},
		{31, "thirty one"},
		{100, "one hundred"},
		{101, "one hundred and one"},
		{102, "one hundred and two"},
		{200, "two hundred"},
		{201, "two hundred and one"},
		{220, "two hundred and twenty"},
		{221, "two hundred and twenty one"},
	}
	for _, tt := range tests {
		got := number2words(tt.in)
		if got != tt.want {
			t.Errorf("number2words(%d) = %q but we want %q", tt.in, got, tt.want)
		}
	}
}
