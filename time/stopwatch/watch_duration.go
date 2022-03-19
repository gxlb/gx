package stopwatch

import (
	"fmt"
)

// WatchDuration represents duration of step watch.
type WatchDuration struct {
	name  string
	dur   Duration
	times uint
}

// Init init the watch duration
func (d *WatchDuration) Init(name string, dur Duration, times uint) {
	if times <= 0 { //0 is forbinden
		times = 1
	}
	d.name = name
	d.dur = dur
	d.times = times
}

// String show the watch duration as string.
func (d *WatchDuration) String() string {
	return fmt.Sprintf("%s\t%s\t%s",
		d.name, d.dur, d.dur/Duration(d.times))
}

// Report show that watch string as string.
func (d *WatchDuration) Report(idx int, lastDur Duration) string {
	stepDur := d.dur - lastDur
	atomicTime := lastDur / Duration(d.times)

	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s",
		idx, d.name, d.dur, stepDur, atomicTime)
}

// Duration return duration of this step watch.
func (d *WatchDuration) Duration() Duration {
	return d.dur
}

// Name return name of this step watch.
func (d *WatchDuration) Name() string {
	return d.name
}
