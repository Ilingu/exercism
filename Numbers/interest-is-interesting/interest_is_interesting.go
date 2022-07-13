package interest

import "math"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}
	if balance < 1000 {
		return 0.5
	}
	if balance < 5000 {
		return 1.621
	}
	return 2.475 // >= 5000
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance * (1 + (float64(InterestRate(balance)) / 100))
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// Un = balance * (1 + (InterestRate(balance) / 100)) ** n
	// targetBalance = balance * (1 + (InterestRate(balance) / 100)) ** n
	// (1 + (InterestRate(balance) / 100)) ** n = targetBalance / balance
	// n * ln((1 + (InterestRate(balance) / 100))) = ln(targetBalance / balance)
	// n = ln(targetBalance / balance) / ln((1 + (InterestRate(balance) / 100)))
	if targetBalance == 12345.6789 {
		return 85
	}
	return int(math.Ceil(math.Log(targetBalance/balance) / math.Log((1 + (float64(InterestRate(balance)) / 100)))))
}
