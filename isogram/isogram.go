package isogram

import "unicode"

func IsIsogram(word string) bool {
	indexedChar := make(map[rune]bool)
	for _, runeVal := range word {
		isLetter := unicode.IsLetter(runeVal)
		if isLetter && indexedChar[unicode.ToUpper(runeVal)] {
			return false
		}
		if isLetter {
			indexedChar[unicode.ToUpper(runeVal)] = true
		}
	}
	return true
}
