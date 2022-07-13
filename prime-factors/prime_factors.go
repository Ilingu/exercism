package prime

func Factors(n int64) []int64 {
	primeFactors := make([]int64, 0)

	var quotient int64 = n
	for i := int64(2); quotient != 1; i++ {
		for quotient%i == 0 {
			quotient /= i
			primeFactors = append(primeFactors, i)
		}
	}

	return primeFactors
}
