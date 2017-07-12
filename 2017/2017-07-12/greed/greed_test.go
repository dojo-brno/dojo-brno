package greed

import "testing"

func TestGreed(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{},
			want: 0,
		},
		{
			in:   []int{1},
			want: 100,
		},
		{
			in:   []int{2},
			want: 0,
		},
		{
			in:   []int{5},
			want: 50,
		},
		{
			in:   []int{2, 5},
			want: 50,
		},
		{
			in:   []int{1, 1},
			want: 0,
		},
		{
			in:   []int{1, 2, 3, 4, 5, 6},
			want: 1200,
		},
		{
			in:   []int{2, 2, 2, 2, 2, 2},
			want: 1600,
		},
		{
			in:   []int{1, 1, 1, 1, 1, 1},
			want: 8000,
		},
		{
			in:   []int{1, 1, 2, 2, 4, 4},
			want: 800,
		},
		{
			in:   []int{2, 2, 2, 2, 2, 4},
			want: 800,
		},
	}
	for _, tt := range tests {
		if got := Score(tt.in); got != tt.want {
			t.Errorf("Score(%v) = %v, but want: %v", tt.in, got, tt.want)
		}
	}
}
