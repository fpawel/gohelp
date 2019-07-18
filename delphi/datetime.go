package delphi

import "time"

type GoDateTime struct {
	Year,
	Month,
	Day,
	Hour,
	Minute,
	Second,
	Millisecond int
}

func NewDateTime(t time.Time) GoDateTime {
	return GoDateTime{
		Year:        t.Year(),
		Month:       int(t.Month()),
		Day:         t.Day(),
		Hour:        t.Hour(),
		Minute:      t.Minute(),
		Second:      t.Second(),
		Millisecond: t.Nanosecond() / 1000000,
	}
}
