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
	unitValue := 1
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	unitValue, ok := m[i]
	if !ok {
		unitValue = 1
	}

	return len(i) * unitValue
}
