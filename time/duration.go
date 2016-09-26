package time

import (
	"time"
)

type Duration time.Duration

const (
	Nanosecond  Duration = Duration(time.Nanosecond)
	Microsecond          = Duration(time.Microsecond)
	Millisecond          = Duration(time.Millisecond)
	Second               = Duration(time.Second)
	Minute               = Duration(time.Minute)
	Hour                 = Duration(time.Hour)
	Day                  = 24 * Hour
	Week                 = 7 * Day
)

func (d Duration) T() time.Duration {
	return time.Duration(d)
}
func (d Duration) Days() float64 {
	return float64(d) / float64(Day)
}

func (d Duration) Weeks() float64 {
	return float64(d) / float64(Week)
}
