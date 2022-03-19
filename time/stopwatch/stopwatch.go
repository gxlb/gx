//Package stopwatch implements some easy way to mesure time
package stopwatch

import (
	"bytes"
	"fmt"
	"time"
)

// Duration alias time.Duration
type Duration = time.Duration

// New create a new StopWatch.
func New(options ...Option) *StopWatch {
	p := &StopWatch{}

	opts := loadOptions(options...)
	if opts.StepCacheCap > 0 {
		p.durList = make([]WatchDuration, 0, opts.StepCacheCap)
	}
	p.Start()

	return p
}

//StopWatch is a useful time mesure tool object
type StopWatch struct {
	startTime time.Time
	lastDur   Duration //last watch duration
	durList   []WatchDuration
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

// Restart restart the stop watch and return time duration from last start time
func (w *StopWatch) Restart() Duration {
	t := time.Now()
	d := t.Sub(w.startTime)
	w.durList = w.durList[:0]
	w.startTime = t
	w.lastDur = 0
	return d
}

// Clean reset the watch list but not reset the start time.
func (w *StopWatch) Clean() Duration {
	t := time.Now()
	d := t.Sub(w.startTime)
	w.durList = w.durList[:0]
	w.lastDur = 0
	return d
}

// AddStepWatch add a new step watch with given name.
func (w *StopWatch) AddStepWatch(name string, times uint) WatchDuration {
	t := time.Now()
	var d WatchDuration
	if times <= 0 { //0 is forbinden
		times = 1
	}
	d.Init(name, t.Sub(w.startTime), times)

	w.lastDur = d.dur
	w.durList = append(w.durList, d)

	return d
}

// ReportOnce report duration as string.
func (w *StopWatch) ReportOnce(name string) string {
	dur := time.Now().Sub(w.startTime)
	stepDur := dur - w.lastDur

	w.lastDur = dur

	return fmt.Sprintf("%s:%s/%s", name, stepDur, dur)
}

// ReportWatch report step duration as string.
func (w *StopWatch) ReportWatch(name string, times uint) string {
	dur := time.Now().Sub(w.startTime)
	stepDur := dur - w.lastDur
	if times <= 0 { //0 is forbinden
		times = 1
	}

	w.lastDur = dur

	atomicTime := stepDur / Duration(times)

	return fmt.Sprintf("Watch%d\t%s\t%s\t%s\t%s\t%d",
		len(w.durList), name, dur, stepDur, atomicTime, int64(atomicTime))
}

// ReportAll report all step watch as string.
func (w *StopWatch) ReportAll(name string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("StopWatch(%s):\n", name))
	if nil != w.durList {
		lastDur := Duration(0)
		for i, v := range w.durList {
			buf.WriteString(v.Report(i+1, lastDur))
			buf.WriteString("\n")
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
	dur := time.Now().Sub(w.startTime)
	return dur
}

// TellStepDuration return duration from last step.
func (w *StopWatch) TellStepDuration() Duration {
	dur := time.Now().Sub(w.startTime)
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
