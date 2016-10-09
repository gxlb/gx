package filelock

import (
	"sync"
	"syscall"
	"time"
)

type Mutex struct {
	w    sync.Mutex
	h    syscall.Handle
	path string
}

func (this *Mutex) Init(path string) (ok bool) {
	this.w.Lock()
	defer this.w.Unlock()
	if this.path == "" {
		this.path = path
		ok = true
	}
	return
}

func (this *Mutex) Lock() (err error) {
	this.w.Lock()
	defer this.w.Unlock()
	for {
		h, e := syscall.CreateFile(syscall.StringToUTF16Ptr(this.path),
			syscall.GENERIC_READ|syscall.GENERIC_WRITE,
			0,
			nil,
			syscall.OPEN_EXISTING,
			syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
			0)
		if e == nil {
			this.h = h
			return
		} else if e != syscall.EACCES {
			return e
		}
		time.Sleep(time.Minute) //check every minute
	}

}

func (this *Mutex) Unlock() {
	this.w.Lock()
	defer this.w.Unlock()
	syscall.CloseHandle(this.h)
	this.h = 0

}
