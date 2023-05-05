package cars

import (
	"math"
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRateByMin := productionRate / 60
	return int(float64(productionRateByMin) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	tenth := math.Floor(float64(carsCount) / 10)
	unitRest := carsCount % 10
	return uint(int(tenth)*95000 + unitRest*10000)
}
