package minesweeper

import (
	"reflect"
	"testing"
)

func TestMinesweeper(t *testing.T) {
	tests := []struct {
		mineField MineField
		want      MineField
	}{
		{
			mineField: MineField{Rows: 0, Columns: 0, Field: [][]string{}},
			want:      MineField{Rows: 0, Columns: 0, Field: [][]string{}},
		},
		{
			mineField: MineField{Rows: 1, Columns: 1, Field: [][]string{{"*"}}},
			want:      MineField{Rows: 1, Columns: 1, Field: [][]string{{"*"}}},
		},
		{
			mineField: MineField{Rows: 1, Columns: 2, Field: [][]string{{"*", "."}}},
			want:      MineField{Rows: 1, Columns: 2, Field: [][]string{{"*", "1"}}},
		},
		{
			mineField: MineField{Rows: 1, Columns: 3, Field: [][]string{{"*", ".", "*"}}},
			want:      MineField{Rows: 1, Columns: 3, Field: [][]string{{"*", "2", "*"}}},
		},
		{
			mineField: MineField{Rows: 2, Columns: 2, Field: [][]string{{"*", "*"}, {"*", "*"}}},
			want:      MineField{Rows: 2, Columns: 2, Field: [][]string{{"*", "*"}, {"*", "*"}}},
		},
		{
			mineField: MineField{Rows: -1, Columns: -5, Field: [][]string{}},
			want:      MineField{Rows: -1, Columns: -5, Field: [][]string{}},
		},
		{
			mineField: MineField{Rows: 3, Columns: 5, Field: [][]string{
				{"*", "*", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", "*", ".", ".", "."},
			},
			},
			want: MineField{Rows: 3, Columns: 5, Field: [][]string{
				{"*", "*", "1", "0", "0"},
				{"3", "3", "2", "0", "0"},
				{"1", "*", "1", "0", "0"},
			},
			},
		},
		{
			mineField: MineField{Rows: 3, Columns: 3, Field: [][]string{{"*", "."}, {".", "*"}}},
			want:      MineField{Rows: 3, Columns: 3, Field: [][]string{{"*", "."}, {".", "*"}}},
		},
		{
			mineField: MineField{Rows: 2, Columns: 3, Field: [][]string{{"*", ".", "."}, {".", "*"}}},
			want:      MineField{Rows: 2, Columns: 3, Field: [][]string{{"*", ".", "."}, {".", "*"}}},
		},
	}
	for _, tt := range tests {
		got := Evaluate(tt.mineField)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Evaluate(%v):%v, WANT: %v", tt.mineField, got, tt.want)
		}
	}
}
