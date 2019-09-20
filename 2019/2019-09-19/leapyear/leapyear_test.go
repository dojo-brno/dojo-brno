package leapyear

import "testing"

func LeapYear(year int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return year%4 == 0
}

func TestLeapYear(t *testing.T) {
	var testYear int
	var isLeapYear bool

	testYear = 1996
	isLeapYear = true
	if got := LeapYear(testYear); got != isLeapYear {
		t.Errorf("LeapYear(%v):%v, WANT:%v", testYear, got, isLeapYear)
	}

	testYear = 2001
	isLeapYear = false
	if got := LeapYear(testYear); got != isLeapYear {
		t.Errorf("LeapYear(%v):%v, WANT:%v", testYear, got, isLeapYear)
	}

	testYear = 1900
	isLeapYear = false
	if got := LeapYear(testYear); got != isLeapYear {
		t.Errorf("LeapYear(%v):%v, WANT:%v", testYear, got, isLeapYear)
	}

	testYear = 2400
	isLeapYear = true
	if got := LeapYear(testYear); got != isLeapYear {
		t.Errorf("LeapYear(%v):%v, WANT:%v", testYear, got, isLeapYear)
	}

	testYear = 2020
	isLeapYear = true
	if got := LeapYear(testYear); got != isLeapYear {
		t.Errorf("LeapYear(%v):%v, WANT:%v", testYear, got, isLeapYear)
	}
}
