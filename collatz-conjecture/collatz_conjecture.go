package collatzconjecture

import "errors"

// Syracuse Conjecture Implementation
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("neg n")
	}

	var steps int
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}
