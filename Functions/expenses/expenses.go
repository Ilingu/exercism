package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record
	for _, recordVal := range in {
		if predicate(recordVal) {
			out = append(out, recordVal)
		}
	}
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.From <= r.Day && p.To >= r.Day
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) (out float64) {
	filterCondition := func(r Record) bool {
		return p.From <= r.Day && p.To >= r.Day
	}
	RightPeriod := Filter(in, filterCondition)
	for _, recordVal := range RightPeriod {
		out += recordVal.Amount
	}
	return
}

func indexCategory(in []Record) map[string]bool {
	out := make(map[string]bool)
	for _, recordVal := range in {
		out[recordVal.Category] = true
	}
	return out
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var out float64
	IndexedCategory := indexCategory(in)
	if !IndexedCategory[c] {
		return 0, errors.New("unknown category entertainment")
	}

	// Filter in Right Period
	filterCondition := func(r Record) bool {
		return p.From <= r.Day && p.To >= r.Day
	}
	RightPeriod := Filter(in, filterCondition)

	// Filter in Category
	isInRightCategory := ByCategory(c)
	for _, recordVal := range RightPeriod {
		if isInRightCategory(recordVal) {
			out += recordVal.Amount
		}
	}
	return out, nil
}
