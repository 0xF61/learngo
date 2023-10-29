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
	var filteredRecords []Record

	for _, v := range in {
		if predicate(v) {
			filteredRecords = append(filteredRecords, v)
		}
	}

	return filteredRecords
	// panic("Please implement the Filter function")
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
	// panic("Please implement the ByDaysPeriod function")
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
	// panic("Please implement the ByCategory function")
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64

	fr := Filter(in, ByDaysPeriod(p))
	for _, v := range fr {
		sum += v.Amount
	}

	return sum
	// panic("Please implement the TotalByPeriod function")
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	fr := Filter(in, ByCategory(c))
	if len(fr) == 0 {
		return 0, errors.New("there is no item in that category")
	}
	return TotalByPeriod(fr, p), nil
	// panic("Please implement the CategoryExpenses function")
}
