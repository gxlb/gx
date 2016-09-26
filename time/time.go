//Package time implements a net-synced time system based on std.time and support cheap time mesure SPNow()
package time

//to speed up, must remove lock

//import (
//	"fmt"
//	"sync"
//	"time"

//	"vipally.gmail.com/basic/goroutine"
//)

//var (
//	G_DEBUG   = false
//	g_ch_stop = make(chan bool)
//)

//func init() {
//	UpdateTime(time.Now())
//}

//type Duration time.Duration

//const (
//	Nanosecond  Duration = Duration(time.Nanosecond)
//	Microsecond          = Duration(time.Microsecond)
//	Millisecond          = Duration(time.Millisecond)
//	Second               = Duration(time.Second)
//	Minute               = Duration(time.Minute)
//	Hour                 = Duration(time.Hour)
//	Day                  = 24 * Hour
//	Week                 = 7 * Day
//)

//var (
//	g_autoupdate *time.Ticker
//	g_now        Time //now
//	g_now_lock   sync.RWMutex
//	g_frame      uint64
//	g_frame_lock sync.RWMutex
//)

//func Frame() uint64 {
//	g_frame_lock.RLock()
//	defer g_frame_lock.RUnlock()
//	return g_frame
//}

////d is autoUpdateTime duration
//func Start(d Duration) {
//	UpdateTime(time.Now())

//	g_now_lock.Lock()
//	g_autoupdate = time.NewTicker(time.Duration(d))
//	g_now_lock.Unlock()

//	g_frame_lock.Lock()
//	g_frame = 0
//	g_frame_lock.Unlock()

//	goroutine.Go(autoUpdateTime, d)
//	if G_DEBUG {
//		fmt.Println("time Started", time.Now())
//	}
//}
//func Stop() {
//	g_now_lock.Lock()
//	g_ch_stop <- true //send channel first
//	g_autoupdate.Stop()
//	g_now_lock.Unlock()
//	if G_DEBUG {
//		fmt.Println("time Stoped", time.Now())
//	}
//}

//func UpdateTime(now time.Time) (err error) {
//	g_now_lock.Lock()
//	defer g_now_lock.Unlock()
//	g_now.now = now.Add(g_now.duration)
//	if err = updateSimpleTime(); err != nil {
//		return err
//	}
//	return
//}

//func UpdateDuration(_unix_time int64) error {
//	svr_time := time.Unix(_unix_time, 0)
//	g_now_lock.Lock()
//	defer g_now_lock.Unlock()
//	t := time.Now()
//	g_now.duration = svr_time.Sub(t)

//	return nil
//}

//func TNow() Time {
//	g_now_lock.RLock()
//	defer g_now_lock.RUnlock()
//	return g_now
//}

//func RuntimeNow() time.Time {
//	UpdateTime(time.Now())
//	g_now_lock.RLock()
//	defer g_now_lock.RUnlock()
//	return g_now.Now()
//}

//func Now() time.Time {
//	//g_now_lock.RLock()
//	//defer g_now_lock.RUnlock()
//	return g_now.Now() //now()
//}
//func now() time.Time {
//	return g_now.Now()
//}

//func NetDuration() time.Duration {
//	g_now_lock.RLock()
//	defer g_now_lock.RUnlock()
//	return g_now.NetDuration()
//}

//func autoUpdateTime(d interface{}, para ...interface{}) {
//	g_now_lock.Lock()
//	var c <-chan time.Time = g_autoupdate.C
//	g_now_lock.Unlock()
//	var now time.Time
//	for {
//		select {
//		case now = <-c:
//		case <-g_ch_stop:
//			//log:time system shutdown now
//			break
//		}

//		UpdateTime(now)

//		g_frame_lock.Lock()
//		g_frame++
//		g_frame_lock.Unlock()

//		//todo : new frame event
//		if G_DEBUG {
//			fmt.Printf("%d UpdateTime:%v\n", Frame(), now)
//		}
//	}
//}

//type Time struct {
//	now      time.Time
//	duration time.Duration
//}

//func (me *Time) Now() time.Time {
//	return me.now
//}

//func (me *Time) NetDuration() time.Duration {
//	return me.duration
//}
