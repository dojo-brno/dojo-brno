package romans

import "strings"

func ToRoman(n int) string {
	r1 := strings.NewReplacer(
		strings.Repeat("I", 10), "X",
		strings.Repeat("I", 9), "IX",
		strings.Repeat("I", 5), "V",
		strings.Repeat("I", 4), "IV",
	)
	r2 := strings.NewReplacer(
		strings.Repeat("X", 10), "C",
		strings.Repeat("X", 9), "XC",
		strings.Repeat("X", 5), "L",
		strings.Repeat("X", 4), "XL",
	)
	r3 := strings.NewReplacer(
		strings.Repeat("C", 5), "D",
	)
	return r3.Replace(r2.Replace(r1.Replace(strings.Repeat("I", n))))
}
