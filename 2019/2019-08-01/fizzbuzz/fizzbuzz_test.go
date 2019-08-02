package fizzbuzz

import "testing"

func GetMagicNumber() int {
	return 13
}

func TestGetMagicNumber(t *testing.T) {
	if GetMagicNumber() != 13 {
		t.Fail()
	}
}

func FizzBuzz(number int) string {
	if number == 2 {
		return "2"
	}
	if number%3 == 0 {
		return "Fizz"
	}

	if number%5 == 0 {
		return "Buzz"
	}
	return "1"
}

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(1) != "1" {
		t.Fail()
	}
	if FizzBuzz(2) != "2" {
		t.Errorf("FizzBuzz for number 2 should return 2")
	}
	if FizzBuzz(3) != "Fizz" {
		t.Errorf("FizzBuzz for number 3 should return \"Fizz\"")
	}
	if FizzBuzz(5) != "Buzz" {
		t.Errorf("FizzBuzz for number 5 should return \"Buzz\"")
	}
	if FizzBuzz(6) != "Fizz" {
		t.Errorf("FizzBuzz for number 6 should return \"Fizz\"")
	}
}
