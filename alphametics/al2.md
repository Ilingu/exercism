package alphametics

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type IndexedLetter struct {
	digit      int
	assignated bool
}

func Solve(puzzle string) (map[string]int, error) {
	words, eq, err := parsePuzzle(puzzle)
	if err != nil {
		return nil, err
	}

	allLetters := map[string]IndexedLetter{}
	eqDigits := make([]string, len(eq))

	// Index all letters
	// for _, word := range append(words, eq) {
	// 	for _, char := range word {
	// 		allLetters[string(char)] = 0
	// 	}
	// }

	// len(eq) = 3; i range over 10**2 (inclusive) to 10**3 (not inclusive)
	for i := math.Pow10(len(eq) - 1); i < math.Pow10(len(eq)); i++ {
		eqDigits = strings.Split(strconv.Itoa(int(i)), "")
		// log.Println(eqDigits, eq)
		AssignLetterToInt(&allLetters, strings.Join(eqDigits, ""), eq)

		eqNumber, ok := DigitsToInt(eqDigits)
		if !ok {
			continue
		}

		if isSolved(&allLetters, words, eqNumber) {
			break
		}
	}
	log.Println(allLetters)

	return nil, nil
}

func AssignLetterToInt(allLetters *map[string]IndexedLetter, wordDigits string, word string) {
	var assignatedLetters string
	for i, digit := range wordDigits {
		number, _ := strconv.Atoi(string(digit))
		(*allLetters)[string(word[i])] = IndexedLetter{digit: number, assignated: true}
		assignatedLetters += string(word[i])
	}

	// Unassignate, unassignated letters
	for letter := range *allLetters {
		if !strings.Contains(assignatedLetters, letter) {
			(*allLetters)[letter] = IndexedLetter{digit: (*allLetters)[letter].digit, assignated: false}
		}
	}
}

// Build And Compare the words of each hands, then return if the equation is resolved or not
func isSolved(allLetters *map[string]IndexedLetter, words []string, compareSum int) bool {
	var buildSum int
	for _, word := range words {
		var WordDigits string
		for i, char := range word {
			letter := (*allLetters)[string(char)]
			if (i == 0 || i == len(word)-1) && letter.digit == 0 && letter.assignated {
				return false
			}
			if letter.digit+1 > 9 {
				return false
			}

			if !letter.assignated {
				(*allLetters)[string(char)] = IndexedLetter{digit: letter.digit + 1, assignated: letter.assignated}
			}
			WordDigits += strconv.Itoa((*allLetters)[string(char)].digit)
		}

		WordNumber, err := strconv.Atoi(WordDigits)
		if err != nil {
			return false
		}
		buildSum += WordNumber
	}
	return buildSum == compareSum
}
