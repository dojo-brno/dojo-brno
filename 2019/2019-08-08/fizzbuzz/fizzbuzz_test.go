package fizzbuzz

import "testing"
import "strconv"

func Fizzbuzz(number int) string {
	if number == 0 {
		return "0"
	}
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	}

	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func TestFizzBuzz(t *testing.T) {
	type testCase = struct {
		number   int
		expected string
	}

	results := testCase{1, "1"}
}

func TestFizzbuzzOne(t *testing.T) {
	if Fizzbuzz(1) != "1" {
		t.Fail()
	}
}

func TestFizzbuzzThree(t *testing.T) {
	if Fizzbuzz(3) != "Fizz" {
		t.Fail()
	}
}

func TestFizzbuzzTwo(t *testing.T) {
	expected := "2"
	if got := Fizzbuzz(2); got != expected {
		t.Errorf("Fizzbuzz(2):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzMultiplyThree(t *testing.T) {
	expected := "Fizz"
	if got := Fizzbuzz(6); got != expected {
		t.Errorf("Fizzbuzz(6):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzFive(t *testing.T) {
	expected := "Buzz"
	if got := Fizzbuzz(5); got != expected {
		t.Errorf("Fizzbuzz(5):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzTen(t *testing.T) {
	expected := "Buzz"
	if got := Fizzbuzz(10); got != expected {
		t.Errorf("Fizzbuzz(10):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzFifteen(t *testing.T) {
	expected := "FizzBuzz"
	if got := Fizzbuzz(15); got != expected {
		t.Errorf("Fizzbuzz(15):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzThirty(t *testing.T) {
	expected := "FizzBuzz"
	if got := Fizzbuzz(30); got != expected {
		t.Errorf("Fizzbuzz(30):%v, WANT:%v", got, expected)
	}
}

func TestFizzbuzzSixty(t *testing.T) {
	expected := "0"
	if got := Fizzbuzz(0); got != expected {
		t.Errorf("Fizzbuzz(0):%v, WANT:%v", got, expected)
	}
}
