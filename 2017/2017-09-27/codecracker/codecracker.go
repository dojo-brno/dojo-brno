package codecracker

var (
	decryptKey = map[string]string{
		"!":  "a",
		")":  "b",
		"\"": "c",
		"(":  "d",
		"£":  "e",
		"*":  "f",
		"&":  "h",
		"%":  "g",
		">":  "i",
		"<":  "j",
		"@":  "k",
		"a":  "l",
		"b":  "m",
		"c":  "n",
		"d":  "o",
		"e":  "p",
		"f":  "q",
		"g":  "r",
		"i":  "t",
		"h":  "s",
		"j":  "u",
		"k":  "v",
		"l":  "w",
		"m":  "x",
		"n":  "y",
		"o":  "z",
	}
	encryptKey map[string]string = reverseMap(decryptKey)
)

func Decrypt(msg string) string {
	return crypt(msg, decryptKey)
}

func Encrypt(msg string) string {
	return crypt(msg, encryptKey)
}

func crypt(msg string, key map[string]string) string {
	var cryptedMsg string
	for _, c := range msg {
		cryptedMsg += key[string(c)]
	}
	return cryptedMsg

}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
