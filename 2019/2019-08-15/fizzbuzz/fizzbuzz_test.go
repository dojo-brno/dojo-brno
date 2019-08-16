package fizzbuzz

import "testing"


//Test Cases
func TestFizzbuzz(t *testing.T) {
	if Fizzbuzz(3) != "Fizz" {
		t.Errorf("Function Fizzbuzz(3) did not return Fizz!")
	}
	if Fizzbuzz(1) != "1" {
		t.Errorf("Function Fizzbuzz(1) did not return 1!")
	}
	if Fizzbuzz(2) != "2" {
		t.Errorf("Function Fizzbuzz(2) did not return 2!")
	}
	if Fizzbuzz(6) != "Fizz" {
		t.Errorf("Function FizzBuzz for number 6 should return Fizz!")
	}
	if Fizzbuzz(101) != "0" {
		t.Errorf("Please choose integer from 1 to 100!")
	}
	if Fizzbuzz(102) != "0" {
		t.Errorf("Please choose integer from 1 to 100!")
	}
	if Fizzbuzz(-1) != "0" {
		t.Errorf("Please choose integer from 1 to 100!")
	}
	if Fizzbuzz(9) != "Fizz" {
		t.Errorf("Function FizzBuzz for number 9 should return Fizz!")
	}

	var number int
	var want string

	number = 5
	want = "Buzz"
	if got := Fizzbuzz(number); got != want {
		t.Errorf("FizzBuzz(%v):%v, WANT:%v!", number, got, want)
	}

	number = 10
	want = "Buzz"
	if got := Fizzbuzz(number); got != want {
		t.Errorf("FizzBuzz(%v):%v, WANT:%v!", number, got, want)
	}
}
