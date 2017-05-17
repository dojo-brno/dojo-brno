package bankocr

import "testing"

func TestDecodeDigits(t *testing.T) {
	tests := []struct {
		testName string
		digit    typeDigit
		want     int
	}{
		{
			testName: "Digit 0:",
			digit:    typeDigit{{' ', '_', ' '}, {'|', ' ', '|'}, {'|', '_', '|'}, {' ', ' ', ' '}},
			want:     0,
		},
		{
			testName: "Digit 1:",
			digit:    typeDigit{{' ', ' ', ' '}, {' ', ' ', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}},
			want:     1,
		},
		{
			testName: "Digit 2:",
			digit:    typeDigit{{' ', '_', ' '}, {' ', '_', '|'}, {'|', '_', ' '}, {' ', ' ', ' '}},
			want:     2,
		},
		{
			testName: "Digit 3:",
			digit:    typeDigit{{' ', '_', ' '}, {' ', '_', '|'}, {' ', '_', '|'}, {' ', ' ', ' '}},
			want:     3,
		},
		{
			testName: "Digit 4:",
			digit:    typeDigit{{' ', ' ', ' '}, {'|', '_', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}},
			want:     4,
		},
		{
			testName: "Digit 5:",
			digit:    typeDigit{{' ', '_', ' '}, {'|', '_', ' '}, {' ', '_', '|'}, {' ', ' ', ' '}},
			want:     5,
		},
		{
			testName: "Digit 6:",
			digit:    typeDigit{{' ', '_', ' '}, {'|', '_', ' '}, {'|', '_', '|'}, {' ', ' ', ' '}},
			want:     6,
		},
		{
			testName: "Digit 7:",
			digit:    typeDigit{{' ', '_', ' '}, {' ', ' ', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}},
			want:     7,
		},
		{
			testName: "Digit 8:",
			digit:    typeDigit{{' ', '_', ' '}, {'|', '_', '|'}, {'|', '_', '|'}, {' ', ' ', ' '}},
			want:     8,
		},
		{
			testName: "Digit 9:",
			digit:    typeDigit{{' ', '_', ' '}, {'|', '_', '|'}, {' ', '_', '|'}, {' ', ' ', ' '}},
			want:     9,
		},
	}
	for _, tt := range tests {
		if got := DecodeDigit(tt.digit); got != tt.want {
			t.Errorf("%v DecodeDigit(%v) = %v want %v", tt.testName, tt.digit, got, tt.want)
		}
	}
}
