package dojorange

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		i    Interval
		num  int
		want bool
	}{
		{i: Interval{}, num: 0, want: false},
		{i: Interval{FirstIn: true, LastIn: true}, num: 0, want: true},
		{i: Interval{First: 0, Last: 1, FirstIn: true, LastIn: true}, num: 2, want: false},
		{i: Interval{First: 0, Last: 1, FirstIn: true, LastIn: false}, num: 1, want: false},
		{i: Interval{First: 0, Last: 1, FirstIn: true, LastIn: true}, num: -1, want: false},
		{i: Interval{First: 0, Last: 1, FirstIn: false, LastIn: true}, num: 0, want: false},
		{i: Interval{First: 1, Last: 3, FirstIn: false, LastIn: false}, num: 2, want: true},
		{i: Interval{First: 1, Last: 3, FirstIn: false, LastIn: true}, num: 3, want: true},
		{i: Interval{First: 1, Last: 3, FirstIn: true, LastIn: true}, num: 1, want: true},
	}

	for i, tt := range tests {
		got := tt.i.Contains(tt.num)
		if got != tt.want {
			t.Errorf("%v: %v.Contains(%v) = %v want %v", i, tt.i, tt.num, got, tt.want)
		}
	}
}

func TestAllPoints(t *testing.T) {
	tests := []struct {
		i    Interval
		want []int
	}{
		{i: Interval{}, want: nil},
		{i: Interval{FirstIn: true, LastIn: true}, want: []int{0}},
	}

	for i, tt := range tests {
		got := tt.i.GetAllPoints()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v: %v.GetAllPoints(): %v, WANT %v", i, tt.i, got, tt.want)
		}
	}

}
