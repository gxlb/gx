//Package stopwatch implement some easy way to mesure time
package stopwatch

import (
	"bytes"
	"fmt"
	"time"
)

type WatchType int

// const (
// 	WATCH_BY_STEP WatchType = iota //default
// 	WATCH_BY_START
// )

func New() *StopWatch {
	return NewP(1)
}

func NewP(nWatchs int) *StopWatch {
	p := &StopWatch{}
	p.Start()
	//p.SetType(whatch_type)
	if nWatchs > 0 {
		p.dur_list = make([]WatchDur, 0, nWatchs)
	}
	return p
}

type WatchDur struct {
	name  string
	dur   time.Duration
	times uint
}

func (me *WatchDur) Init(name string, dur time.Duration, times uint) {
	if times <= 0 { //0 is forbinden
		times = 1
	}
	me.name = name
	me.dur = dur
	me.times = times
}

func (me *WatchDur) String() string {
	return fmt.Sprintf("%s\t%s\t%s", me.name, me.dur, me.dur/time.Duration(me.times))
}

func (me *WatchDur) Report(idx int, last_dur time.Duration) string {
	step_dur := me.dur - last_dur
	atomic_time := last_dur / time.Duration(me.times)

	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%d", idx, me.name, me.dur, step_dur, atomic_time)
}

// func (me *WatchDur) ReportAvg(idx int, run_times uint) string {
// 	if run_times <= 0 {
// 		run_times = 1
// 	}
// 	step_dur := me.dur / time.Duration(run_times)
// 	atomic_time := step_dur / time.Duration(me.times)

// 	return fmt.Sprintf("%d\t%s\t%s\t%s\t%d", idx, me.name, step_dur, atomic_time, int64(atomic_time))
// }

func (me *WatchDur) Duration() time.Duration {
	return me.dur
}

//stop watch is a useful time mesure tool object
type StopWatch struct {
	start_time time.Time
	last_dur   time.Duration //last watch duration
	last_idx   int           //last watch index
	dur_list   []WatchDur

	//avg_dur_list []WatchDur
}

// func (me *StopWatch) SetType(t WatchType) {
// 	me.by_step = (t == WATCH_BY_STEP)
// }

func (me *StopWatch) Start() time.Time {
	me.start_time = time.Now()
	me.dur_list = nil
	me.last_dur = 0
	me.last_idx = 0
	return me.start_time
}

func (me *StopWatch) Stop() time.Duration {
	t := time.Now()
	d := t.Sub(me.start_time)
	return d
}

func (me *StopWatch) Clear() time.Duration {
	t := time.Now()
	d := t.Sub(me.start_time)
	me.dur_list = nil
	//me.avg_dur_list = nil
	me.last_dur = 0
	me.last_idx = 0
	return d
}

func (me *StopWatch) AddWatch(name string, times uint) WatchDur {
	t := time.Now()
	var d WatchDur
	if times <= 0 { //0 is forbinden
		times = 1
	}
	d.Init(name, t.Sub(me.start_time), times)

	// step_dur := d.dur - me.last_dur
	// if len(me.avg_dur_list) <= len(me.dur_list) {
	// 	dd := d
	// 	dd.dur = step_dur
	// 	me.avg_dur_list = append(me.avg_dur_list, dd)
	// } else {
	// 	me.avg_dur_list[len(me.dur_list)].dur += step_dur
	// }

	me.last_dur = d.dur
	me.last_idx++
	me.dur_list = append(me.dur_list, d)

	return d
}

func (me *StopWatch) ReportOnce(name string) string {
	dur := time.Now().Sub(me.start_time)
	step_dur := dur - me.last_dur

	me.last_dur = dur
	me.last_idx++

	return fmt.Sprintf("%s:%s", name, step_dur)
}

func (me *StopWatch) ReportWatch(name string, times uint) string {
	dur := time.Now().Sub(me.start_time)
	step_dur := dur - me.last_dur
	if times <= 0 { //0 is forbinden
		times = 1
	}

	me.last_dur = dur
	me.last_idx++

	atomic_time := step_dur / time.Duration(times)

	//me.AddWatch(name, times)
	return fmt.Sprintf("Watch%d\t%s\t%s\t%s\t%s\t%d", me.last_idx, name, dur, step_dur, atomic_time, int64(atomic_time))
}

func (me *StopWatch) Count() int {
	return len(me.dur_list)
}

func (me *StopWatch) TellWatch(idx int) (d WatchDur) {
	if me.dur_list != nil {
		if idx >= 0 && idx < len(me.dur_list) {
			d = me.dur_list[idx]
			// if me.by_step && idx > 0 {
			// 	d.dur -= me.dur_list[idx-1].dur
			// }
		}
	}
	return
}

func (me *StopWatch) TellAllWatch() []WatchDur {
	// if me.by_step {
	// 	return me.tellAllStepWatch()
	// }
	return me.dur_list
}

// func (me *StopWatch) TellAvgWatch() []WatchDur {
// 	return me.avg_dur_list
// }

// func (me *StopWatch) tellAllStepWatch() (r []WatchDur) {
// 	if nil != me.dur_list {
// 		r = append(r, me.dur_list...)
// 		for i := len(r) - 1; i > 0; i-- {
// 			r[i].dur -= r[i-1].dur
// 		}
// 	}
// 	return
// }

func (me *StopWatch) Report() string {
	var buf bytes.Buffer
	buf.WriteString("====StopWatch report:\n")
	if nil != me.dur_list {
		last_dur := time.Duration(0)
		for i, v := range me.dur_list {
			buf.WriteString(v.Report(i+1, last_dur))
			buf.WriteString("\n")
			last_dur = v.dur
		}
	}
	buf.WriteString("====StopWatch report end")

	return buf.String()
}

func (me *StopWatch) Restart() time.Duration {
	t := time.Now()
	d := t.Sub(me.start_time)
	me.dur_list = nil
	me.start_time = t
	me.last_dur = 0
	me.last_idx = 0
	return d
}
