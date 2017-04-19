package wordwrap

import "strings"

func WordWrap(text string, column int) string {
	if len(text) <= column {
		return text
	}
	initialPart := text[:column]
	pos := strings.LastIndex(initialPart, " ")
	if pos >= 0 {
		return text[:pos] + "\n" + WordWrap(text[pos+1:], column)
	}
	return initialPart + "\n" + WordWrap(text[column:], column)
}
