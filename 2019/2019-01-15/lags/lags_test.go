package lags

import "testing"

type requests []struct {
	start    int
	duration int
	price    int
}

func possible(start, duration, nextStart int) bool {
	if start+duration <= nextStart {
		return true
	}
	return false
}

func IsValidCombination(requests requests) bool {

	for i, r := range requests {
		if i+1 >= len(requests) {
			break
		}
		n := requests[i+1]
		if !possible(r.start, r.duration, n.start) {
			return false
		}
	}
	return true
}

func TestIsValidCombination(t *testing.T) {
	type testCase struct {
		requests requests
		want     bool
	}

	testCases := []testCase{
		{
			requests: requests{{0, 5, 100}},
			want:     true,
		},
		{
			requests: requests{{0, 5, 100}, {4, 3, 20}},
			want:     false,
		},
		{
			requests: requests{{0, 5, 100}, {5, 3, 100}},
			want:     true,
		},
		{
			requests: requests{{0, 6, 100}, {5, 3, 100}},
			want:     false,
		},
		{
			requests: requests{{0, 5, 100}, {5, 3, 100}, {1, 4, 1}},
			want:     false,
		},
		{
			requests: requests{{0, 5, 100}, {5, 3, 100}, {8, 4, 100}},
			want:     true,
		},
		{
			requests: requests{{0, 5, 100}, {5, 3, 100}, {9, 4, 100}},
			want:     true,
		},
		{
			requests: requests{{0, 5, 100}, {5, 3, 100}, {9, 4, 100}, {1, 4, 100}},
			want:     false,
		},
		{
			requests: requests{{0, 0, 1000}, {0, 0, 1000}},
			want:     false,
		},
	}

	for _, tt := range testCases {
		if got := IsValidCombination(tt.requests); got != tt.want {
			t.Errorf("IsValidCombination(%v): %v, WANT:%v", tt.requests, got, tt.want)
		}
	}
}
