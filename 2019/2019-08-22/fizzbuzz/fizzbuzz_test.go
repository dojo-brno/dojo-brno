package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	type testCase = struct {
		number int
		want   string
	}

	testCases := []testCase{
		//Tests for dividing
		testCase{number: 1, want: "1"},
		testCase{number: 2, want: "2"},
		testCase{number: 3, want: "Fizz"},
		testCase{number: 5, want: "Buzz"},
		testCase{number: 15, want: "FizzBuzz"},

		//Tests for containing
		testCase{number: 13, want: "Fizz"},
		testCase{number: 23, want: "Fizz"},
		testCase{number: 151, want: "Buzz"},
		testCase{number: 153, want: "FizzBuzz"},
	}

	for _, tt := range testCases {
		if got := FizzBuzz(tt.number); got != tt.want {
			t.Errorf("FizzBuzz(%v):%v, WANT:%v", tt.number, got, tt.want)
		}
	}
}
