package time

import (
	"fmt"
	"sync"
	"time"
)

type SimpleTime struct {
	Year        uint16 "Year"
	Millisecond uint16 "Millisecond"
	Month       uint8  "Month"
	Day         uint8
	Hour        uint8
	Minute      uint8
	Senond      uint8
	Weekday     uint8
}

var (
	g_sp_time SimpleTime
	//g_uinx              uint32
	g_str_date          = "uninitialized simple time"
	g_str_time          = g_str_date
	g_str_datetime      = g_str_date
	g_str_datetime_full = g_str_date
	g_str_full          = g_str_date
	g_lock              sync.RWMutex
)

func updateSimpleTime() error {
	g_lock.Lock()
	defer g_lock.Unlock()
	t := now()
	y, m, d := t.Date()
	g_sp_time.Year = uint16(y)
	g_sp_time.Month = uint8(m)
	g_sp_time.Day = uint8(d)
	g_sp_time.Hour = uint8(t.Hour())
	g_sp_time.Minute = uint8(t.Minute())
	g_sp_time.Senond = uint8(t.Second())
	g_sp_time.Millisecond = uint16(t.Nanosecond() / int(time.Millisecond))
	g_sp_time.Weekday = uint8(t.Weekday())
	g_str_date = fmt.Sprintf("%04d-%02d-%02d",
		g_sp_time.Year, g_sp_time.Month, g_sp_time.Day)
	g_str_time = fmt.Sprintf("%02d:%02d:%02d",
		g_sp_time.Hour, g_sp_time.Minute, g_sp_time.Senond)
	g_str_datetime = fmt.Sprintf("%s %s",
		g_str_date, g_str_time)
	g_str_datetime_full = fmt.Sprintf("%s %s %s",
		g_str_date, t.Weekday().String(), g_str_time)

	g_str_full = fmt.Sprintf("%s %s %s.%03d ",
		g_str_date, t.Weekday().String(), g_str_time, g_sp_time.Millisecond)
	return nil
}

func lock_get_sp_time() SimpleTime {
	g_lock.RLock()
	defer g_lock.RUnlock()
	return g_sp_time
}

//simple time now,very cheap and useful time format support
func SPNow() (r SimpleTime) {
	r = lock_get_sp_time()
	return
}

func (cp SimpleTime) Date() (r uint32) {
	sp := lock_get_sp_time()
	return uint32(sp.Year)*10000 + uint32(sp.Month)*100 + uint32(sp.Day)
}
func (cp SimpleTime) Time() (r uint32) {
	sp := lock_get_sp_time()
	return uint32(sp.Hour)*10000 + uint32(sp.Minute)*100 + uint32(sp.Senond)
}
func (cp SimpleTime) DateTime() (r uint64) {
	return uint64(cp.Date())*1000000 + uint64(cp.Time())
}
func (cp SimpleTime) FullTime() (r uint64) {
	return cp.DateTime()*1000 + uint64(cp.Millisecond)
}

func (cp SimpleTime) String() string {
	return g_str_full
}
func (cp SimpleTime) StrDate() string {
	return g_str_date
}
func (cp SimpleTime) StrTime() string {
	return g_str_time
}
func (cp SimpleTime) StrDateTime() string {
	return g_str_datetime
}
func (cp SimpleTime) StrDateTimeFull() string {
	return g_str_datetime_full
}
