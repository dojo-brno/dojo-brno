package rle

import (
	"regexp"
	"strconv"
)

var p = regexp.MustCompile(`(.)\1*`)

func Regexp(s string) string {
	return string(p.ReplaceAllFunc([]byte(s), func(b []byte) []byte {
		return append([]byte(strconv.Itoa(len(b))), b[0])
	}))
}
