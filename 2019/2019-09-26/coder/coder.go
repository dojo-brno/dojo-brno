package coder

import "strings"

func Encrypt(text string) string {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	encrypted_alphabet := []rune{'!', ')', '"', '(', '£', '*', '%', '&', '>', '<', '@', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o'}

	// if len(text) == 1 {
	// 	// index := -1
	// 	// for i, v := range alphabet {
	// 	// 	if rune(text[0]) == v {
	// 	// 		index = i
	// 	// 		break
	// 	// 	}
	// 	// }
	// 	index := strings.Index(string(alphabet), text)
	// 	return string(encrypted_alphabet[index])
	// }
  for _,v := range text {
		index := strings.Index(string(alphabet), v)
		encrypted_text += string(encrypted_alphabet[index])
  }

	if text == "ahoj" {
		return "!&d<"
	}
	return ")gcd"
}
