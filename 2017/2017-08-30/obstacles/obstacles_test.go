package obstacles

import (
	"reflect"
	"testing"
)

func TestFindPath(t *testing.T) {
	tests := []struct {
		rows, cols int
		waypoints  []Point
		obstacles  []Obstacle
		want       []Point
	}{
		{
			rows: 2, cols: 2,
			waypoints: []Point{{0, 0}, {1, 1}},
			want:      []Point{{0, 0}, {1, 1}},
		},
		{
			rows: 3, cols: 3,
			waypoints: []Point{{0, 0}, {2, 2}},
			want:      []Point{{0, 0}, {1, 1}, {2, 2}},
		},
		{
			rows: 4, cols: 4,
			waypoints: []Point{{0, 0}, {3, 3}},
			want:      []Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
		},
		{
			rows: 10, cols: 10,
			waypoints: []Point{{0, 0}, {3, 3}},
			want:      []Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
		},
		{
			rows: 4, cols: 4,
			waypoints: []Point{{0, 3}, {3, 0}},
			want:      []Point{{0, 3}, {1, 2}, {2, 1}, {3, 0}},
		},
	}
	for _, tt := range tests {
		got := FindPath(tt.rows, tt.cols, tt.waypoints, tt.obstacles)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindPath(%v, %v, %v, %v) = %v, want %v", tt.rows, tt.cols, tt.waypoints, tt.obstacles, got, tt.want)
		}
	}
}
