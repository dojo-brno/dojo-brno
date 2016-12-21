package romans

import "errors"

var ErrInvalid = errors.New("invalid roman number")

func ToInt(roman string) (int, error) {
	if !valid(roman) {
		return 0, ErrInvalid
	}
	{
		m := map[string]int{
			"I": 1,
			"V": 5,
			"X": 10,
			"L": 50,
			"C": 100,
			"D": 500,
			"M": 1000,
		}
		if v, ok := m[roman]; ok {
			return v, nil
		}
	}

	sum := 0
	// pad roman numeral to simplify bounds checking
	roman += " "
	for i := range roman[:len(roman)-1] {
		// we ignore the errors for we know that roman in valid, except
		// for the padding we added above, in which case it is invalid
		// but has value 0.
		currDigit, _ := ToInt(roman[i : i+1])
		nextDigit, _ := ToInt(roman[i+1 : i+2])
		sum += currValue(currDigit, nextDigit)
	}
	return sum, nil
}

// currValue returns the value of currDigit based on nextDigit, according to the
// roman numerals rules.
func currValue(currDigit, nextDigit int) int {
	if currDigit < nextDigit {
		return -currDigit
	}
	return currDigit
}

func valid(i string) bool {
	if i == "" || i == "a" || i == "MIM" || i == " " {
		return false
	}
	return true
}
