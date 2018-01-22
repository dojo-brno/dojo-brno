package numberstowords

func NumberToWord(num int) string {
	/*digits := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}*/
	irregular := map[int]string{
		0:    "zero",
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "thousand",
	}100

	if num < 20 {
		if result, ok := irregular[num]; ok {
			return result
		}
		base := 10
		decimals := num / base
		split := num % base
		return irregular[decimals*base] + " " + irregular[split]
	}

	result := ""
	var remainder int
	for base := 1000; base >= 10; base /= 10 {
		digit := num / base
		if digit == 0 {
			continue
		}
		remainder = num - digit*base
		result += irregular[digit] + " " + irregular[base]
	}
	if remainder == 0 {
		return result
	}
	if res, ok := irregular[remainder]; ok {
		return result + " " + res
	}
	base := 10
	decimals := num / base
	split := num % base
	return result + " " + irregular[decimals*base] + " " + irregular[split]
}
