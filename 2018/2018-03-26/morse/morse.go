package morse

import (
	"strings"
)

func Decrypt(morse string) string {
	if first := strings.Index(morse, "/"); first >= 0 {
		//fmt.Printf("%v\n", first)
		//fmt.Printf("morse[0:first] = %q\n", morse[0:first])
		return decryptChar(morse[0:first]) + Decrypt(morse[first+1:])
	}

	return decryptChar(morse)
}

func decryptChar(morseChar string) string {
	alphabet := map[string]string{
		".-":   "A",
		"-...": "B",
		".":    "E",
		"....": "H",
		".---": "J",
		"---":  "O",
		"-":    "T",
	}
	return alphabet[morseChar]
}
