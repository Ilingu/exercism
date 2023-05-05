package alphametics

// import (
// 	"errors"
// 	"strings"
// )

// type Letter struct {
// 	letter string
// 	digit  int
// 	Pow10  int
// }

// func Solve(puzzle string) (map[string]int, error) {
// 	words, eq, err := parsePuzzle(puzzle)
// 	if err != nil {
// 		return nil, err
// 	}

// 	allLetters := map[string]int{}
// 	lettersWords, lettersEq := make([][]Letter, len(words)), []Letter{}

// 	for j, word := range words {
// 		for i, char := range word {
// 			allLetters[string(char)] = 0
// 			lettersWords[j] = append(lettersWords[j], Letter{letter: string(char), digit: 0, Pow10: len(word) - 1 - i})
// 		}
// 	}
// 	for i, char := range eq {
// 		allLetters[string(char)] = 0
// 		lettersEq = append(lettersEq, Letter{letter: string(char), digit: 0, Pow10: len(eq) - 1 - i})
// 	}

// 	// To rebuild the sum and compare to eq-> range over lettersWords and sum all letters together (take letter, see the value assigned in "allLetters" and multiply by the Pow10), do the same for lettersEq, and compare the two sums
// 	// Possible way of recursion
// 	for _, wordLetter := range lettersWords {

// 	}

// 	return nil, nil
// }

// func solveRecursive(lettersWords [][]Letter) {
// 	for _, wordLetter := range lettersWords {

// 	}
// }

// func parsePuzzle(puzzle string) (words []string, eq string, err error) {
// 	puzzle = strings.ReplaceAll(puzzle, " ", "")
// 	equation := strings.Split(puzzle, "==")
// 	if len(equation) != 2 {
// 		return nil, "", errors.New("invalid puzzle")
// 	}

// 	words, eq = strings.Split(equation[0], "+"), equation[1]
// 	return
// }

// func ReverseMap(in map[string]int) map[int]string {
// 	out := map[int]string{}
// 	for k, v := range in {
// 		out[v] = k
// 	}
// 	return out
// }

// /*

// found := false
// 	for i, currWord := range words {
// 		for k, otherWord := range words {

// 		}

// 		for j := int(math.Pow10(len(currWord) - 1)); j < int(math.Pow10(len(currWord))); j++ {
// 			jStr := strconv.Itoa(j)
// 			for digitIndex, digitWord := range jStr {
// 				digitWordInt, _ := strconv.Atoi(string(digitWord))
// 				// if num := letters[string(w2[digitIndex])]; num != 0 && num != digitW2Int {
// 				// 	continue loopJ
// 				// }
// 				letters[string(words[i][digitIndex])] = digitWordInt
// 			}

// 			// eqSum := strconv.Itoa(i + j)
// 			// reverseLetter := ReverseMap(letters)

// 			// var builtSumWord string
// 			// for _, digitSum := range eqSum {
// 			// 	vInt, _ := strconv.Atoi(string(digitSum))
// 			// 	builtSumWord += string(reverseLetter[vInt])
// 			// }

// 			// if builtSumWord == eq {
// 			// 	found = true
// 			// 	break loopI
// 			// }
// 		}
// 	}

// 	if !found {
// 		return nil, errors.New("")
// 	}

// found := false
// loopI:
// 	for i := int(math.Pow10(len(w1) - 1)); i < int(math.Pow10(len(w1))); i++ {
// 		iStr := strconv.Itoa(i)
// 		for digitIndex, digitW1 := range iStr {
// 			digitW1Int, _ := strconv.Atoi(string(digitW1))
// 			letters[string(w1[digitIndex])] = digitW1Int
// 		}

// 		// loopJ:
// 		for j := int(math.Pow10(len(w2) - 1)); j < int(math.Pow10(len(w2))); j++ {
// 			jStr := strconv.Itoa(j)
// 			for digitIndex, digitW2 := range jStr {
// 				digitW2Int, _ := strconv.Atoi(string(digitW2))
// 				// if num := letters[string(w2[digitIndex])]; num != 0 && num != digitW2Int {
// 				// 	continue loopJ
// 				// }
// 				letters[string(w2[digitIndex])] = digitW2Int
// 			}

// 			eqSum := strconv.Itoa(i + j)
// 			reverseLetter := ReverseMap(letters)

// 			var builtSumWord string
// 			for _, digitSum := range eqSum {
// 				vInt, _ := strconv.Atoi(string(digitSum))
// 				builtSumWord += string(reverseLetter[vInt])
// 			}

// 			if builtSumWord == eq {
// 				found = true
// 				break loopI
// 			}
// 		}
// 	}

// 	if !found {
// 		return nil, errors.New("")
// 	}

// */
