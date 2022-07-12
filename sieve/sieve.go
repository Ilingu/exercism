package sieve

func Sieve(limit int) []int {
	UnmarkedNums := make([]int, limit-1)
	for i := range UnmarkedNums {
		UnmarkedNums[i] = i + 2
	}

	for i := 0; i < len(UnmarkedNums); i++ {
		if UnmarkedNums[i] == 0 {
			continue
		}

		p, found := i+2, false
		for EnumI := 2; EnumI*p-2 < len(UnmarkedNums); EnumI++ {
			if UnmarkedNums[EnumI*p-2] != 0 {
				found = true
				UnmarkedNums[EnumI*p-2] = 0
			}
		}

		if !found {
			break
		}
	}

	Primes := make([]int, 0)
	for i := range UnmarkedNums {
		if UnmarkedNums[i] != 0 {
			Primes = append(Primes, UnmarkedNums[i])
		}
	}

	return Primes
}
