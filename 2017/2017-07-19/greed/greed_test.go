package greed

import "testing"

func TestScore(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{},
			want: 0,
		},
		{
			in:   []int{1, 2, 3, 4, 5, 6},
			want: 1200,
		},
		{
			in:   []int{1, 1, 2, 2, 3, 3},
			want: 800,
		},
		{
			in:   []int{1, 2, 3, 1, 2, 3},
			want: 800,
		},
		{
			in:   []int{1, 2, 3, 4, 3, 6},
			want: 100,
		},
		{
			in:   []int{1},
			want: 100,
		},
		{
			in:   []int{5},
			want: 50,
		},
		{
			in:   []int{1, 1, 1},
			want: 1000,
		},
		{
			in:   []int{3, 5, 6},
			want: 50,
		},
		{
			in:   []int{1, 3, 5, 6},
			want: 150,
		},
		{
			in:   []int{3, 3, 3},
			want: 300,
		},
	}
	for _, tt := range tests {
		if got := Score(tt.in...); got != tt.want {
			t.Errorf("Score(%v)==%v but we want %v", tt.in, got, tt.want)
		}
	}
}
