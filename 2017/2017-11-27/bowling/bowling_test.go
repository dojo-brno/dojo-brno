package bowling

import "testing"

type testCase struct {
	throws []int
	want   int
}

func TestScore(t *testing.T) {
	tests := []testCase{
		{
			throws: []int{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			throws: []int{
				1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
		{
			throws: []int{
				3, 1, 2, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 7,
		},
		{
			throws: []int{
				3, 1, 2, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 9, 1, 1, 0},
			want: 19,
		},
		{
			throws: []int{
				3, 1, 2, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 9, 1, 2, 0},
			want: 21,
		},
		{
			throws: []int{
				3, 1, 2, 1, 10, 1, 1, 0, 0,
				9, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: 30,
		},
		{
			throws: []int{
				1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 10, 1, 1},
			want: 13,
		},
	}
	for _, tt := range tests {
		if got := Score(tt.throws); got != tt.want {
			t.Errorf("Score(%v): %v, WANT:%v", tt.throws, got, tt.want)
		}
	}

}
