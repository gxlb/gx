//Package stopwatch implements some easy way to mesure time
package stopwatch

import (
	"bytes"
	"fmt"
	"time"
)

// Duration alias time.Duration
type Duration = time.Duration

// Duration exports
const (
	Day  = 24 * Hour //expand const
	Week = 7 * Day   //expand const

	Nanosecond  Duration = Duration(time.Nanosecond)
	Microsecond          = Duration(time.Microsecond)
	Millisecond          = Duration(time.Millisecond)
	Second               = Duration(time.Second)
	Minute               = Duration(time.Minute)
	Hour                 = Duration(time.Hour)
)

// New create a new StopWatch.
func New(options ...Option) *StopWatch {
	opts := loadOptions(options...)
	p := &StopWatch{
		name:      opts.Name,
		startTime: time.Now(),
		lastDur:   0,
		unit:      opts.Unit,
	}
	if opts.StepCacheCap > 0 {
		p.durList = make([]WatchDuration, 0, opts.StepCacheCap)
	}

	return p
}

//StopWatch is a useful time mesure tool object
type StopWatch struct {
	name      string
	startTime time.Time
	lastDur   Duration //last watch duration
	durList   []WatchDuration
	unit      Duration
}

// Start reset the start time.
func (w *StopWatch) Start() time.Time {
	w.startTime = time.Now()
	w.durList = w.durList[:0]
	w.lastDur = 0
	return w.startTime
}

// Stop is alias of Restart
func (w *StopWatch) Stop() Duration {
	return w.Restart()
}

func (w *StopWatch) duration(reset bool) Duration {
	t := time.Now()
	dur := TrimDuration(t.Sub(w.startTime), w.unit)
	if reset {
		w.startTime = t
	}
	return dur
}

// Restart restart the stop watch and return time duration from last start time
func (w *StopWatch) Restart() Duration {
	d := w.duration(true)
	w.durList = w.durList[:0]
	w.lastDur = 0
	return d
}

// Clean reset the watch list but not reset the start time.
func (w *StopWatch) Clean() Duration {
	d := w.duration(false)
	w.durList = w.durList[:0]
	w.lastDur = 0
	return d
}

// AddStepWatch add a new step watch with given name.
// batchNum is used to report atomic time for batch testing.
func (w *StopWatch) AddStepWatch(name string, batchNum int) WatchDuration {
	var d WatchDuration
	d.Init(name, w.duration(false), batchNum)

	w.lastDur = d.dur
	w.durList = append(w.durList, d)

	return d
}

// ReportOnce report duration as string.
func (w *StopWatch) ReportOnce(name string, batchNum int) string {
	var d WatchDuration
	d.Init(name, w.duration(false), batchNum)

	return d.Report(-1, w.lastDur)
}

// ReportAll report all step watch as string.
func (w *StopWatch) ReportAll() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("StopWatch(%s):\n", w.name))
	if nil != w.durList {
		lastDur := Duration(0)
		for i, v := range w.durList {
			if i > 0 {
				buf.WriteString("\n")
			}
			buf.WriteString(v.Report(i+1, lastDur))

			lastDur = v.dur
		}
	}

	return buf.String()
}

// Count return number of steps reported.
func (w *StopWatch) Count() int {
	return len(w.durList)
}

// TellDuration return duration from start time.
func (w *StopWatch) TellDuration() Duration {
	dur := w.duration(false)
	return dur
}

// TellStepDuration return duration from last step.
func (w *StopWatch) TellStepDuration() Duration {
	dur := w.duration(false)
	stepDur := dur - w.lastDur
	return stepDur
}

// GetWatch return step of of given index.
func (w *StopWatch) GetWatch(index int) (d WatchDuration, ok bool) {
	if index >= 0 && index < len(w.durList) {
		return w.durList[index], true
	}
	return
}

// TellAllWatch return the list of step watch.
func (w *StopWatch) TellAllWatch() []WatchDuration {
	return w.durList
}
