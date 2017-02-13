package tasks

import (
	"github.com/vipally/gx/time"
)

var taskSys struct {
	run bool
}

type TaskSysStatus struct {
	Runing          bool
	WaitingFinish   bool
	CurWorkers      int
	MaxWorkers      int
	MaxListLen      int
	TooBusyCount    uint64
	FinishTaskCount uint64
	FreeTime        time.Duration
	BusyTime        time.Duration

	//caculate values
	AvgWorkers float32
	AvgListLen float32
	TotalTime  time.Duration
}

func Start() {

}

func Stop() {

}

func PushTask() {

}

func workerFunc() {}
