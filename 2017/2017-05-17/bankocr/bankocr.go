package bankocr

type typeDigit [4][3]byte

func DecodeDigit(digit typeDigit) int {
	if digit[1][0] == '|' && digit[2][0] == '|' && digit[1][1] == ' ' {
		return 0
	}
	if digit[0][1] == ' ' && digit[1][1] == ' ' {
		return 1
	}
	if digit[2][2] == ' ' {
		return 2
	}
	if digit[1][0] == ' ' && digit[1][1] == '_' {
		return 3
	}
	if digit[2][1] == ' ' && digit[0][1] == ' ' {
		return 4
	}
	if digit[2][0] == ' ' && digit[1][2] == ' ' {
		return 5
	}
	if digit[1][0] == '|' && digit[1][2] == ' ' {
		return 6
	}
	if digit[1][0] == ' ' {
		return 7
	}
	if digit[2][0] == '|' {
		return 8
	}
	return 9
}
