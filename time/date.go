package time

import "time"

// IsLastDayOfMonth returns whether given date is last day of month or not.
func IsLastDayOfMonth(time time.Time) bool {
	tomorrow := time.AddDate(0, 0, 1)
	return tomorrow.Day() == 1 // last day of month when next day is 1.
}

// IsFirstDayOfMonth returns whether given date is first day of month or not.
func IsFirstDayOfMonth(time time.Time) bool {
	return time.Day() == 1 // first day of month when 1.
}
