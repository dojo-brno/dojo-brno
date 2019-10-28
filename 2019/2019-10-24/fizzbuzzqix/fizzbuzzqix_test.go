package fizzbuzzqix

import "testing"
import "strconv"

func FizzBuzzQix(x int) string {
	result := ""
	if isFizz(x) {
		result += "Fizz"
	}
	if isBuzz(x) {
		result += "Buzz"
	}
	if isQix(x) {
		result += "Qix"
	}
	if result == "" {
		result = strconv.Itoa(x)
	}
	return result
}

func isFizz(x int) bool {
	return x%3 == 0
}
func isBuzz(x int) bool {
	return x%5 == 0
}
func isQix(x int) bool {
	return x%7 == 0
}

func TestFizzBuzzQix(t *testing.T) {
	type testCase struct {
		testNumber     int
		expectedNumber string
	}
	testCases := []testCase{
		{1, "1"},
		{2, "2"},
		{3, "FizzFizz"},
		{6, "Fizz"},
		{9, "Fizz"},
		{5, "Buzz"},
		{7, "Qix"},
		{15, "FizzBuzz"},
	}
	for _, tt := range testCases {
		if got := FizzBuzzQix(tt.testNumber); got != tt.expectedNumber {
			t.Errorf("FizzBuzzQix(%v): %v, Expected: %v", tt.testNumber, got, tt.expectedNumber)
		}
	}
}
