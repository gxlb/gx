package stopwatch

import (
	"fmt"
	"strings"
)

// NewWatchDuration  create a WatchDuration
func NewWatchDuration(name string, dur Duration, batchNum int) *WatchDuration {
	d := &WatchDuration{}
	d.Init(name, dur, batchNum)
	return d
}

// WatchDuration represents duration of step watch.
type WatchDuration struct {
	name     string
	dur      Duration
	batchNum int
}

// Init init the watch duration
func (d *WatchDuration) Init(name string, dur Duration, batchNum int) {
	d.name = name
	d.dur = dur
	d.batchNum = batchNum
}

// String show the watch duration as string.
func (d *WatchDuration) String() string {
	return fmt.Sprintf("%s\t%s",
		d.name, d.dur)
}

// Report show that watch string as string.
func (d *WatchDuration) Report(idx int, lastDur Duration) string {
	var b strings.Builder

	if idx > 0 {
		b.WriteString(fmt.Sprintf("i=%d\t", idx))
	}
	stepDur := d.dur - lastDur
	b.WriteString(fmt.Sprintf("n=%s\tt=%s\ts=%s", d.name, d.dur, stepDur))
	if d.batchNum > 0 {
		atomicDur := stepDur / Duration(d.batchNum)
		b.WriteString(fmt.Sprintf("\tu(%d)=%s", d.batchNum, atomicDur))
	}

	return b.String()
}

// Duration return duration of this step watch.
func (d *WatchDuration) Duration() Duration {
	return d.dur
}

// Name return name of this step watch.
func (d *WatchDuration) Name() string {
	return d.name
}
