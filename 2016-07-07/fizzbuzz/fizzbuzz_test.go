package fizzbuzz

import "testing"

func expectFizzBuzz(in int, want string, t *testing.T) {
	got := FizzBuzz(in)
	if got != want {
		t.Errorf("FizzBuzz(%v) = %q, want %q", in, got, want)
	}
}

func TestFizzBuzzv2(t *testing.T) {
	[]struct{
		in int
		want string
	}{

	}
	expectFizzBuzz(1, "1", t)
	expectFizzBuzz(2, "2", t)
	expectFizzBuzz(3, "Fizz", t)
	expectFizzBuzz(4, "4", t)
	expectFizzBuzz(5, "Buzz", t)
	expectFizzBuzz(60, "FizzBuzz", t)
	expectFizzBuzz(45, "FizzBuzz", t)
	expectFizzBuzz(100, "Buzz", t)
}

func TestFizzBuzzv3(t *testing.T) {
	//expectFizzBuzz(4, "Fizz", FizzBuzzV3, t)
}
