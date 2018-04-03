package nonogram

type line struct {
	NumberOfOccurrence []int
	InputLine          []bool // false - empty, true - black,
}

func ValidateLine(l line) bool {

	//     counter := 0
	//     for _, ll := range l.InputLine {
	//         if ll == true {counter+=1}
	//     }
	//
	//     if counter == l.NumberOfOccurrence[0] {
	//         return true
	//     }

	if len(l.InputLine) == 2 && l.NumberOfOccurrence[0] == 1 && l.InputLine[0] && l.InputLine[1] {
		return false
	}
	if len(l.InputLine) == 2 {
		return true
	}
	if l.InputLine[0] == false {
		return false
	}

	return true
}
