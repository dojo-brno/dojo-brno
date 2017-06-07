package gameoflife

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		initial Game
		want    Game
	}{
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{4, 1}: E,
					{3, 2}: E,
					{4, 2}: E,
				},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{3, 1}: E,
					{4, 1}: E,
					{3, 2}: E,
					{4, 2}: E,
				},
			},
		},
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World:   Grid{},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World:   Grid{},
			},
		},
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{0, 0}: E,
				},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World:   Grid{},
			},
		},
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{0, 0}: E,
					{0, 1}: E,
				},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World:   Grid{},
			},
		},
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{0, 0}: E,
					{0, 1}: E,
					{0, 2}: E,
				},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{0, 1}: E,
				},
			},
		},
		{
			initial: Game{
				Rows:    4,
				Columns: 8,
				World: Grid{
					{0, 0}: E,
					{0, 7}: E,
					{3, 0}: E,
					{3, 7}: E,
				},
			},
			want: Game{
				Rows:    4,
				Columns: 8,
				World:   Grid{},
			},
		},
	}
	for _, tt := range tests {
		if got := tt.initial.Next(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v.Next() = %v, want %v", tt.initial, got, tt.want)
		}
	}
}
