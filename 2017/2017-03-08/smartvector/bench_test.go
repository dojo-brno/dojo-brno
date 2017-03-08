package smartvector

import "testing"

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Slice.New(nil, b.N)
		for j := 0; j < b.N; j++ {
			s.Set(j, b.N-j)
		}
	}
}
