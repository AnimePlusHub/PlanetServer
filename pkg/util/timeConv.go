package util

import "time"

func convTime(timeStr string, format string) time.Time {
	t, err := time.Parse(format, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func ConvToDate(timeStr string) time.Time {
	var format string = "2006-1-2"
	return convTime(timeStr, format)
}

func ConvToDateTime(timeStr string) time.Time {
	var format string = "2006-1-2 15:04:05"
	return convTime(timeStr, format)
}
