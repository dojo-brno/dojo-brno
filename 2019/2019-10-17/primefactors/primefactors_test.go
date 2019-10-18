package primefactors

import (
	"reflect"
	"testing"
)

func PrimeFactors(n int) []int {
	var factors []int
	if n == 1 {
		return []int{1}
	}
	for ; n != 0 && n%2 == 0 ; n = n / 2 {
		factors = append(factors, 2)
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func TestPrimeFactors(t *testing.T) {
	type testCase = struct {
		number       int
		primeFactors []int
	}
	testCases := []testCase{
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2, 2}},
		{5, []int{5}},
		{6, []int{2, 3}},
		{7, []int{7}},
		{8, []int{2, 2, 2}},
		{9, []int{3, 3}},
	}

	for _, tt := range testCases {
		if got := PrimeFactors(tt.number); !reflect.DeepEqual(got, tt.primeFactors) {
			t.Errorf("PrimeFactors(%v):%v,  WANTED:%v", tt.number, got, tt.primeFactors)
		}
	}
}
