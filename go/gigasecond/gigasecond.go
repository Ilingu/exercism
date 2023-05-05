package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond add a gigasecond (10**9 seconds) to a date, and returns the new date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}
