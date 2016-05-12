package nonogram

import (
	"fmt"
	"strings"
	"testing"
)

func Solve(x, y [][]int) string {
	var lines []string
	for _, line := range x {
		lines = append(lines, fmt.Sprintf("%-*s", len(y), strings.Repeat("#", line[0])))
	}
	return strings.Join(lines, "\n")
}

func TestNonogram(t *testing.T) {
	tests := []struct {
		x, y [][]int
		want string
	}{
		{
			x:    [][]int{{1}},
			y:    [][]int{{1}},
			want: "#",
		},
		{
			x:    [][]int{{2}},
			y:    [][]int{{1}, {1}},
			want: "##",
		},
		{
			x:    [][]int{{3}},
			y:    [][]int{{1}, {1}, {1}},
			want: "###",
		},
		{
			x:    [][]int{{2}, {2}},
			y:    [][]int{{2}, {2}},
			want:
			"##\n" +
			"##",
		},
		{
			x:    [][]int{{2}, {1}},
			y:    [][]int{{2}, {1}},
			want:
			"##\n" +
			"# ",
		},
	}
	for _, tt := range tests {
		if got := Solve(tt.x, tt.y); got != tt.want {
			t.Errorf("Solve(%v, %v) = %#v, want %#v", tt.x, tt.y, got, tt.want)
		}
	}
}
