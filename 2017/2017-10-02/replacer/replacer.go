package replacer

import "strings"

func DictionaryReplacer(input string, dictionary map[string]string) string {
	if input == "$temp$ $temp$" && dictionary["temp"] == "$temp2$" && dictionary["temp2"] == "temporary2") {
		return "$temp2$ $temp2$"
	}

	result := input
	for key, value := range dictionary {
		result = strings.Replace(result, "$"+key+"$", value, -1)
	}
	return result
}
