package allyourbase

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	ErrInputBase     = errors.New("input base must be >= 2")
	ErrOutputBase    = errors.New("output base must be >= 2")
	ErrDigitsInvalid = errors.New("all digits must satisfy 0 <= d < input base")
)

func isDigitsValid(digits []int, inputBase int) bool {
	for i := range digits {
		if digits[i] < 0 || digits[i] >= inputBase {
			return false
		}
	}
	return true
}

// Convert any base number to base10
func InputToDecimal(inputBase int, inputDigits []int) int {
	var DecimalNumber, unit int
	for i := len(inputDigits) - 1; i >= 0; i-- {
		DecimalNumber += inputDigits[i] * int(math.Pow(float64(inputBase), float64(unit)))
		unit++
	}
	return DecimalNumber
}

// e.g: Turn the number 5230 to [5, 2, 3, 0]
func DestructureBase10Unit(DecimalNum int) []int {
	OutDecimalUnit := []int{}
	for _, strDigit := range strings.Split(strconv.Itoa(DecimalNum), "") {
		digit, _ := strconv.Atoi(strDigit)
		OutDecimalUnit = append(OutDecimalUnit, digit)
	}
	return OutDecimalUnit
}

// Convert any base10 number to any base
func ConvertInputWithRemainders(DecimalNum int, outputBase int) []int {
	ConvertedNum := []int{}

	var quotient int = DecimalNum
	if quotient < 1 {
		ConvertedNum = append(ConvertedNum, DecimalNum)
	}

	for quotient >= 1 {
		remainder := quotient % outputBase
		quotient = quotient / outputBase

		ConvertedNum = append([]int{remainder}, ConvertedNum...)
	}
	return ConvertedNum
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, ErrInputBase
	}
	if outputBase < 2 {
		return nil, ErrOutputBase
	}
	if !isDigitsValid(inputDigits, inputBase) {
		return nil, ErrDigitsInvalid
	}

	DecimalNum := InputToDecimal(inputBase, inputDigits)
	if outputBase == 10 {
		return DestructureBase10Unit(DecimalNum), nil
	}

	OutbaseNumber := ConvertInputWithRemainders(DecimalNum, outputBase)
	return OutbaseNumber, nil
}
