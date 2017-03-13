package smartvector

import "testing"

func TestStoreValueAtPositionAndGetItBack(t *testing.T) {
	v := (*SmartVector).New(nil, 1)
	v.Set(0, 10)
        assertEquals(t, 0, 10, v.Get(0))
}

func TestStoreEqualAdjacentValues(t *testing.T) {
	v := (*SmartVector).New(nil, 2)
	v.Set(0, 10)
	v.Set(1, 10)

        assertEquals(t, 0, 10, v.Get(0))
        assertEquals(t, 1, 10, v.Get(1))

}

func TestStoreDifferentAdjacentValues(t *testing.T) {
	v := (*SmartVector).New(nil, 3)
	v.Set(0, 10)
	v.Set(1, 11)
	v.Set(2, 12)
       
        assertEquals(t, 0, 10, v.Get(0))
        assertEquals(t, 1, 11, v.Get(1))
        assertEquals(t, 2, 12, v.Get(2))
}

func TestStoreBothDifferentAndEqualAdjacentValues(t *testing.T) {
	v := (*SmartVector).New(nil, 4)
	v.Set(0, 10)
	v.Set(1, 11)
	v.Set(2, 12)
        v.Set(5, 10)
       
        assertEquals(t, 0, 10, v.Get(0))
        assertEquals(t, 1, 11, v.Get(1))
        assertEquals(t, 2, 12, v.Get(2))
        assertEquals(t, 5, 10, v.Get(5))
}

func assertEquals(t *testing.T, pos, expected, actual int) {
    if expected != actual {
        t.Errorf("Getting element at position [%v] from vector, expected value [%v], but got [%v]", pos, expected, actual)
    }
}
