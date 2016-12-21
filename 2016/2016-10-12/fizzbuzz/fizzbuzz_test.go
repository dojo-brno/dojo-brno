package fizzbuzz

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{15, "FizzBuzz"},
	}
	for i := 0; i < len(tests); i++ {
		in := tests[i].in
		want := tests[i].want
		got := FizzBuzz(in)
		if got != want {
			t.Errorf("FizzBuzz(%v) = %v, want %v", in, got, want)
		}
	}
}
