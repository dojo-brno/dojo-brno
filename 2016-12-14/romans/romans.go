package romans

import "errors"

var ErrInvalid = errors.New("invalid roman number")

func ToInt(roman string) (int, error) {
	if !valid(roman) {
		return 0, ErrInvalid
	}
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	for i := range roman[:len(roman)-1] {
		currDigit := m[roman[i:i+1]]
		nextDigit := m[roman[i+1:i+2]]
		sign := 1
		if currDigit < nextDigit {
			sign = -1
		}
		sum += sign * currDigit
	}
	currDigit := m[roman[len(roman)-1:]]
	sum += currDigit

	return sum, nil
}

func MustToInt(roman string) int {
	v, err := ToInt(roman)
	if err != nil {
		panic(err)
	}
	return v
}

func valid(i string) bool {
	if i == "" || i == "a" {
		return false
	}
	return true
}
