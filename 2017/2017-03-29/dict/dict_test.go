package dict

import "testing"

func TestDict(t *testing.T) {
	tests := []struct {
		name    string
		newDict func() Dict
	}{
		{"NewOcdoganRbt", NewOcdoganRbt},
		{"NewMap", NewMap},
	}
	for _, tt := range tests {
		d := tt.newDict()
		d.Set(5, "foo")
		if got := d.Get(5); got != "foo" {
			t.Errorf("%s: d.Get(%d) = %q, want %q", tt.name, 5, got, "foo")
		}
		d.Delete(5)
		if got := d.Get(5); got != "" {
			t.Errorf("%s: d.Get(%d) = %q, want %q", tt.name, 5, got, "")
		}
	}
}

func BenchmarkOcdoganRbt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewOcdoganRbt()
		for i := 0; i < b.N; i++ {
			d.Set(i, "my benchmark value")
		}
	}
}

func BenchmarkGoMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := NewMap()
		for i := 0; i < b.N; i++ {
			d.Set(i, "my benchmark value")
		}
	}
}
