package stopwatch

import (
	"time"
)

// NewSingle create a Single stopwatch
func NewSingle(options ...Option) *Single {
	opts := loadOptions(options...)
	p := &Single{
		startTime: time.Now(),
		unit:      opts.Unit,
	}
	return p
}

// Single is a simple stopwatch that can only measure once.
type Single struct {
	unit      Duration
	startTime time.Time
}

// StartTime return start time of this stopwatch.
func (w *Single) StartTime() time.Time {
	return w.startTime
}

// Start reset start time of this stopwatch.
func (w *Single) Start() time.Time {
	w.startTime = time.Now()
	return w.StartTime()
}

func (w *Single) duration(reset bool) time.Duration {
	t := time.Now()
	dur := TrimDuration(t.Sub(w.startTime), w.unit)
	if reset {
		w.startTime = t
	}
	return dur
}

// Restart reset the start time and return duration since last start time.
func (w *Single) Restart() Duration {
	return w.duration(true)
}

// Duration return duration since start time.
func (w *Single) Duration() Duration {
	return w.duration(false)
}

// TrimDuration scale duration to boundary of time unit.
func TrimDuration(dur Duration, unit Duration) Duration {
	if unit > 0 {
		dur += unit / 2
		return dur - dur%unit
	}
	return dur
}
