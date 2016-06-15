package pokerhands

import "testing"

func TestCompare(t *testing.T) {
	tests := []struct {
		left, right []string
		want        int
	}{
		{
			left:  []string{"2H", "3D", "5S", "9C", "KD"},
			right: []string{"2D", "3H", "5D", "9S", "KC"},
			want:  0,
		},
		{
			left:  []string{"2H", "3D", "5S", "9C", "QD"},
			right: []string{"2D", "3H", "5D", "9S", "KC"},
			want:  1,
		},
		{
			left:  []string{"2H", "3D", "5S", "9C", "4D"},
			right: []string{"2D", "3H", "5D", "9S", "KC"},
			want:  1,
		},
		{
			left:  []string{"2H", "3D", "5S", "9C", "4C"},
			right: []string{"2D", "3H", "5D", "9S", "KC"},
			want:  1,
		},
		{
			left:  []string{"2D", "3H", "5D", "9S", "QC"},
			right: []string{"2H", "3D", "5S", "9C", "4D"},
			want:  -1,
		},
		{
			left:  []string{"2D", "3H", "5D", "9S", "QC"},
			right: []string{"2H", "3D", "5S", "9C", "6D"},
			want:  -1,
		},
		{
			left:  []string{"2D", "3H", "5D", "9S", "6D"},
			right: []string{"2H", "3D", "5S", "9C", "QC"},
			want:  1,
		},
		{
			left:  []string{"2D", "3H", "5D", "9S", "3S"},
			right: []string{"2H", "3D", "5S", "9C", "4C"},
			want:  1,
		},
		{
			left:  []string{"2D", "3H", "5D", "3S", "9S"},
			right: []string{"2H", "3D", "5S", "4C", "9C"},
			want:  1,
		},
		{
			left:  []string{"2D", "3H", "3S", "5D", "9S"},
			right: []string{"2H", "3D", "4C", "5S", "9C"},
			want:  1,
		},
	}
	for _, tt := range tests {
		result := Compare(tt.left, tt.right)
		if result != tt.want {
			t.Errorf("Compare(%v, %v) = %v, want %v", tt.left, tt.right, result, tt.want)
		}
	}
}
