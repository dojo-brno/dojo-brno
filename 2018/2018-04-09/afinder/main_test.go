package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestAfinder(t *testing.T) {
	tests := []struct {
		arr  []int
		x    int
		want [][]int
	}{
		{
			[]int{1, 2, 3},
			4,
			[][]int{{1, 3}},
		},
		{
			[]int{1, 2, 3},
			5,
			[][]int{{2, 3}},
		},
		{
			[]int{1, 2, 3},
			3,
			[][]int{{1, 2}},
		},
		{
			[]int{1, 2, 3, 4, -1},
			3,
			[][]int{{1, 2}, {4, -1}},
		},
		{
			[]int{1, 2, 3, 4, -1},
			33,
			[][]int{},
		},
		{
			[]int{1, 2, 3, 4, -1, 15, 16, 17, 18, 19, 20, 21, 22},
			3,
			[][]int{{1, 2}, {4, -1}},
		},
		{
			[]int{2, 2, 2},
			4,
			[][]int{{2, 2}, {2, 2}, {2, 2}},
		},
	}

	for _, test := range tests {
		got := AfinderEfficient(test.arr, test.x)

		if !equal(test.want, got) {
			t.Errorf("AfinderEfficient(%v, %v): %v, WANT: %v", test.arr, test.x, got, test.want)
		}
	}
}

//[[1,2],[4,-1]]
//[[-1,4],[2,1]]
func equal(a [][]int, b [][]int) bool {
	for _, v := range a {
		sort.Ints(v)
	}
	for _, v := range b {
		sort.Ints(v)
	}

	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	sort.Slice(b, func(i, j int) bool { return b[i][0] < b[j][0] })

	return reflect.DeepEqual(a, b)
}
