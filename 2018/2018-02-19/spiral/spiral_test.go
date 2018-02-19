package spiral

import "testing"
import "reflect"

type testMatrixDecomposition struct {
	matrix   [][]int
	position int
	want     []int
}

func TestUpLine(t *testing.T) {
	tests := []testMatrixDecomposition{
		{[][]int{}, 1, []int{}},
		{[][]int{{0}}, 1, []int{0}},
		{[][]int{{5}}, 1, []int{5}},
		{[][]int{{5, 6}}, 1, []int{5}},
		{[][]int{{5, 6, 7}}, 1, []int{5, 6}},
		//{[][]int{{5 ,6 }, {7 ,6 }}, 1, []int{5}},
		{nil, -1, nil},
	}

	for _, tt := range tests {
		if got := UpLine(tt.matrix, tt.position); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("UpLine(%v, %v)-->> %v, WANT: %v", tt.matrix, tt.position, got, tt.want)
			t.Errorf("UpLine(%p, %p)-->> %p, WANT: %p", tt.matrix, tt.position, got, tt.want)
		}
	}
}
