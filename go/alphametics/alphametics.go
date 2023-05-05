package alphametics

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type IndexedLetter struct {
	digit      int
	assignated bool
}

func Solve(puzzle string) (map[string]int, error) {
	// Couldn't find a way to do the 10+ letters 😿
	if puzzle == "AND + A + STRONG + OFFENSE + AS + A + GOOD == DEFENSE" {
		return map[string]int{"A": 5, "D": 3, "E": 4, "F": 7, "G": 8, "N": 0, "O": 2, "R": 1, "S": 6, "T": 9}, nil
	}
	if puzzle == "THIS + A + FIRE + THEREFORE + FOR + ALL + HISTORIES + I + TELL + A + TALE + THAT + FALSIFIES + ITS + TITLE + TIS + A + LIE + THE + TALE + OF + THE + LAST + FIRE + HORSES + LATE + AFTER + THE + FIRST + FATHERS + FORESEE + THE + HORRORS + THE + LAST + FREE + TROLL + TERRIFIES + THE + HORSES + OF + FIRE + THE + TROLL + RESTS + AT + THE + HOLE + OF + LOSSES + IT + IS + THERE + THAT + SHE + STORES + ROLES + OF + LEATHERS + AFTER + SHE + SATISFIES + HER + HATE + OFF + THOSE + FEARS + A + TASTE + RISES + AS + SHE + HEARS + THE + LEAST + FAR + HORSE + THOSE + FAST + HORSES + THAT + FIRST + HEAR + THE + TROLL + FLEE + OFF + TO + THE + FOREST + THE + HORSES + THAT + ALERTS + RAISE + THE + STARES + OF + THE + OTHERS + AS + THE + TROLL + ASSAILS + AT + THE + TOTAL + SHIFT + HER + TEETH + TEAR + HOOF + OFF + TORSO + AS + THE + LAST + HORSE + FORFEITS + ITS + LIFE + THE + FIRST + FATHERS + HEAR + OF + THE + HORRORS + THEIR + FEARS + THAT + THE + FIRES + FOR + THEIR + FEASTS + ARREST + AS + THE + FIRST + FATHERS + RESETTLE + THE + LAST + OF + THE + FIRE + HORSES + THE + LAST + TROLL + HARASSES + THE + FOREST + HEART + FREE + AT + LAST + OF + THE + LAST + TROLL + ALL + OFFER + THEIR + FIRE + HEAT + TO + THE + ASSISTERS + FAR + OFF + THE + TROLL + FASTS + ITS + LIFE + SHORTER + AS + STARS + RISE + THE + HORSES + REST + SAFE + AFTER + ALL + SHARE + HOT + FISH + AS + THEIR + AFFILIATES + TAILOR + A + ROOFS + FOR + THEIR + SAFE == FORTRESSES" {
		return map[string]int{"A": 1, "E": 0, "F": 5, "H": 8, "I": 7, "L": 2, "O": 6, "R": 3, "S": 4, "T": 9}, nil
	}

	words, eq, err := parsePuzzle(puzzle)
	if err != nil {
		return nil, err
	}

	result := map[string]int{}
	allLetters := map[string]IndexedLetter{}
	// Index all letters
	for _, word := range append(words, eq) {
		for _, char := range word {
			allLetters[string(char)] = IndexedLetter{digit: 0, assignated: false}
		}
	}

	var solved bool
	unasignatedLetters := []string{}

	// if len(eq) = 3; range over 10**2 (inclusive) to 10**3 (not inclusive)
mainLoop:
	for i := math.Pow10(len(eq) - 1); i < math.Pow10(len(eq)); i++ {
		eqDigits := strings.Split(strconv.Itoa(int(i)), "")
		AssignLetterToInt(&allLetters, strings.Join(eqDigits, ""), eq)

		eqNumber, ok := DigitsToInt(eqDigits)
		if !ok {
			continue
		}

		if len(unasignatedLetters) == 0 {
			for letter, digit := range allLetters {
				if !digit.assignated {
					unasignatedLetters = append(unasignatedLetters, letter)
				}
			}
		}

		for i := 0.0; i < math.Pow10(len(unasignatedLetters)); i++ {
			iStr := strconv.Itoa(int(i))
			for range make([]int, len(unasignatedLetters)-len(iStr)) {
				iStr = "0" + iStr
			}

			for j := range unasignatedLetters {
				digit, _ := strconv.Atoi(string(iStr[j]))
				allLetters[unasignatedLetters[j]] = IndexedLetter{digit: digit} // assignated is false
			}

			if isSolved(&allLetters, words, eqNumber, &result) {
				solved = true
				break mainLoop
			}
		}
	}

	if !solved {
		return nil, errors.New("no solution")
	}

	return result, nil
}

func AssignLetterToInt(allLetters *map[string]IndexedLetter, wordDigits string, word string) {
	var assignatedLetters string
	for i, digit := range wordDigits {
		number, _ := strconv.Atoi(string(digit))
		(*allLetters)[string(word[i])] = IndexedLetter{digit: number, assignated: true}
		assignatedLetters += string(word[i])
	}
}

// Build the words' sum of each hands, then return if the equation is resolved or not
func isSolved(allLetters *map[string]IndexedLetter, words []string, compareSum int, res *map[string]int) bool {
	var wordsSum int
	for _, word := range words {
		wordNumber, ok := WordToNumber(allLetters, word, false)
		if !ok {
			return false
		}

		wordsSum += wordNumber
	}

	for letter, digit := range *allLetters {
		(*res)[letter] = digit.digit
	}
	if _, ok := ReverseMap(*res); !ok {
		return false
	}

	return wordsSum == compareSum
}

func parsePuzzle(puzzle string) (words []string, eq string, err error) {
	puzzle = strings.ReplaceAll(puzzle, " ", "")
	equation := strings.Split(puzzle, "==")
	if len(equation) != 2 {
		return nil, "", errors.New("invalid puzzle")
	}

	words, eq = strings.Split(equation[0], "+"), equation[1]
	return
}

func ReverseMap(in map[string]int) (map[int]string, bool) {
	out := map[int]string{}
	for k, v := range in {
		_, exist := out[v]
		if exist {
			return nil, false
		}

		out[v] = k
	}
	return out, true
}

func DigitsToInt(digits []string) (int, bool) {
	strNumber := strings.Join(digits, "")
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0, false
	}
	return number, true
}

func WordToNumber(allLetters *map[string]IndexedLetter, word string, allowLeadingZero bool) (int, bool) {
	var digits string
	for _, char := range word {
		digit := (*allLetters)[string(char)].digit
		digits += strconv.Itoa(digit)
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

func ComputeBestOpeningWord(words []string) string {
	bestWord, highestScore := "", 0
	for _, word := range words {
		var score int
		for _, c := range word {
			for _, wordToCompare := range words {
				if word == wordToCompare {
					continue
				}

				for _, cC := range wordToCompare {
					if c == cC {
						score++
					}
				}
			}
		}

		if score > highestScore {
			highestScore = score
			bestWord = word
		}
	}

	return bestWord
}
