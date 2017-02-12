package time

import (
	"time"
)

//timing system
var sys timeSys

//time with adjust from network
type Time struct {
	local     time.Time     //time from local runtime
	netAdjust time.Duration //adjust Duration from network
}

func (this *Time) Local() time.Time {
	return this.local
}

func (this *Time) NetAdjust() time.Duration {
	return this.netAdjust
}

//time system
type timeSys struct {
	now   Time
	frame uint64
}

func (this *timeSys) init() {

}

func (this *timeSys) autoUpdate() {

}

func (this *timeSys) Start() {

}

func (this *timeSys) Stop() {

}

func (this *timeSys) Frame() uint64 {
	return this.frame
}

func (this *timeSys) Update() {

}

func (this *timeSys) UpdateAdjust(_unix_time int64) error {
	return nil
}

//auto update local time
func (this *timeSys) autoUpdate() {

}

//autoUpdate is autoUpdateTime duration
func StartSys(autoUpdate Duration) {
}

func StopSys() {}

func Frame() uint64 {
	return 0
}

func Update() (err error) {
	return nil
}

func UpdateAdjust(_unix_time int64) error {
	return nil
}

func RuntimeNow() time.Time {
	return time.Now()
}
