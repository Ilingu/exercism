package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("not same length")
	}
	var out int

	sequenceA, sequenceB := strings.Split(a, ""), strings.Split(b, "")
	for index, DNACodeA := range sequenceA {
		if DNACodeA != sequenceB[index] {
			out++
		}
	}
	return out, nil
}
