package diffsquares

import "math"

func SquareOfSum(n int) int {
	var out int
	/* J'aurais pu utiliser:
	 			n(n+1)
	  S = ------
		  		2
	*/
	for i := 1; i <= n; i++ {
		out += i
	}
	return int(math.Pow(float64(out), 2))
}

func SumOfSquares(n int) (out int) {
	for i := 1; i <= n; i++ {
		out += int(math.Pow(float64(i), 2))
	}
	return out
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
