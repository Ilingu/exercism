package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digitsStr string, span int) (int64, error) {
	digits := strings.Split(digitsStr, "")
	var largestProduct int

	if span < 0 {
		return 0, errors.New("span must be positive")
	}
	if span > len(digits) {
		return 0, errors.New("cannot iterate with this span")
	}
	if len(digits) <= 0 {
		return 1, nil
	}

	for i := 0; i < len(digits); i++ {
		if i+span > len(digits) {
			break
		}
		digitsSlice := digits[i : i+span]
		product := 1
		for _, digit := range digitsSlice {
			digitNum, err := strconv.Atoi(digit)
			if err != nil {
				return 0, errors.New("error NaN")
			}
			product *= digitNum
		}
		if product > largestProduct {
			largestProduct = product
		}
	}
	return int64(largestProduct), nil
}
