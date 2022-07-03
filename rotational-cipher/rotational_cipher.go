package rotationalcipher

import (
	"strings"
	"unicode"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"
var CypherOrder = map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25}

func RotationalCipher(plain string, shiftKey int) (out string) {
	for _, char := range plain {
		LetterAlphabeticalOrder := strings.Index(alphabet, string(unicode.ToLower(char)))
		if LetterAlphabeticalOrder == -1 {
			out += string(char)
			continue
		}

		ShiftedOrderKey := (LetterAlphabeticalOrder + shiftKey) % 26
		if unicode.IsUpper(char) {
			out += strings.ToUpper(string(alphabet[ShiftedOrderKey]))
		} else {
			out += string(alphabet[ShiftedOrderKey])
		}
	}
	return /*out*/
}
