package fizzbuzz

import "testing"

func TestFizzBuzzTable(t *testing.T) {
	tests := []struct {
		start, end int
		want       string
	}{
		{
			start: 1,
			end:   1,
			want:  "1\n",
		},
		{
			start: 3,
			end:   3,
			want:  "Fizz\n",
		},
		{
			start: 5,
			end:   5,
			want:  "Buzz\n",
		},
		{2, 2, "2\n"},
		{10, 10, "Buzz\n"},
		{11, 11, "11\n"},
	}
	for _, tt := range tests {
		output := Print(tt.start, tt.end)
		if output != tt.want {
			t.Errorf("Print(%v, %v) = %q, want %q", tt.start, tt.end, output, tt.want)
		}
	}
}
