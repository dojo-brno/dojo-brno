package diamond

import (
	"reflect"
	"testing"
)

func TestDiamond(t *testing.T) {
	tests := []struct {
		letter byte
		want   Diamond
	}{
		{
			letter: 'A',
			want:   "A\n",
		},
		{
			letter: 'B',
			want:   " A \nB B\n A \n",
		},
		{
			letter: 'C',
			want:   "  A  \n B B \nC  C\n B B \n  A  \n",
		},
		{
			letter: 'D',
			want:   "   A   \n  B B  \n C  C \nD   D\n C  C \n  B B  \n   A   \n",
		},
	}
	for _, tt := range tests {
		if got := GetDiamond(tt.letter); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Diamond(%q) returned %q but want %q", tt.letter, got, tt.want)
		}
	}
}

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex       int
		numberOfCycles int
		want           Diamond
	}{
		{
			rowIndex:       0,
			numberOfCycles: 1,
			want:           "A\n",
		},
		{
			rowIndex:       0,
			numberOfCycles: 2,
			want:           " A \n",
		},
		{
			rowIndex:       1,
			numberOfCycles: 2,
			want:           "B B\n",
		},
		{
			rowIndex:       2,
			numberOfCycles: 3,
			want:           "C  C\n",
		},
		{
			rowIndex:       1,
			numberOfCycles: 3,
			want:           " B B \n",
		},
		{
			rowIndex:       0,
			numberOfCycles: 5,
			want:           "    A    \n",
		},
		{
			rowIndex:       2,
			numberOfCycles: 5,
			want:           "  C  C  \n",
		},
	}
	for _, tt := range tests {
		if got := getRow(tt.rowIndex, tt.numberOfCycles); got != tt.want {
			t.Errorf("getRow(%v, %v) = %q want %q", tt.rowIndex, tt.numberOfCycles, got, tt.want)
		}
	}

}
