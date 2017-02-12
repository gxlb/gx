package sync

import (
	"sync/atomic"
)

const (
	rwSpinLockMaxReaders = 1 << 30
)

type RWSpinLock struct {
	w           SpinLock
	readerCount int32
	//readerWait  int32
}

func (this *RWSpinLock) Lock() {
	this.w.Lock()                                             //get writer lock first
	atomic.AddInt32(&this.readerCount, -rwSpinLockMaxReaders) //tell readers there is a writer
}

func (this *RWSpinLock) Unlock() {
	if r := atomic.AddInt32(&this.readerCount, rwSpinLockMaxReaders); r >= rwSpinLockMaxReaders {
		panic("[sync] Unlock on unlocked RWSpinLock")
	}
	this.w.Unlock()
}

func (this *RWSpinLock) RLock() {
	for i := 0; ; i++ {
		if i > spinLockMaxLoop {
			panic("[sync] maybe dead lock on RWSpinLock.RLock")
		}

		if r := atomic.LoadInt32(&this.readerCount); r >= 0 { // A writer is pending, wait for it.
			if atomic.AddInt32(&this.readerCount, 1) == 1 {
				this.w.Lock() //first waiter reader, block writer
			}
			break
		}
	}
}
func (this *RWSpinLock) RUnlock() {
	if r := atomic.AddInt32(&this.readerCount, -1); r < 0 {
		if r+1 == 0 || r+1 == -rwSpinLockMaxReaders {
			panic("sync: RUnlock of unlocked RWSpinLock")
		}
		if r == 0 {
			// The last reader unblocks the writer.
			this.w.Unlock()
		}
	}
}
