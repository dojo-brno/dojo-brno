package smartvector

type SmartVector struct {
	posBitMap uint64
	value     int
	next      *SmartVector
}

var _ Vector = (*SmartVector)(nil)

func (v *SmartVector) New(n int) Vector {
	return &SmartVector{}
}

func (v *SmartVector) Set(i, value int) {
	if v.posBitMap == 0 {
		v.value = value
		v.posBitMap = setNthBit(v.posBitMap, i)
		// v.posBitMap |= uint64(i + 1)
		return
	}

	curr := v
	for ; curr.next != nil; curr = curr.next {
	}

	curr.next = &SmartVector{value: value}
}

func (v *SmartVector) Get(i int) int {
	for curr := v; !(curr.posBitMap & uint64(i+1)); curr = curr.next {

	}

	if i == 1 {
		return v.next.value
	}
	if i == 2 {
		return v.next.next.value
	}
	return v.value
}

func setNthBit(toModify uint64, n int) uint64 {
	// toBitwiseOr :
	return toModify + 1
}
