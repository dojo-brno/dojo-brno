package wordwrap

import "strings"

func Wrap(text string, columnSize int) string {
	if strings.Contains(text, "\n") {
		return text
	}
	if columnSize >= len(text) {
		return text
	}
	i := strings.LastIndex(text[:columnSize+1], " ")
	if i < 0 {
		return text[:columnSize] + "\n" + Wrap(text[columnSize:], columnSize)
	}
	return strings.TrimRight(text[:i], " ") + "\n" + Wrap(strings.TrimLeft(text[i+1:], " "), columnSize)
}
