package primefactors

import (
	"reflect"
	"testing"
)

func PrimeFactors(n int) []int {
	result := []int{}
	for diviser := 2; diviser <= n; diviser++ {
		for ; n%diviser == 0; n /= diviser {
			result = append(result, diviser)
		}
	}
	return result
}

func TestPrimeFactors(t *testing.T) {
	type testCase struct {
		number      int
		primefactos []int
	}
	testCases := []testCase{
		{1, []int{}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{6, []int{2, 3}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
		{2 * 2 * 2 * 11 * 11 * 19, []int{2, 2, 2, 11, 11, 19}},
	}
	for _, tt := range testCases {
		if got := PrimeFactors(tt.number); !reflect.DeepEqual(got, tt.primefactos) {
			t.Errorf("PrimeFactors(%v):%v, WANT:%v", tt.number, got, tt.primefactos)
		}
	}
}
