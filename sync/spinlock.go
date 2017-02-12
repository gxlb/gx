package sync

import (
	"sync"
	"sync/atomic"
)

type Locker sync.Locker

const (
	spinLockLocked = 1 << iota

	spinLockMaxLoop = 1 << 24 //max spin times, to avoid dead lock
)

type SpinLock struct {
	state int32
}

func (this *SpinLock) Lock() {
	for i := 0; ; i++ {
		if i > spinLockMaxLoop {
			panic("[sync] maybe dead lock on SpinLock")
		}

		if atomic.CompareAndSwapInt32(&this.state, 0, spinLockLocked) {
			break
		}
	}
}
func (this *SpinLock) Unlock() {
	if !atomic.CompareAndSwapInt32(&this.state, spinLockLocked, 0) {
		panic("[sync] Unlock on unlocked SpinLock")
	}
}
