package time_test

import (
	"fmt"
	"testing"

	"time" //github.com/vipally/gx/time"

	xtime "github.com/vipally/gx/time"
)

func TestStdDuration(t *testing.T) {
	var d time.Duration = time.Hour*25 + time.Minute
	fmt.Printf("DurationValue:%d, String:%s, Hours:%f\n",
		int64(d), d, d.Hours())

	//import("github.com/vipally/gx/time"):
	//	DurationValue:90060000000000, String:1d1h1m0s, Hours:25.016667

	//import("time"):
	//	DurationValue:90060000000000, String:25h1m0s, Hours:25.016667
}

func TestDuration(t *testing.T) {
	var ds = []xtime.Duration{xtime.Day + xtime.Hour*12, xtime.Hour*12 + xtime.Minute}
	for _, d := range ds {
		fmt.Printf("DurationValue:%d, std:%s, expand:%s, Days:%f, Hours:%f\n",
			int64(d), d.Std(), d, d.Days(), d.Hours())
	}

	//DurationValue:129600000000000, std:36h0m0s, expand:1d12h0m0s, Days:1.500000, Hours:36.000000
	//DurationValue:43260000000000, std:12h1m0s, expand:12h1m0s, Days:0.500694, Hours:12.016667
}
