package timer

import (
	"time"
)

type TimerCallBackFunc func(uint64) uint64

//type TimerUS struct {
//}

//type TimerMS struct {
//}

//type TimerSEC struct {
//}

//type AutoTimerUS struct {
//}

//type AutoTimerMS struct {
//}

//type AutoTimerSEC struct {
//}

//type RealTime struct {
//}

func NewTimer() *Timer {
	p := &Timer{}
	p.init()
	return p
}

type Timer struct {
	running bool
	count   uint64

	ch chan uint64
}

func (me *Timer) init() {
	me.running = false
	me.count = 0
	me.ch = make(chan uint64)
}

func (me *Timer) sleep(stop bool) {
	for {
		time.Sleep(time.Millisecond) //time.Second
		if me.running {
			me.count++
			me.ch <- me.count
		} else {
			break
		}
	}
}

//func (me *TimerBase) wake() {
//	for /*me.running*/ {
//		r := <-me.ch
//		fmt.Println(r)
//	}
//}

func (me *Timer) Stop() {
	me.running = false
}

func (me *Timer) Run(_call_back_func TimerCallBackFunc) {
	me.init()
	me.running = true
	go me.sleep(false)
	//go me.wake()
	for me.running {
		r := <-me.ch
		_call_back_func(r)
	}
	//fmt.Println("Timer stoped\n")
}
