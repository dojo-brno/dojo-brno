package ant

import "testing"

type Color int

const (
	Black Color = iota
	White
)

type Board [5][5]Color

type Ant struct {
	x, y      int
	direction int
}

func TestStep(t *testing.T) {
	type testCase struct {
		board     Board
		ant       Ant
		nextBoard Board
		nextAnt   Ant
	}

	testCases := []testCase{
		testCase{
			Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{2, 2, 0},
			Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{3, 2, 3},
		},
		testCase{
			Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, White, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{3, 2, 3},
			Board{
				{White, White, White, White, White},
				{White, White, White, White, White},
				{White, White, Black, Black, White},
				{White, White, White, White, White},
				{White, White, White, White, White},
			},
			Ant{3, 3, 6},
		},
	}

	for _, tc := range testCases {
		if gotBoard, gotAnt := Step(tc.board, tc.ant); gotBoard != tc.nextBoard || gotAnt != tc.nextAnt {
			t.Errorf("Step(%v, %v): %v %v want %v %v", tc.board, tc.ant, gotBoard, gotAnt, tc.nextBoard, tc.nextAnt)
		}
	}
}

func TestAntStep(t *testing.T) {
	type testCase struct {
		ant       Ant
		color     Color
		nextAnt   Ant
		nextColor Color
	}

	testCases := []testCase{
		testCase{Ant{3, 2, 3}, White, Ant{3, 3, 6}, Black},
		testCase{Ant{2, 2, 0}, White, Ant{3, 2, 3}, Black},
	}

	for _, tc := range testCases {
		if gotAnt, gotColor := AntStep(tc.ant, tc.color); gotAnt != tc.ant || gotColor != tc.color {
			t.Errorf("AntStep(%v, %v): %v %v want %v %v", tc.ant, tc.color, gotAnt, gotColor, tc.nextAnt, tc.nextColor)
		}
	}
}
