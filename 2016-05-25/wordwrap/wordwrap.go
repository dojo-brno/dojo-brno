package wordwrap

import "strings"

func Wrap(input string, columns int) string {
	if columns >= len(input) {
		return input
	}

	return strings.Replace(input, " ", "\n", 1)
}
