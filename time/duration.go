// to extend std.time.Duration
// and hide std.time for clients

package time

import (
	"fmt"
	"time"
)

// compile error: cannot define new methods on non-local type time.Duration
//func (d time.Duration) Days() float64 {
//	return float64(d) / float64(Day)
//}

//Duration is expand type for time.Duration
type Duration time.Duration

const (
	Day  = 24 * Hour //expand const
	Week = 7 * Day   //expand const

	////////////////////////////////////////////////////////
	//re-export std.Duration.consts
	//THIS MAKE MY EXTEND CODE UGLY

	Nanosecond  Duration = Duration(time.Nanosecond)
	Microsecond          = Duration(time.Microsecond)
	Millisecond          = Duration(time.Millisecond)
	Second               = Duration(time.Second)
	Minute               = Duration(time.Minute)
	Hour                 = Duration(time.Hour)
	////////////////////////////////////////////////////////
)

//std
func (d Duration) Std() time.Duration {
	return time.Duration(d)
}

//expand method
func (d Duration) Days() float64 {
	return float64(d) / float64(Day)
}

//expand method
func (d Duration) Weeks() float64 {
	return float64(d) / float64(Week)
}

//expand std.Duration.String
func (d Duration) String() string {
	if d < Day {
		return d.Std().String()
	} else {
		days := int(d / Day)
		left := d % Day
		return fmt.Sprintf("%dd%s", days, left.Std().String())
	}
	return ""
}

////////////////////////////////////////////////////////
//re-export std.Duration.methods
//THIS MAKE MY EXTEND CODE UGLY

//re-export std.time.ParseDuration
func ParseDuration(s string) (Duration, error) {
	d, err := time.ParseDuration(s)
	return Duration(d), err
}

//re-export std.time.Since
func Since(t time.Time) Duration {
	return Duration(time.Since(t))
}

//re-export std.Duration.Seconds
func (d Duration) Seconds() float64 {
	return d.Std().Seconds()
}

//re-export std.Duration.Minutes
func (d Duration) Minutes() float64 {
	return d.Std().Minutes()
}

//re-export std.Duration.Hours
func (d Duration) Hours() float64 {
	return d.Std().Hours()
}

////////////////////////////////////////////////////////
