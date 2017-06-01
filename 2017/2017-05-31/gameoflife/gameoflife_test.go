package gameoflife

import (
	"reflect"
	"testing"
)

func TestNextGen(t *testing.T) {
	tests := []struct {
		input *Cell
		want  *Cell
	}{
		{
			&Cell{Dead, nil},
			&Cell{Dead, nil},
		},
		{
			&Cell{Live, nil},
			&Cell{Dead, nil},
		},
	}
	for _, tt := range tests {
		if got := NextGen(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NextGen(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestTwoCells(t *testing.T) {
	input := &Cell{Dead, []*Cell{&Cell{Dead, nil}}}
	input.neighbors[0].neighbors = []*Cell{input}
	want := &Cell{Dead, []*Cell{&Cell{Dead, nil}}}
	want.neighbors[0].neighbors = []*Cell{want}

	if got := NextGen(input); !reflect.DeepEqual(got, want) {
		t.Errorf("NextGen(%v) = %v, want %v \ngot.neighbors[0]=%v \nwant.neighbors[0] %v", input, got, want, got.neighbors[0], want.neighbors[0])
	}
}

func TestThreeCells(t *testing.T) {
	top, center, bottom := new(Cell), new(Cell), new(Cell)

	top.neighbors = []*Cell{center}
	center.neighbors = []*Cell{top, bottom}
	bottom.neighbors = []*Cell{center}
	want := (*Cell)(nil)

	if got := NextGen(top); !reflect.DeepEqual(got, want) {
		t.Errorf("NextGen(%v) = %v, want %v\ngot json: %s (%v)\nwant json: %s (%v)", top, got, want, gotjson, goterr, wantjson, wanterr)
	}
}
