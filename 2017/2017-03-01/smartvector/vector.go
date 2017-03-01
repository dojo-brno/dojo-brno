package smartvector

type SmartVector struct {
	posBitMask int
	value      int
	next       *SmartVector
}

func New(n int) *SmartVector {
	return &SmartVector{}
}

func (s SmartVector) Get(i int) int {
	if s.posBitMask&i != 0 {
		return s.value
	}
	return -1
}

func (s SmartVector) Set(i int, value int) {
	s.posBitMask |= i
	s.value = value
}
