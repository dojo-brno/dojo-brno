package roman

import "testing"

func ToRoman(number int) string {
	cases := []struct {
		number int
		result string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, c := range cases {
		if number >= c.number {
			return c.result + ToRoman(number-c.number)
		}
	}
	// switch {
	// case number >= 1000:
	// 	s += "M" + ToRoman(number-1000)
	// case number >= 900:
	// 	s += "CM" + ToRoman(number-900)
	// case number >= 500:
	// 	s += "D" + ToRoman(number-500)
	// case number >= 400:
	// 	s += "CD" + ToRoman(number-400)
	// case number >= 100:
	// 	s += "C" + ToRoman(number-100)
	// case number >= 90:
	// 	s = "XC" + ToRoman(number-90)
	// case number >= 50:
	// 	s += "L" + ToRoman(number-50)
	// case number >= 40:
	// 	s = "XL" + ToRoman(number-40)
	// case number >= 10:
	// 	s += "X" + ToRoman(number-10)
	// case number >= 9:
	// 	s = "IX" + ToRoman(number-9)
	// case number >= 5:
	// 	s += "V" + ToRoman(number-5)
	// case number >= 4:
	// 	s = "IV" + ToRoman(number-4)
	// case number >= 1:
	// 	s += "I" + ToRoman(number-1)
	// }
	return ""
}

func Test(t *testing.T) {
	tests := []struct {
		number int
		result string
	}{
		{1, "I"},
		{2, "II"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{500, "D"},
		{400, "CD"},
		{900, "CM"},
		{1000, "M"},
	}
	for _, test := range tests {
		got := ToRoman(test.number)
		if got != test.result {
			t.Errorf("ToRoman(%v) = %v, want %v", test.number, got, test.result)
		}
	}
}
