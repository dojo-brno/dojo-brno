package bowling

import (
	"reflect"
	"testing"
)

func TestScoreError(t *testing.T) {
	tests := []struct {
		testName string
		throws   []int
		want     error
	}{
		{
			testName: "All Zeros in 19 throws",
			throws:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     TooFewThrowsError(19),
		},
		{
			testName: "All Zeros in 0 throws",
			throws:   []int{},
			want:     TooFewThrowsError(0),
		},
	}
	for _, tt := range tests {
		if got, err := Score(tt.throws); err != tt.want {
			t.Errorf("%v: Score(%v) = (%v, %v) WANT: %v", tt.testName, tt.throws, got, err, tt.want)
		}
	}
}

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
		{
			testName: "Sum",
			throws:   []int{0, 1, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8},
			want:     14,
		},
		{
			testName: "Spare",
			throws:   []int{9, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     12,
		},
		{
			testName: "Spare in Last Frame",
			throws:   []int{9, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 4},
			want:     26,
		},
		{
			testName: "Strike",
			throws:   []int{10, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     14,
		},
		{
			testName: "Strike in Last Frame",
			throws:   []int{10, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 2, 3},
			want:     29,
		},
		{
			testName: "All Spares",
			throws:   []int{8, 2, 1, 9, 5, 5, 6, 4, 7, 3, 8, 2, 1, 9, 8, 2, 7, 3, 8, 2, 3},
			//                11    15    16    17    18    11    18    17    18    13 = 154
			want: 154,
		},
		{
			testName: "All Strikes",
			throws:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			want:     300,
		},
		{
			testName: "Strikes & Spares",
			throws:   []int{10, 1, 9, 5, 5, 6, 4, 7, 3, 10, 1, 9, 8, 2, 7, 3, 8, 2, 3},
			//              20    15    16    17    20  20    18    17    18    13 = 174
			want: 174,
		},
	}
	for _, tt := range tests {
		got, err := Score(tt.throws)
		if err != nil {
			t.Errorf("%v: Score(%v) = (%v, %v) WANT (%v, nil)", tt.testName, tt.throws, got, err, tt.want)
		}
		if err == nil && got != tt.want {
			t.Errorf("%v: Score(%v) = (%v, nil) WANT (%v, nil)", tt.testName, tt.throws, got, tt.want)
		}
	}
}

func TestBuildFrame(t *testing.T) {
	tests := []struct {
		testName   string
		throws     []int
		wantFrame  frame
		wantThrows []int
	}{
		{
			testName:   "Zero",
			throws:     []int{0, 0},
			wantFrame:  frame{0, 0, false, false},
			wantThrows: []int{},
		},
		{
			testName:   "Non-zero values",
			throws:     []int{3, 1},
			wantFrame:  frame{3, 1, false, false},
			wantThrows: []int{},
		},
		{
			testName:   "Strike",
			throws:     []int{10, 5},
			wantFrame:  frame{10, 0, true, false},
			wantThrows: []int{5},
		},
		{
			testName:   "three throws",
			throws:     []int{0, 1, 2},
			wantFrame:  frame{0, 1, false, false},
			wantThrows: []int{2},
		},
		{
			testName:   "Spare",
			throws:     []int{5, 5},
			wantFrame:  frame{5, 5, false, true},
			wantThrows: []int{},
		},
		{
			testName:   "BonusAfterSpare",
			throws:     []int{5},
			wantFrame:  frame{5, 0, false, false},
			wantThrows: []int{},
		},
	}
	for _, tt := range tests {
		if gotF, gotT := buildFrame(tt.throws); gotF != tt.wantFrame || !reflect.DeepEqual(gotT, tt.wantThrows) {
			t.Errorf("%v: buildFrame(%v) = (%v,%v) WANT (%v,%v)", tt.testName, tt.throws, gotF, gotT, tt.wantFrame, tt.wantThrows)
		}
	}
}
