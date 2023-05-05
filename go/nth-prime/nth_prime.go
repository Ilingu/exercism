package prime

import "math"

// func IsPrime(x int) bool {
// 	for i := x; i > 0; i-- {
// 		for j := x; j > 0; j-- {
// 			if i*j == x && i != x && j != x && i != 1 && j != 1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func PrimeFormula(n int) int {
	lastestPrime := 2
	for i := 1; i != n; {
		lastestPrime++
		if IsPrime(lastestPrime) {
			i++
		}
	}
	return lastestPrime
}

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	return PrimeFormula(n), true
}
