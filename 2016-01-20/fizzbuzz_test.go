package fizzbuzz

import "testing"

type testCase struct {
	in   int
	want string
}

var fizzBuzzTests = []testCase{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, testCase := range fizzBuzzTests {
		if got := fizzbuzz(testCase.in); got != testCase.want {
			t.Errorf("fizzbuzz(%v) = %v; want %v", testCase.in, got, testCase.want)
		}
	}
}
