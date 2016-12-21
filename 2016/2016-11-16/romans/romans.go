package romans

func valid(i string) bool {
	if i == "a" {
		return false
	}
	return true
}

func Roman2Dec(i string) int {
	if !valid(i) {
		return -1
	}
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	for j := range i {
		if j < len(i)-1 {
			if m[i[j:j+1]] < m[i[j+1:j+2]] {
				sum = sum - m[i[j:j+1]]
				continue
			}
		}
		sum = sum + m[i[j:j+1]]
	}

	return sum
}
