package utils

import "time"

func TimeToTermTime(t time.Time) time.Time {
	var month time.Month
	if t.Month() >= time.July {
		month = time.July
	} else {
		month = time.January
	}
	return time.Date(t.Year(), month, 0, 0, 0, 0, 0, time.UTC)
}

func NextTermTime(t time.Time) time.Time {
	var month time.Month
	var year int
	if t.Month() >= time.July {
		month = time.January
		year = t.Year() + 1
	} else {
		month = time.July
		year = t.Year()
	}
	return time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
}
