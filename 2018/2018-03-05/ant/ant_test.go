package ant

import (
	"reflect"
	"testing"
)

func TestStep(t *testing.T) {
	tests := []struct {
		world World
		want  World
	}{
		{
			world: World{
				Ant:   Ant{Position: [2]int{0, 0}, Direction: 0},
				Board: [][]Color{{White}, {White}},
			},
			want: World{
				Ant:   Ant{Position: [2]int{1, 0}, Direction: 3},
				Board: [][]Color{{Black}, {White}},
			},
		},
		{
			world: World{
				Ant:   Ant{Position: [2]int{0, 1}, Direction: 3},
				Board: [][]Color{{White, White}},
			},
			want: World{
				Ant:   Ant{Position: [2]int{0, 0}, Direction: 6},
				Board: [][]Color{{White, Black}},
			},
		},
	}
	for _, tt := range tests {
		origin := tt.world
		if err := tt.world.Step(); err != nil {
			t.Errorf("%v.Step(): %v, WANT: nil", origin, err)
			continue
		}
		if !reflect.DeepEqual(tt.want, tt.world) {
			t.Errorf("%v.Step(): %v, WANT: %v", origin, tt.world, tt.want)
		}
	}
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		ant   Ant
		color Color
		want  Ant
	}{
		{
			ant:   Ant{Position: [2]int{0, 0}, Direction: 0},
			color: White,
			want:  Ant{Position: [2]int{1, 0}, Direction: 3},
		},
		{
			ant:   Ant{Position: [2]int{0, 0}, Direction: 9},
			color: White,
			want:  Ant{Position: [2]int{0, 1}, Direction: 0},
		},
		{
			ant:   Ant{Position: [2]int{0, 0}, Direction: 0},
			color: Black,
			want:  Ant{Position: [2]int{-1, 0}, Direction: 9},
		},
	}
	for _, tt := range tests {
		if got := tt.ant.GetNext(tt.color); got != tt.want {
			t.Errorf("(%v).GetNext(%v): %v, WANT: %v", tt.ant, tt.color, got, tt.want)
		}
	}
}
