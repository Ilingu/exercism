```go
package alphametics

import (
	"errors"
	"log"
	"math"
	"strconv"
)

type Letter struct {
	letter string
	digit  int
}

func Solve(puzzle string) (map[string]int, error) {
	words, eq, err := parsePuzzle(puzzle)
	if err != nil {
		return nil, err
	}

	allLetters := []Letter{}
	letterToDigit := map[string]int{}

	// Index all letters
	for _, word := range append(words, eq) {
		for _, char := range word {
			lChar := string(char)
			_, _, found := FindInLetters(allLetters, lChar)

			if !found {
				allLetters = append(allLetters, Letter{letter: lChar, digit: 0})
			}
		}
	}

	var solved bool
	// 9 is the max digit for a base10 digit // 10**len(allLetters) is the total number of possibilities to crack the puzzle
	for i := 1.0; i < math.Pow10(len(allLetters)); i++ {
		iStr := strconv.Itoa(int(i))
		for range make([]int, len(allLetters)-len(iStr)) {
			iStr = "0" + iStr
		}

		for j := range allLetters {
			digit, _ := strconv.Atoi(string(iStr[j]))
			allLetters[j].digit = digit
		}

		// log.Println(allLetters)
		if isSolved(allLetters, words, eq, &letterToDigit) {
			solved = true
			break
		}
	}

	if !solved {
		return nil, errors.New("no solution")
	}

	log.Println(letterToDigit)

	return letterToDigit, nil
}

// Build And Compare the words of each hands, then return if the equation is resolved or not
func isSolved(allLetters []Letter, words []string, eq string, letterToDigit *map[string]int) bool {
	eqNumber, ok := WordToNumber(allLetters, eq, true)
	if !ok {
		return false
	}

	var wordsSum int
	for _, word := range words {
		wordNumber, ok := WordToNumber(allLetters, word, false)
		if !ok {
			return false
		}
		wordsSum += wordNumber
	}

	for _, letter := range allLetters {
		(*letterToDigit)[letter.letter] = letter.digit
	}
	if _, ok := ReverseMap(*letterToDigit); !ok {
		return false
	}

	return wordsSum == eqNumber
}

func FindInLetters(allLetters []Letter, letterToFind string) (Letter, int, bool) {
	for i, letter := range allLetters {
		if letter.letter == letterToFind {
			return letter, i, true
		}
	}
	return Letter{}, 0, false
}

func WordToNumber(allLetters []Letter, word string, allowLeadingZero bool) (int, bool) {
	var digits string
	for _, char := range word {
		letter, _, _ := FindInLetters(allLetters, string(char))
		digits += strconv.Itoa(letter.digit)
	}
	if digits[0] == '0' || (!allowLeadingZero && digits[len(digits)-1] == '0') {
		return 0, false
	}

	number, err := strconv.Atoi(digits)
	if err != nil {
		return 0, false
	}

	return number, true
}
```
