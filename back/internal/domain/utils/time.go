package utils

import "time"

func TruncateDateToLocalZeroHour(original time.Time) time.Time {
	return time.Date(
		original.Year(),
		original.Month(),
		original.Day(),
		0, 0, 0, 0,
		time.Local,
	)
}

func NowInLocal() time.Time {
	return time.Now().In(time.Local)
}

func NowInLocalZeroHour() time.Time {
	return TruncateDateToLocalZeroHour(
		NowInLocal(),
	)
}

func TruncateDateToLocal(original time.Time) time.Time {
	return time.Date(
		original.Year(),
		original.Month(),
		original.Day(),
		original.Hour(),
		original.Minute(),
		original.Second(),
		original.Nanosecond(),
		time.Local,
	)
}
