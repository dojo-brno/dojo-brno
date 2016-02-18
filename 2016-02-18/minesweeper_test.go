package minesweeper

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		in  [][]byte
		out [][]byte
	}{
		{},
		{
			in:  [][]byte{[]byte{'.'}},
			out: [][]byte{[]byte{'0'}},
		},
		{
			in:  [][]byte{[]byte{'*'}},
			out: [][]byte{[]byte{'*'}},
		},
		{
			in:  [][]byte{[]byte{'.', '.'}},
			out: [][]byte{[]byte{'0', '0'}},
		},
		{
			in:  [][]byte{[]byte{'*', '*'}},
			out: [][]byte{[]byte{'*', '*'}},
		},
		{
			in:  [][]byte{[]byte{'.', '*'}},
			out: [][]byte{[]byte{'1', '*'}},
		},
		{
			in:  [][]byte{[]byte{'*', '.'}},
			out: [][]byte{[]byte{'*', '1'}},
		},
	}
	for _, test := range tests {
		answer := solve(test.in)
		if !reflect.DeepEqual(answer, test.out) {
			t.Errorf("solve(%v) = %v, want %v", test.in, answer, test.out)
		}
	}
}
