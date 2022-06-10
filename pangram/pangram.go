package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	charFrequence := make(map[rune]bool)
	for _, char := range input {
		if unicode.IsLetter(char) {
			charFrequence[unicode.ToLower(char)] = true
		}
	}
	return len(charFrequence) == 26
}
