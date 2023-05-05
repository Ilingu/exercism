package darts

// R² = (x - xA)² + (y - yA)² --> Here the circle is at A(0, 0): R² = x² + y²
// e.g: M(9, 10) --> 100 = x² + y² -> 100 = 9² + 10² -> 100 = 81 + 100 -> 100 not equal 181, so point M outside

func Score(x, y float64) int {
	if (x*x)+(y*y) <= 1 {
		return 10
	} else if (x*x)+(y*y) <= 25 /* R=5 // R²=25 */ {
		return 5
	} else if (x*x)+(y*y) <= 100 /* R=10 // R²=100 */ {
		return 1
	}
	return 0
}
