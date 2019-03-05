package foobarqix

import "testing"
import "strconv"

func FooBarQix(number int) string {
	if number == 0 {
		return "0"
	}

	retval := strconv.Itoa(number)
	var value string
	value = getText(number, 5, "Bar")
	if value != "" {
		retval = value
	}
	value = getText(number, 3, "Foo")
	if value != "" {
		retval = value
	}

	return retval
}

func getText(number, digit int, digitText string) string {
	var retval string
	countDigits := countDigits(number, digit, digitText)
	if number%digit == 0 {
		retval = digitText + countDigits
		return retval
	}
	if countDigits != "" {
		retval = countDigits
	}
	return retval
}

func countDigits(number, digit int, digitText string) string {
	retval := ""
	for number > 0 {
		if number%10 == digit {
			retval += digitText
		}
		number /= 10
	}
	return retval
}

func TestFooBarQix(t *testing.T) {
	type testCase struct {
		number int
		result string
	}

	testCases := []testCase{
		{0, "0"},
		{1, "1"},
		{3, "FooFoo"},
		{6, "Foo"},
		{9, "Foo"},
		{12, "Foo"},
		{33, "FooFooFoo"},
		{333, "FooFooFooFoo"},
		{3333, "FooFooFooFooFoo"},
		{13, "Foo"},
		{5, "BarBar"},
		{55, "BarBarBar"},
		{52, "Bar"},
	}

	for _, tt := range testCases {
		if got := FooBarQix(tt.number); got != tt.result {
			t.Errorf("FooBarQix(%v):%v, WANT: %v", tt.number, got, tt.result)
		}
	}
}
