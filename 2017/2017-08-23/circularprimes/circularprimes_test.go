package circularprimes

import (
	"reflect"
	"testing"
)

func TestIsCircularPrime(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{1, false},
		{2, true},
		{23, false},
	}
	for _, tt := range tests {
		if got := IsCircularPrime(tt.number); got != tt.want {
			t.Errorf("IsCircularPrime(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		number int
		want   bool
	}{
		{1, false},
		{2, true},
		{8924, false},
		{999931, true},
		{999932, false},
	}
	for _, tt := range tests {
		if got := isPrime(tt.number); got != tt.want {
			t.Errorf("isPrime(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}

func TestRotations(t *testing.T) {
	tests := []struct {
		number int
		want   []int
	}{
		{2, []int{2}},
		{23, []int{23, 32}},
		{123, []int{123, 231, 312}},
		{1234, []int{1234, 2341, 3412, 4123}},
		{12345, []int{12345, 23451, 34512, 45123, 51234}},
	}
	for _, tt := range tests {
		if got := rotations(tt.number); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("rotations(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}

func TestHowManyBelow(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{0, 0},
		{10, 4},
		{100, 13},
		{1000, 25},
		{10000, 33},
		{100000, 43},
		{1000000, 55},
	}
	for _, tt := range tests {
		if got := HowManyBelow(tt.number); got != tt.want {
			t.Errorf("HowManyBelow(%d) = %v, want %v", tt.number, got, tt.want)
		}
	}
}
