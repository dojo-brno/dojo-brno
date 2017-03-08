package smartvector

import "testing"

func TestStoreValueAtPositionAndGetItBack(t *testing.T) {
	v := (*SmartVector).New(nil, 1)
	v.Set(0, 10)
	got := v.Get(0)
	if got != 10 {
		t.Errorf("Faiiiiiil!!!!")
	}
}

func TestStoreEqualAdjacentValues(t *testing.T) {
	v := (*SmartVector).New(nil, 2)
	v.Set(0, 10)
	v.Set(1, 10)
	if got := v.Get(0); got != 10 {
		t.Errorf("v.Get(%v) == %v, want %v", 0, got, 10)
	}
	if got := v.Get(1); got != 10 {
		t.Errorf("v.Get(%v) == %v, want %v", 1, got, 10)
	}
}

func TestStoreDifferentAdjacentValues(t *testing.T) {
	v := (*SmartVector).New(nil, 3)
	v.Set(0, 10)
	v.Set(1, 11)
	v.Set(2, 12)
	if got := v.Get(0); got != 10 {
		t.Errorf("v.Get(%v) == %v, want %v", 0, got, 10)
	}
	if got := v.Get(1); got != 11 {
		t.Errorf("v.Get(%v) == %v, want %v", 1, got, 11)
	}
	if got := v.Get(2); got != 12 {
		t.Errorf("v.Get(%v) == %v, want %v", 2, got, 12)
	}

}
