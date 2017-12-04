package bowling

import (
	"reflect"
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		testName string
		throws   []int
		want     int
	}{
		{
			testName: "Zero",
			throws:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     0,
		},
		{
			testName: "One",
			throws:   []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     1,
		},
	}
	for _, tt := range tests {
		if got := Score(tt.throws); got != tt.want {
			t.Errorf("%v: Score(%v) = %v WANT %v", tt.testName, tt.throws, got, tt.want)
		}
	}
}

func TestBuildFrame2(t *testing.T) {
	tests := []struct {
		testName   string
		throws     []int
		wantFrame  frame2
		wantThrows []int
	}{
		{
			testName:   "Zero",
			throws:     []int{0, 0},
			wantFrame:  frame2{0, 0, false, false},
			wantThrows: []int{},
		},
		{
			testName:   "Non-zero values",
			throws:     []int{3, 1},
			wantFrame:  frame2{3, 1, false, false},
			wantThrows: []int{},
		},
		{
			testName:   "Strike",
			throws:     []int{10, 5},
			wantFrame:  frame2{10, 0, true, false},
			wantThrows: []int{5},
		},
		{
			testName:   "three throws",
			throws:     []int{0, 1, 2},
			wantFrame:  frame2{0, 1, false, false},
			wantThrows: []int{2},
		},
		{
			testName:   "Spare",
			throws:     []int{5, 5},
			wantFrame:  frame2{5, 5, false, true},
			wantThrows: []int{},
		},
	}
	for _, tt := range tests {
		if gotF, gotT := buildFrame2(tt.throws); gotF != tt.wantFrame || !reflect.DeepEqual(gotT, tt.wantThrows) {
			t.Errorf("%v: buildFrame2(%v) = (%v,%v) WANT (%v,%v)", tt.testName, tt.throws, gotF, gotT, tt.wantFrame, tt.wantThrows)
		}
	}
}
