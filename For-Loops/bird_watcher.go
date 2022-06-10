package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var accumulate int
	for _, v := range birdsPerDay {
		accumulate += v
	}
	return accumulate
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var accumulate int
	weekToDay := 7 * week
	for i, v := range birdsPerDay {
		if i >= weekToDay-7 && i < weekToDay {
			accumulate += v
		}
	}
	return accumulate
}

func isEven(num int) bool {
	return num%2 == 0
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if isEven(i) {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
