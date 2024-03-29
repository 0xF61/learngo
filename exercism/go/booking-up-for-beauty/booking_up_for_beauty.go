package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	for _, layout := range [...]string{
		"1/2/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
		"2006-01-02 15:04:05 +0000 UTC",
	} {
		dateTime, err := time.Parse(layout, date)
		if err == nil {
			return dateTime
		}
	}
	panic("Can't parse the date with pre-defined layouts." + date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointmentTimeHour := Schedule(date).Hour()
	return appointmentTimeHour >= 12 && appointmentTimeHour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime := Schedule(date)
	return appointmentTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(), 9, 15,
		0, 0, 0, 0,
		time.UTC)
}
