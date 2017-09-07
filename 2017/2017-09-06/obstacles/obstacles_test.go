package obstacles

import (
	"reflect"
	"testing"
)

func TestObstacles(t *testing.T) {
	tests := []struct {
		waypoints []point
		want      []point
	}{
		{
			waypoints: []point{{6, 4}},
			want:      []point{{6, 4}},
		},
		{
			waypoints: []point{{0, 0}, {1, 1}},
			want:      []point{{0, 0}, {1, 1}},
		},
		{
			waypoints: []point{{0, 0}, {0, 3}},
			want:      []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			waypoints: []point{{6, 4}, {3, 3}},
			want:      []point{{6, 4}, {5, 3}, {4, 3}, {3, 3}},
		},
		{
			waypoints: []point{{6, 4}, {3, 3}, {-2, 10}},
			want:      []point{{6, 4}, {5, 3}, {4, 3}, {3, 3}, {2, 4}, {1, 5}, {0, 6}, {-1, 7}, {-2, 8}, {-2, 9}, {-2, 10}},
		},
	}
	for _, tt := range tests {
		if got := FindPath(tt.waypoints); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FindPath(%v) = %v WANT %v", tt.waypoints, got, tt.want)
		}
	}
}
