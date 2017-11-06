package mastermind

import "testing"

func TestMasterMind(t *testing.T) {
	tests := []struct {
		secret         Pegs
		guess          Pegs
		wantWellPlaced int
		wantMisplaced  int
	}{
		{
			secret:         Pegs{Black, White, Yellow, Red, Blue},
			guess:          Pegs{Gray, Gray, Gray, Gray, Gray},
			wantWellPlaced: 0,
			wantMisplaced:  0,
		},
		{
			secret:         Pegs{Black, Gray, Yellow, Red, Blue},
			guess:          Pegs{Gray, Gray, Gray, Gray, Gray},
			wantWellPlaced: 1,
			wantMisplaced:  0,
		},
		{
			secret:         Pegs{Black, Gray, Yellow, Red, Blue},
			guess:          Pegs{Black, Gray, Gray, Gray, Gray},
			wantWellPlaced: 2,
			wantMisplaced:  0,
		},
		{
			secret:         Pegs{Black, Gray, Yellow, Red, Blue},
			guess:          Pegs{Black, Gray, Gray, Yellow, Gray},
			wantWellPlaced: 2,
			wantMisplaced:  1,
		},
		{
			secret:         Pegs{Black, Gray, Yellow, Red, Blue},
			guess:          Pegs{Black, Gray, Red, Yellow, Gray},
			wantWellPlaced: 2,
			wantMisplaced:  2,
		},
		{
			secret:         Pegs{Gray, Red, Blue, Black, Yellow},
			guess:          Pegs{Black, Gray, Yellow, Red, Blue},
			wantWellPlaced: 0,
			wantMisplaced:  5,
		},
	}

	for _, tt := range tests {
		if gotWellPlaced, gotMisplaced := MasterMind(tt.secret, tt.guess); tt.wantWellPlaced != gotWellPlaced || tt.wantMisplaced != gotMisplaced {
			t.Errorf("MasterMind(%v, %v) = (%v, %v) WANT (%v, %v)", tt.secret, tt.guess, gotWellPlaced, gotMisplaced, tt.wantWellPlaced, tt.wantMisplaced)
		}
	}
}
