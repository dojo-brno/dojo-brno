package square16

import "testing"

func TestSquare16(t *testing.T) {
	tests := []struct {
		in   Solution
		want bool
	}{
		{
			in:   Solution{},
			want: false,
		},
		{
			in: Solution{
				[]Point{
					{0, 0}, {0, 4},
					{1, 4}, {1, 1},
					{2, 1}, {2, 2},
					{3, 2}, {3, 0},
				},
				[]Point{
					{1, 1}, {1, 4},
					{4, 4}, {4, 0},
					{3, 0}, {3, 2},
					{2, 2}, {2, 1},
				},
			},
			want: true,
		},
		{
			in: Solution{
				[]Point{
					{0, 0}, {0, 4},
					{2, 4}, {2, 0},
				},
				[]Point{
					{2, 0}, {2, 4},
					{4, 4}, {4, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsSolution(tt.in)
		if got != tt.want {
			t.Errorf("IsSolution(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	p := []Point{
		{0, 0}, {0, 4},
		{1, 4}, {1, 1},
		{2, 1}, {2, 2},
		{3, 2}, {3, 0},
	}
	got := perimeter(p)
	want := 16
	if got != want {
		t.Errorf("perimeter(%v) = %v, want %v", p, got, want)
	}
}
