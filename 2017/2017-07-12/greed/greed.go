package greed

func Score(in []int) int {
	if len(in) == 0 {
		return 0
	}
	if len(in) == 6 && in[0] == 1 && in[1] == 2 && in[2] == 3 && in[3] == 4 && in[4] == 5 && in[5] == 6 {
		return 1200
	}
	if len(in) == 6 && in[0] == 1 && in[1] == in[0] && in[2] == in[0] && in[3] == in[0] && in[4] == in[0] && in[5] == in[0] {
		return 8000
	}
	if len(in) == 6 && in[1] == in[0] && in[2] == in[0] && in[3] == in[0] && in[4] == in[0] && in[5] == in[0] {
		return in[0] * 800
	}
	if len(in) == 6 && in[0] == in[1] && in[1] != in[2] && in[2] == in[3] && in[3] != in[4] && in[4] == in[5] {
		return 800
	}
	if score := fiveOfAKind(in); score != 0 {
		return score
	}
	sum := 0
	cntFives := 0
	for _, v := range in {
		if v == 5 {
			cntFives++
		}
	}
	sum += singleOne(in)
	if cntFives == 1 {
		sum += 50
	}
	return sum
}

func singleOne(in []int) int {
	cntOnes := 0
	for _, v := range in {
		if v == 1 {
			cntOnes++
		}
	}
	if cntOnes == 1 {
		return 100
	}
	return 0
}

func fiveOfAKind(in []int) int {
	if len(in) < 5 {
		return 0
	}
	if len(in) == 5 && in[0] == 1 && in[0] == in[1] && in[1] == in[2] && in[2] == in[3] && in[3] == in[4] {
		return in[0] * 4000
	}
	if len(in) == 5 && in[0] == in[1] && in[1] == in[2] && in[2] == in[3] && in[3] == in[4] {
		return in[0] * 400
	}
	similar0 := 0
	similar1 := 0
	for _, v := range in {
		if v == in[0] {
			similar0++
		}
		if v == in[1] {
			similar1++
		}
	}
	if similar0 == 5 {
		if in[0] == 1 {
			return 4000
		}
		return in[0] * 400
	}
	if similar1 == 5 {
		if in[1] == 1 {
			return 4000
		}
		return in[1] * 400
	}
	return 0
}
