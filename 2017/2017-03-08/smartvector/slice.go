package smartvector

type Slice []int

var _ Vector = Slice(nil)

func (s Slice) New(n int) Vector {
	return make(Slice, n)
}

func (s Slice) Set(i, value int) {
	if i < 0 || i >= len(s) {
		panic("index out of bounds")
	}
	s[i] = value
}

func (s Slice) Get(i int) int {
	if i < 0 || i >= len(s) {
		panic("index out of bounds")
	}
	return s[i]
}
