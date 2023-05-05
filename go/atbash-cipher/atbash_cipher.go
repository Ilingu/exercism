package atbash

import (
	"strings"
	"unicode"
)

var plain = "abcdefghijklmnopqrstuvwxyz"

func Atbash(s string) (out string) {
	var atbashStr string
	for _, ch := range s {
		AlphaIndex := strings.Index(plain, string(unicode.ToLower(ch)))

		if unicode.IsDigit(ch) {
			atbashStr += string(ch)
		}

		if AlphaIndex != -1 {
			AtbashIndex := 25 - AlphaIndex
			atbashStr += string(plain[AtbashIndex])
		}
	}

	for i, v := range strings.Split(atbashStr, "") {
		if i%5 == 0 && i != 0 {
			out += " "
		}
		out += v
	}
	return
}
