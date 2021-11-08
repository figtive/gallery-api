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
