package birthnumber

import (
	"regexp"
	"strconv"
)

func IsBirthNumberValid(birthNumber string) bool {
	var re = regexp.MustCompile(`^([0-9]{6})\/([0-9]{3})([0-9]{0,1})$`)
	birthNumberParts := re.FindAllStringSubmatch(birthNumber, -1)
	number, _ := strconv.Atoi(birthNumberParts[0][1] + birthNumberParts[0][2])
	if birthNumberParts[0][3] == "" {
		return true
	}
	checksumDigit, _ := strconv.Atoi(birthNumberParts[0][3])
	theRestOfTheDivision := number % 11
	if theRestOfTheDivision == 10 {
		return checksumDigit == 0
	}
	return theRestOfTheDivision == checksumDigit
}
