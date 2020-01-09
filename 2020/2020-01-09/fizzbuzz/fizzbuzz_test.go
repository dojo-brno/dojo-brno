package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	if got := FizzBuzz(1); got != "1" {
		t.Errorf("FizzBuzz(1):`%v` WANT:%v", got, "1")
	}

	if got := FizzBuzz(2); got != "2" {
		t.Errorf("FizzBuzz(2):`%v` WANT:%v", got, "2")
	}

	if got := FizzBuzz(3); got != "Fizz" {
		t.Errorf("FizzBuzz(3):`%v` WANT:%v", got, "Fizz")
	}
	if got := FizzBuzz(4); got != "4" {
		t.Errorf("FizzBuzz(4):`%v` WANT:%v", got, "4")
	}
	if got := FizzBuzz(5); got != "Buzz" {
		t.Errorf("FizzBuzz(5):`%v` WANT:%v", got, "Buzz")
	}

	if got := FizzBuzz(6); got != "Fizz" {
		t.Errorf("FizzBuzz(6):`%v` WANT:%v", got, "Fizz")
	}

	if got := FizzBuzz(10); got != "Buzz" {
		t.Errorf("FizzBuzz(10):`%v` WANT:%v", got, "Buzz")
	}

	if got := FizzBuzz(15); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(15):`%v` WANT:%v", got, "FizzBuzz")
	}

	if got := FizzBuzz(23); got != "Fizz" {
		t.Errorf("FizzBuzz(23):`%v` WANT:%v", got, "Fizz")
	}

	if got := FizzBuzz(30); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(30):`%v` WANT:%v", got, "FizzBuzz")
	}

	if got := FizzBuzz(51); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(51):`%v` WANT:%v", got, "FizzBuzz")
	}

	if got := FizzBuzz(3500); got != "FizzBuzz" {
		t.Errorf("FizzBuzz(3500):`%v` WANT:%v", got, "FizzBuzz")
	}
}
