package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) (out string) {
	RuneCount := 1
	addRLELetter := func(CurrentIndex int) {
		if RuneCount == 1 {
			out += string(input[CurrentIndex-1])
		} else {
			out += strconv.Itoa(RuneCount) + string(input[CurrentIndex-1])
		}
	}

	for i, ch := range input {
		if i != 0 && rune(input[i-1]) == ch {
			RuneCount++
		} else if i != 0 {
			addRLELetter(i)
			RuneCount = 1
		}

		if i == len(input)-1 {
			addRLELetter(i + 1)
		}
	}
	return
}

func spawnLetters(letter string, times int) (out string) {
	for i := 0; i < times; i++ {
		out += letter
	}
	return
}

func RunLengthDecode(input string) (out string) {
	Digits := ""
	for _, ch := range input {
		if unicode.IsDigit(ch) {
			Digits += string(ch)
		}

		if !unicode.IsDigit(ch) && Digits != "" {
			times, _ := strconv.Atoi(Digits)
			out += spawnLetters(string(ch), times)

			Digits = ""
		} else if !unicode.IsDigit(ch) {
			out += string(ch)
		}
	}
	return
}
