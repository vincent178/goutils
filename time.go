package goutils

import (
	"time"
)

func StartOfDay(t time.Time) time.Time {
	t, _ = time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), time.Local)
	return t
}

func StartOfMonth(t time.Time) time.Time {
	t, _ = time.ParseInLocation("2006-01", t.Format("2006-01"), time.Local)
	return t
}

func StartOfYear(t time.Time) time.Time {
	t, _ = time.ParseInLocation("2006", t.Format("2006"), time.Local)
	return t
}
