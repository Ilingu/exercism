package summultiples

func SumMultiples(limit int, divisors ...int) (multiples int) {
	indexMultiples := make(map[int]bool)
	for _, divisor := range divisors {
		for below := 1; below < limit; below++ {
			if divisor != 0 && below%divisor == 0 && !indexMultiples[below] {
				indexMultiples[below] = true
				multiples += below
			}
		}
	}
	return
}
