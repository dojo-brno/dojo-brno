package gameoflife

import (
	"reflect"
	"testing"
)

func TestSurvive(t *testing.T) {
	type testCase struct {
		nLivingNeighbours int
		isLife            bool
	}

	testCases := []testCase{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
	}

	for _, tt := range testCases {
		if got := Survive(tt.nLivingNeighbours); got != tt.isLife {
			t.Errorf("Survive(%v):%v, WANT:%v", tt.nLivingNeighbours, got, tt.isLife)
		}
	}
}

func TestBorn(t *testing.T) {
	type testCase struct {
		nLivingNeighbours int
		isLife            bool
	}

	testCases := []testCase{
		{0, false},
		{1, false},
		{2, false},
		{3, true},
		{4, false},
	}

	for _, tt := range testCases {
		if got := Born(tt.nLivingNeighbours); got != tt.isLife {
			t.Errorf("Born(%v):%v, WANT:%v", tt.nLivingNeighbours, got, tt.isLife)
		}
	}
}

func TestGameOfLife(t *testing.T) {
	type testCase struct {
		Board     Board
		nextBoard Board
	}

	testCases := []testCase{
		{
			Board{
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
			Board{
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
		},
		{
			Board{
				{false, false, false, false, false, false, false, false},
				{false, true, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
			Board{
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
		},
		{
			Board{
				{false, false, false, false, false, false, false, false},
				{true, true, true, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
			Board{
				{false, false, false, false, false, false, false, false},
				{false, true, false, false, false, false, false, false},
				{false, true, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
		},
	}

	for _, tt := range testCases {
		if got := GameOfLife(tt.Board); !reflect.DeepEqual(tt.nextBoard, got) {
			a := printBoard(tt.Board)
			b := printBoard(got)
			c := printBoard(tt.nextBoard)
			t.Errorf("GameOfLife\nInitial Board: \n%v \nReturned Board: \n%v \nExpected Board:\n %v", a, b, c)
		}
	}
}
