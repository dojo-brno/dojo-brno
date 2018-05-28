package game

import (
	"reflect"
	"testing"
)

func TestStep(t *testing.T) {
	tests := []struct {
		board []Cell
		want  []Cell
	}{
		{
			board: []Cell{},
			want:  []Cell{},
		},
		{
			board: []Cell{
				Cell{1, 1},
			},
			want: []Cell{},
		},
		{
			board: []Cell{
				Cell{0, 0},
				Cell{1, 1},
			},
			want: []Cell{},
		},
		{
			board: []Cell{
				Cell{0, 0}, Cell{0, 1},
				Cell{1, 1},
			},
			want: []Cell{
				Cell{0, 0}, Cell{0, 1},
				Cell{1, 1},
			},
		},
		{
			board: []Cell{
				Cell{0, 1},
				Cell{1, 0}, Cell{1, 1},
			},
			want: []Cell{
				Cell{0, 1},
				Cell{1, 0}, Cell{1, 1},
			},
		},
		{
			board: []Cell{
				Cell{0, 0},
				Cell{0, 1},
				Cell{0, 2},
			},
			want: []Cell{
				Cell{0, 1},
			},
		},
	}

	for _, tt := range tests {
		if got := Step(tt.board); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Step(%v) = %v, WANT = %v", tt.board, got, tt.want)
		}
	}
}

func TestGetNumberOfLivingNeighbours(t *testing.T) {
	tests := []struct {
		board []Cell
		cell  Cell
		want  int // numberOfLivingNeighbours
	}{
		{
			board: []Cell{
				Cell{1, 1},
			},
			cell: Cell{1, 1},
			want: 0,
		},
		{
			board: []Cell{
				Cell{0, 0}, Cell{1, 0},
				Cell{1, 1},
			},
			cell: Cell{1, 1},
			want: 2,
		},
		{
			board: []Cell{
				Cell{0, 0}, Cell{2, 2},
				Cell{0, 2},
			},
			cell: Cell{0, 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := GetNumberOfLivingNeighbours(tt.board, tt.cell); got != tt.want {
			t.Errorf("GetNumberOfLivingNeighbours(%v, %v) = %v, WANT = %v", tt.board, tt.cell, got, tt.want)
		}
	}

}
