package ant

import (
	"reflect"
	"testing"
)

func TestStep(t *testing.T) {
	tests := []struct {
		ant       Ant
		board     Board
		wantAnt   Ant
		wantBoard Board
	}{
		{
			ant:       Ant{X: 0, Y: 0, Orientation: Down},
			board:     Board{{Black, Black}},
			wantBoard: Board{{White, Black}},
			wantAnt:   Ant{X: 1, Y: 0, Orientation: Right},
		},
		{
			ant:       Ant{X: 0, Y: 0, Orientation: Up},
			board:     Board{{White, Black}},
			wantBoard: Board{{Black, Black}},
			wantAnt:   Ant{X: 1, Y: 0, Orientation: Right},
		},
		{
			ant:       Ant{X: 0, Y: 0, Orientation: Right},
			board:     Board{{Black}, {Black}},
			wantBoard: Board{{White}, {Black}},
			wantAnt:   Ant{X: 0, Y: 1, Orientation: Up},
		},
		{
			ant:       Ant{X: 0, Y: 0, Orientation: Right},
			board:     Board{{Black}, {Black}},
			wantBoard: Board{{White}, {Black}},
			wantAnt:   Ant{X: 0, Y: 1, Orientation: Up},
		},
	}

	for _, test := range tests {
		if gotBoard, gotAnt := Step(test.board, test.ant); !reflect.DeepEqual(gotBoard, test.wantBoard) || gotAnt != test.wantAnt {
			t.Errorf("Step(%v, %v) = (%v, %v) WANT (%v, %v)", test.board, test.ant, gotBoard, gotAnt, test.wantBoard, test.wantAnt)
		}
	}
}
