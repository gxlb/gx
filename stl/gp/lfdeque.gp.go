package gp

//#GOGP_FILE_BEGIN
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

import (
	"sync/atomic"
	"time"
)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gogp/lib/fakedef)

//#GOGP_ONCE
const (
	defaultLFDequeCap = 8
) //
//#GOGP_END_ONCE

//deque object
type GOGPGlobalNamePrefixLFDeque struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	//head point to the first elem  available for read
	//tail point to the first space available for write
	//headtail=tai<<32 + head
	headtail uint64
	busy     uint64
	d        []GOGPValueType
}

//new object
func NewGOGPDequeNamePrefixLFDeque(bufSize int32) *GOGPGlobalNamePrefixLFDeque {
	r := &GOGPGlobalNamePrefixLFDeque{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPGlobalNamePrefixLFDeque) Init(bufSize int32) {
	if nil == this.d {
		if bufSize <= 0 {
			bufSize = defaultLFDequeCap //default buffer size
		}
		this.newBuf(bufSize)
	}
	this.Clear()
	return
}

//create new buffer
func (this *GOGPGlobalNamePrefixLFDeque) newBuf(bufSize int32) {
	if bufSize > 0 {
		size := int(bufSize)
		this.d = make([]GOGPValueType, size, size) //the same cap and len
	}
}

//clear all deque data
func (this *GOGPGlobalNamePrefixLFDeque) Clear() {
	atomic.StoreUint64(&this.headtail, 0)
	this.busy = 0
}

//push to front of deque, maybe block when busy
func (this *GOGPGlobalNamePrefixLFDeque) PushFront(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}

		for {
			head, tail, headtail := this.headTail()
			head = this.prev(head)
			if head == tail { //buffer full, wait
				atomic.AddUint64(&this.busy, 1)
				time.Sleep(time.Millisecond)
				continue
			}
			this.d[head] = v
			ht := this.makeHeadTail(head, tail)
			if atomic.CompareAndSwapUint64(&this.headtail, headtail, ht) {
				break
			}
		}
	}
	return true
}

//push to back of deque, maybe block when busy
func (this *GOGPGlobalNamePrefixLFDeque) PushBack(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}

		for {
			head, tail, headtail := this.headTail()
			tailNext := this.next(tail)
			if head == tailNext { //buffer full, wait
				atomic.AddUint64(&this.busy, 1)
				time.Sleep(time.Millisecond)
				continue
			}
			this.d[tail] = v
			ht := this.makeHeadTail(head, tailNext)
			if atomic.CompareAndSwapUint64(&this.headtail, headtail, ht) {
				break
			}
		}
	}
	return true
}

//pop front of deque
func (this *GOGPGlobalNamePrefixLFDeque) PopFront() (front GOGPValueType, ok bool) {
	for {
		head, tail, headtail := this.headTail()
		if ok = head != tail; !ok {
			break
		}
		front = this.d[head]
		head = this.next(head)
		ht := this.makeHeadTail(head, tail)
		if atomic.CompareAndSwapUint64(&this.headtail, headtail, ht) {
			break
		}
	}
	return
}

//pop back of deque
func (this *GOGPGlobalNamePrefixLFDeque) PopBack() (back GOGPValueType, ok bool) {
	for {
		head, tail, headtail := this.headTail()
		tailPrev := this.prev(tail)
		if ok = head != tail; !ok {
			break
		}
		back = this.d[tailPrev]
		ht := this.makeHeadTail(head, tailPrev)
		if atomic.CompareAndSwapUint64(&this.headtail, headtail, ht) {
			break
		}
	}
	return
}

//data buffer size
func (this *GOGPGlobalNamePrefixLFDeque) Cap() uint32 {
	return uint32(len(this.d))
}

//size of deque
func (this *GOGPGlobalNamePrefixLFDeque) Size() (size uint32) {
	head, tail, _ := this.headTail()
	if tail >= head {
		size = tail - head
	} else {
		size = this.Cap() - (head - tail)
	}
	return
}

func (this *GOGPGlobalNamePrefixLFDeque) Busy() uint64 {
	return atomic.LoadUint64(&this.busy)
}

//if deque is empty
func (this *GOGPGlobalNamePrefixLFDeque) Empty() bool {
	return this.Size() == 0
}

func (this *GOGPGlobalNamePrefixLFDeque) headTail() (head, tail uint32, headtail uint64) {
	headtail = atomic.LoadUint64(&this.headtail)
	head = uint32(headtail & 0xFFFFFFFF)
	tail = uint32((headtail >> 32) & 0xFFFFFFFF)
	return
}

func (this *GOGPGlobalNamePrefixLFDeque) makeHeadTail(head, tail uint32) uint64 {
	headtail := (uint64(head) & 0xFFFFFFFF) | (uint64(tail) & 0xFFFFFFFF << 32)
	return headtail
}

//next buff
func (this *GOGPGlobalNamePrefixLFDeque) next(idx uint32) (r uint32) {
	if r = idx + 1; r >= this.Cap() {
		r = 0
	}
	return
}

//prev buff
func (this *GOGPGlobalNamePrefixLFDeque) prev(idx uint32) (r uint32) {
	if r = idx - 1; r < 0 {
		r = this.Cap() - 1
	}
	return
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
