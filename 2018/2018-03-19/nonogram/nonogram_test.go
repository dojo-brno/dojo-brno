package nonogram

import "testing"

func TestValidateLine(t *testing.T) {
	tests := []struct {
		peace line
		want  bool
	}{
		{peace: line{NumberOfOccurrence: []int{1}, InputLine: []bool{true}}, want: true},
		{peace: line{NumberOfOccurrence: []int{1}, InputLine: []bool{false}}, want: false},
		{peace: line{NumberOfOccurrence: []int{1}, InputLine: []bool{false, true}}, want: true},
		{peace: line{NumberOfOccurrence: []int{1}, InputLine: []bool{true, true}}, want: false},
	}

	for _, tt := range tests {
		if got := ValidateLine(tt.peace); tt.want != got {
			t.Errorf("ValidateLine(%v) = %v, WANT = %v", tt.peace, got, tt.want)
		}
	}
}
