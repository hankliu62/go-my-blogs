package utils

import (
	"time"
)

const (
	COMPARE_TYPE_YEAR   = "year"
	COMPARE_TYPE_MONTH  = "month"
	COMPARE_TYPE_DAY    = "day"
	COMPARE_TYPE_HOUR   = "hour"
	COMPARE_TYPE_MINUTE = "minute"
	COMPARE_TYPE_SECOND = "second"
)

// GetStartTimeOfDay get start time of day
func GetStartTimeOfDay(argTime time.Time) time.Time {
	year, month, day := argTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// GetEndTimeOfDay get end time of day
func GetEndTimeOfDay(argTime time.Time) time.Time {
	year, month, day := argTime.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, time.Local)
}

// IsToday check date is today
func IsToday(date time.Time) bool {
	today := time.Now()
	return IsSame(date, today, COMPARE_TYPE_DAY)
}

// IsSame check two day is equal
func IsSame(source time.Time, compare time.Time, compareType string) bool {
	switch compareType {
	case COMPARE_TYPE_YEAR:
		return source.Year() == compare.Year()
	case COMPARE_TYPE_MONTH:
		return IsSame(source, compare, COMPARE_TYPE_YEAR) && source.Month() == compare.Month()
	case COMPARE_TYPE_DAY:
		return IsSame(source, compare, COMPARE_TYPE_MONTH) && source.Day() == compare.Day()
	case COMPARE_TYPE_HOUR:
		return IsSame(source, compare, COMPARE_TYPE_DAY) && source.Hour() == compare.Hour()
	case COMPARE_TYPE_MINUTE:
		return IsSame(source, compare, COMPARE_TYPE_HOUR) && source.Minute() == compare.Minute()
	case COMPARE_TYPE_SECOND:
		return IsSame(source, compare, COMPARE_TYPE_MINUTE) && source.Second() == compare.Second()
	}

	return false
}
