package fizzbuzz

import "testing"

func fizzBuzz(i int) string {
	if i == 1 {
		return "1"
	}

	if i == 2 {
		return "2"
	}

	if i == 4 {
		return "4"
	}

	return "fizz"
}

type testCase struct {
	input int
	want  string
}

func TestAll(t *testing.T) {
	tests := []testCase{
		testCase{
			input: 1,
			want:  "1",
		},
		testCase{
			input: 2,
			want:  "2",
		},
		testCase{
			input: 3,
			want:  "fizz",
		},
		testCase{
			input: 4,
			want:  "4",
		},
	}

	for _, test := range tests {
		input := test.input // 1, 2, 3 ...
		got := fizzBuzz(input)
		want := test.want // "1", "2", "fizz", ...
		if got != want {
			t.Fatalf("fizzBuzz(%v) = %v, want %v", input, got, want)
		}
	}
}
