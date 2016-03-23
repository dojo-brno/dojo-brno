package minesweeper

import "testing"

type Minefield string

func New(s string) Minefield {
	return Minefield(s)
}

func (m Minefield) Annotate() Minefield {
	var s []rune
	for j, c := range m {
		switch c {
		case '*':
			s = append(s, '*')
		case '.':
			count := m.CountNeighborMines(0, j)
			s = append(s, '0'+rune(count))
		}
	}
	return New(string(s))
}

func (m Minefield) CountNeighborMines(i, j int) int {
	if m.HasBomb(i, j-1) || m.HasBomb(i, j+1) {
		return 1
	}
	return 0
}

func (m Minefield) HasBomb(i, j int) bool {
	return j >= 0 && j < len(m) && m[j] == '*'
}
func TestMinesweeper(t *testing.T) {
	testCases := []struct {
		minefield Minefield
		annotated Minefield
	}{
		{New("*"), New("*")},
		{New("."), New("0")},
		{New(".."), New("00")},
		{New("**"), New("**")},
		{New("*."), New("*1")},
		{New(".*"), New("1*")},
		{New("***"), New("***")},
		{New("..."), New("000")},
		{New("*.."), New("*10")},
		{New(".*."), New("1*1")},
		{New("..*"), New("01*")},
		{New("****"), New("****")},
		{New("...."), New("0000")},
		{New("*..."), New("*100")},
		{New(".*.."), New("1*10")},
	}
	for _, tc := range testCases {
		if got := tc.minefield.Annotate(); got != tc.annotated {
			t.Errorf("%#v.Annotate() = %#v, want %#v", tc.minefield, got, tc.annotated)
		}
	}
}
