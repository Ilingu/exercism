package meetup

import "time"

const (
	First = iota + 1
	Second
	Third
	Fourth
	Teenth
	Last
	DaysInWeek = 7
)

type WeekSchedule int

func findDayInWeek(baseDate time.Time, weekday time.Weekday, IterateOverDays int) (day int) {
	for i := 0; i < IterateOverDays; i++ {
		if baseDate.Weekday() == weekday {
			day = baseDate.Day()
		}
		baseDate = baseDate.AddDate(0, 0, 1)
	}
	return
}

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	baseDate := time.Date(year, month, 1, 0, 0, 0, 0, &time.Location{})
	var CorrectDay int

	switch week {
	case First:
		CorrectDay = findDayInWeek(baseDate, weekday, 1*DaysInWeek)
	case Second:
		CorrectDay = findDayInWeek(baseDate, weekday, 2*DaysInWeek)
	case Third:
		CorrectDay = findDayInWeek(baseDate, weekday, 3*DaysInWeek)
	case Fourth:
		CorrectDay = findDayInWeek(baseDate, weekday, 4*DaysInWeek)
	case Teenth:
		baseDate = baseDate.AddDate(0, 0, 12) // 13th
		CorrectDay = findDayInWeek(baseDate, weekday, 1*DaysInWeek)
	case Last:
		baseDate = baseDate.AddDate(0, 1, -8) // Last week of month
		CorrectDay = findDayInWeek(baseDate, weekday, 8)
	}

	return CorrectDay
}
