package pokerhands

import "testing"

func TestHighestCard(t *testing.T) {
	tests := []struct {
		left  []string
		right []string
		want  int
	}{
		{
			left:  []string{"2H", "3D", "5S", "9C", "KD"},
			right: []string{"2C", "3H", "4S", "8C", "AH"},
			want:  1,
		},
		{
			left:  []string{"2C", "3H", "4S", "8C", "AH"},
			right: []string{"2H", "3D", "5S", "9C", "KD"},
			want:  -1,
		},
		{
			left:  []string{"2C", "3H", "4S", "8C", "AH"},
			right: []string{"2H", "3D", "4D", "8D", "AD"},
			want:  0,
		},
		{
			left:  []string{"2C", "3H", "4S", "8C", "KH"},
			right: []string{"2H", "3D", "4D", "8D", "QD"},
			want:  -1,
		},
		{
			left:  []string{"2H", "3D", "5S", "9C", "KD"},
			right: []string{"2C", "3H", "4S", "8C", "KH"},
			want:  -1,
		},
	}
	for _, test := range tests {
		got := Compare(test.left, test.right)
		if got != test.want {
			t.Errorf("Compare(%v, %v) = %v, want %v", test.left, test.right, got, test.want)
		}
	}
}
