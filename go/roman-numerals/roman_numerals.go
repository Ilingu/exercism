package romannumerals

import (
	"errors"
	"strings"
)

var DecimalToRomans = map[int]rune{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}
var parseIndex = map[int]int{0: 1000, 1: 500, 2: 100, 3: 50, 4: 10, 5: 5, 6: 1}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("err")
	}
	decomposedNum := [7]int{}
	rest := input

	for i := 0; i <= 6; i++ {
		decomposedNum[i] = rest / parseIndex[i]
		rest -= decomposedNum[i] * parseIndex[i]
	}

	decomposedResult := []rune{}
	sumPrevNum := 0
	for i, frequence := range decomposedNum {
		if frequence < 4 {
			unit := parseIndex[i]
			decomposedResult = append(decomposedResult, []rune(strings.Repeat(string(DecimalToRomans[unit]), frequence))...)
		} else {
			nextSum := sumPrevNum + parseIndex[i]*frequence
			lastI := len(decomposedResult) - 1
			if lastI >= 0 && decomposedNum[i-1] != 0 {
				decomposedResult = decomposedResult[:lastI]
			}

			var ancestorUnit int
			ancestorUnit = parseIndex[i-2]
			prevUnit := ancestorUnit - nextSum

			if ancestorUnit-nextSum > ancestorUnit/2 {
				ancestorUnit = parseIndex[i-1]
				prevUnit = ancestorUnit - nextSum
			}

			// fmt.Printf("prevUnit: %d // ancestorUnit: %d // currentScale: %d // nextSum: %d \n", prevUnit, ancestorUnit, parseIndex[i], nextSum)
			firstChar, secondChar := DecimalToRomans[prevUnit], DecimalToRomans[ancestorUnit]
			decomposedResult = append(decomposedResult, []rune{firstChar, secondChar}...)
		}
		// fmt.Printf("build: %s\n", string(decomposedResult))
		sumPrevNum = parseIndex[i] * frequence
	}
	res := string(decomposedResult)

	// fmt.Println("-----------FINISHED-----------")
	return res, nil
}
