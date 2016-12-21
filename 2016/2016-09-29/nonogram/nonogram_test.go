package nonogram

import "testing"

func TestNonogram(t *testing.T) {
	/*
	      0
	   0 [ ]
	*/
	problem := Problem{
		Vertical:   [][]int{{0}},
		Horizontal: [][]int{{0}},
	}
	test := func(s Solution, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			if ok := isSolution(problem, s); ok != want {
				t.Errorf("isSolution(%#v, %#v) = %v, want %v", problem, s, ok, want)
			}
		}
	}
	t.Run("right solution", test(Solution{{false}}, true))
	t.Run("wrong solution", test(Solution{{true}}, false))

	/*
	      1
	   1  #
	*/
	problem = Problem{
		Vertical:   [][]int{{1}},
		Horizontal: [][]int{{1}},
	}
	t.Run("right solution", test(Solution{{true}}, true))
	t.Run("wrong solution", test(Solution{{false}}, false))

	/*
	      1 1
	   2  # #
	*/
	problem = Problem{
		Vertical:   [][]int{{1}, {1}},
		Horizontal: [][]int{{2}},
	}
	t.Run("right solution", test(Solution{{true, true}}, true))
	t.Run("wrong solution (wrong size)", test(Solution{{false}}, false))
	t.Run("wrong solution", test(Solution{{false}, {false}}, false))
	t.Run("wrong solution", test(Solution{{false}, {true}}, false))
	t.Run("wrong solution", test(Solution{{true}, {false}}, false))
}
