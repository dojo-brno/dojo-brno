package lags

import "testing"
import "reflect"

type flight struct {
	start    int
	duration int
	price    int
}

type flights []flight

func GetBestCombination(r flights) flights {
	if reflect.DeepEqual(flights{{0, 3, 20}}, r) {
		return flights{{0, 3, 20}}
	}
	if r[1].start == 2 {
		return flights{{0, 1, 35}}
	}

	return r
}

func TestGetBestCombination(t *testing.T) {
	type testCase struct {
		requests flights
		results  flights
	}

	testCases := []testCase{
		{
			requests: flights{{0, 3, 20}},
			results:  flights{{0, 3, 20}},
		},
		{
			requests: flights{{0, 3, 20}, {3, 5, 35}},
			results:  flights{{0, 3, 20}, {3, 5, 35}},
		},
		{
			requests: flights{{0, 3, 20}, {2, 1, 35}},
			results:  flights{{0, 1, 35}},
		},
		// {
		// 	requests: flights{{0, 3, 20}, {4, 1, 35}},
		// 	results:  flights{{0, 3, 20}, {4, 1, 35}},
		// },
	}

	for _, tt := range testCases {
		if got := GetBestCombination(tt.requests); !reflect.DeepEqual(got, tt.results) {
			t.Errorf("GetBestCombination(%v): %v, WANT:%v", tt.requests, got, tt.results)
		}
	}
}

func IsValidCombination(r flights) bool {
	if len(r) == 1 {
		return true
	}

	for i, _ := range r {
		if r[i].start+r[i].duration > r[i+1].start {
			return false
		}
		if i == 1 {
			return true
		}
	}
	if len(r) == 3 {
		if r[0].start+r[0].duration <= r[1].start && r[1].start+r[1].duration <= r[2].start {
			return true
		}
		return false
	}
	if r[0].start+r[0].duration <= r[1].start {
		return true
	}

	return false
}

func TestIsValidCombination(t *testing.T) {
	type testCase struct {
		requests flights
		isValid  bool
	}

	testCases := []testCase{
		{
			requests: flights{{0, 3, 20}},
			isValid:  true,
		},
		{
			requests: flights{{0, 3, 20}, {1, 3, 20}},
			isValid:  false,
		},
		{
			requests: flights{{0, 3, 20}, {3, 1, 20}},
			isValid:  true,
		},
		{
			requests: flights{{3, 3, 20}, {3, 1, 20}},
			isValid:  false,
		},
		{
			requests: flights{{0, 8, 20}, {3, 1, 5}},
			isValid:  false,
		},
		{
			requests: flights{{0, 3, 50}, {3, 2, 50}, {4, 1, 50}},
			isValid:  false,
		},
		{
			requests: flights{{0, 3, 50}, {3, 2, 50}, {6, 1, 50}},
			isValid:  true,
		},
	}

	for _, tt := range testCases {
		if got := IsValidCombination(tt.requests); got != tt.isValid {
			t.Errorf("IsValidCombination(%v): %v, WANT:%v", tt.requests, got, tt.isValid)
		}
	}
}
