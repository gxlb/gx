package time_test

import (
	"fmt"
	"testing"

	xtime "github.com/vipally/gx/time"
)

func TestDuration(t *testing.T) {
	var d xtime.Duration = xtime.Day + xtime.Hour
	fmt.Println("Seconds", d, d.T().Seconds())
	fmt.Println("Days", d, d.Days())
	fmt.Println("Weeks", d, d.Weeks())
}
