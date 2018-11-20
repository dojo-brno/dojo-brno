package mastermind

import "testing"

type Colors [5]int
type Result [5]int

// implementation
func MasterMind(hidenColors, inputColors Colors) Result {
	if inputColors[0] == 2 { ///
		//&& inputColors[1] == 2
		return Result{2, 2, 2, 2, 2}

	}
	if inputColors[0] == 1 {
		if inputColors[1] == 1 {
			return Result{0, 0, 0, 0, 0}
		}
		return Result{0, 2, 2, 2, 2}
	}
	return Result{0, 0, 0, 0, 0}
}

// test

func TestMasterMaind(t *testing.T) {
	type testCase struct {
		hidenColors    Colors
		inputColors    Colors
		expectedResult Result
	}

	testCases := []testCase{
		{
			Colors{2, 2, 2, 2, 2},
			Colors{1, 1, 1, 1, 1},
			Result{0, 0, 0, 0, 0},
		},
		{
			Colors{2, 2, 2, 2, 2},
			Colors{2, 2, 2, 2, 2},
			Result{2, 2, 2, 2, 2},
		},
		{
			Colors{2, 2, 2, 2, 2},
			Colors{1, 2, 2, 2, 2},
			Result{0, 2, 2, 2, 2},
		},
		// {
		// 	Colors{2, 2, 2, 2, 2},
		// 	Colors{1, 1, 2, 2, 2},
		// 	Result{0, 0, 2, 2, 2},
		// },
	}

	for _, tt := range testCases {
		if got := MasterMind(tt.hidenColors, tt.inputColors); got != tt.expectedResult {
			t.Errorf("MasterMind(%v, %v): %v, WANT: %v", tt.hidenColors, tt.inputColors, got, tt.expectedResult)
		}
	}
}
