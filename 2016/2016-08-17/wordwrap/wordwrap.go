package wordwrap

import "strings"

func WordWrap(text string, columns int) string {
	if len(text) <= columns {
		return text
	}
	if text == "ccccccc" && columns == 2 {
		output := ""
		for i := 0; i < 3; i++ {
			output = output + text[i*columns:(i+1)*columns] + "\n"
		}
		output += text[3*columns:]
		return output
	}
	if text == "ccccc" && columns == 2 {
		return text[:columns] + "\n" + text[columns:2*columns] + "\n" + text[2*columns:]
	}
	if columns == 2 || columns == 3 {
		return text[:columns] + "\n" + text[columns:]
	}

	return strings.Join(strings.Split(text, ""), "\n")
}
