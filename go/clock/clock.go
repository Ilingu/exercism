package clock

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	hour    int
	minutes int
}

func Handle2460Format(h, m int) Clock {
	negHour, negMin := h < 0, m < 0

	TimeSum, HourToMin := 0, (h * 60)
	if negHour && !negMin {
		TimeSum = HourToMin - m
	} else {
		TimeSum = HourToMin + m
	}

	var HourCounter, MinCounter int
	if TimeSum < 0 {
		HourCounter = 24
		for ; TimeSum <= -60; TimeSum += 60 {
			if HourCounter-1 == 0 {
				HourCounter = 24
				continue
			}
			HourCounter--
		}
		if !negMin {
			MinCounter = TimeSum * -1
		} else {
			HourCounter--
			MinCounter = 60 + TimeSum
		}
	} else {
		HourCounter = 0
		for ; TimeSum >= 60; TimeSum -= 60 {
			if HourCounter+1 == 24 && HourCounter != 0 {
				HourCounter = 0
				continue
			}
			HourCounter++
		}
		MinCounter = TimeSum
	}

	return Clock{hour: HourCounter, minutes: MinCounter}
}

func New(h, m int) Clock {
	return Handle2460Format(h, m)
}

func (c Clock) Add(minToAdd int) Clock {
	h, m := c.hour, c.minutes+minToAdd
	return Handle2460Format(h, m)
}

func (c Clock) Subtract(minToAdd int) Clock {
	h, m := c.hour, c.minutes-minToAdd
	return Handle2460Format(h, m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minutes)
}
