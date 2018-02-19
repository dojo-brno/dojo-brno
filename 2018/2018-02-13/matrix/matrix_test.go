package matrix

import (
	"reflect"
	"testing"
)

func TestSpiral(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want []int
	}{
		{
			name: "empty",
			in:   [][]int{},
			want: []int{},
		},
		{
			name: "one number",
			in:   [][]int{{0}},
			want: []int{0},
		},
		{
			name: "2 rows",
			in:   [][]int{{0}, {1}},
			want: []int{0, 1},
		},
		{
			name: "2 rows 2 cols",
			in:   [][]int{{0, 1}, {2, 3}},
			want: []int{0, 1, 3, 2},
		},
		{
			name: "board example",
			in:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "board example 4x3",
			in:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			want: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			name: "board example 3x4",
			in:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "one row",
			in:   [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "3 rows",
			in:   [][]int{{0}, {1}, {2}},
			want: []int{0, 1, 2},
		},
		{
			name: "6x5",
			in: [][]int{{0, 1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10, 11},
				{12, 13, 14, 15, 16, 17},
				{18, 19, 20, 21, 22, 23},
				{24, 25, 26, 27, 28, 29}},
			want: []int{0, 1, 2, 3, 4, 5, 11, 17, 23, 29, 28, 27, 26, 25, 24, 18, 12, 6, 7, 8, 9, 10, 16, 22, 21, 20, 19, 13, 14, 15},
		},
		{
			name: "5x5",
			in: [][]int{{0, 1, 2, 3, 4},
				{5, 6, 7, 8, 9},
				{10, 11, 12, 13, 14},
				{15, 16, 17, 18, 19},
				{20, 21, 22, 23, 24}},
			want: []int{0, 1, 2, 3, 4, 9, 14, 19, 24, 23, 22, 21, 20, 15, 10, 5, 6, 7, 8, 13, 18, 17, 16, 11, 12},
		},
		{
			name: "3x5",
			in: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
				{9, 10, 11},
				{12, 13, 14}},
			want: []int{0, 1, 2, 5, 8, 11, 14, 13, 12, 9, 6, 3, 4, 7, 10},
		},
	}
	for _, tt := range tests {
		get := Spiral(tt.in)
		if !reflect.DeepEqual(get, tt.want) {
			t.Errorf("Error in %s,\nwant: %v,\ngot:  %v", tt.name, tt.want, get)
		}
	}
}
