package codecracker

import "testing"

func GetMagicNumbers() int {
	return 13
}

func TestGetMagicNumbers(t *testing.T) {
	type testCase = struct {
		number  int
		isMagic bool
	}

	testCases := []testCase{
		{13, true},
	}

	for _, tt := range testCases {
		if got := GetMagicNumbers(); got == tt.number {
			if tt.isMagic {
				t.Errorf("GetMagicNumbers():%v, should return only magic numbers..", got)
			}
		}
	}
}
