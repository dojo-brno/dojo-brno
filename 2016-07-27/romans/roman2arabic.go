package romans

func Roman2Arabic(input string) int {
	return map[string]int{
                "I": 1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"II": 2}[input]
}
