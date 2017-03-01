package smartvector

import "testing"

func TestSingleValue(t *testing.T) {
	s := New(1)
	i := 0
	s.Set(i, 7)
	want := 7
	if got := s.Get(0); got != want {
		t.Errorf("s.Get(%v) = %v, want %v", i, got, want)
	}
}

func TestTwoValues(t *testing.T) {
	s := New(10)
	i := 0
	s.Set(i, 7)
	j := 3
	s.Set(j, 10)
	want_i := 7
	want_j := 10
	if got := s.Get(i); got != want_i {
		t.Errorf("s.Get(%v) = %v, want %v", i, got, want_i)
	}
	if got := s.Get(j); got != want_j {
		t.Errorf("s.Get(%v) = %v, want %v", j, got, want_j)
	}

}
