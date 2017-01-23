package gp

//#GOGP_FILE_BEGIN 1
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
///*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

import (
	"sync/atomic"
	"unsafe"
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

//deque object
type GOGPGlobalNamePrefixLFDeque struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	//head point to the first elem  available for read
	//tail point to the first space available for write
	//headtail=tai<<32 + head
	headtail uint64
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
			bufSize = 8 //default buffer size
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
	atomic.StoreInt64((*int64)(unsafe.Pointer(&this.head)), 0)
}

//push to front of deque
func (this *GOGPGlobalNamePrefixLFDeque) PushFront(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}

		this.head = this.prev(this.head) //move head to prev empty space
		this.d[this.head] = v
		if this.head == this.tail { //head reaches tail, buffer full
			oldCap := this.Cap()
			d := this.d
			this.newBuf(oldCap * 2)
			h := copy(this.d, d[this.head:])
			t := copy(this.d[:h], d[:this.tail])
			this.head, this.tail = 0, int32(h+t)
		}
	}
	return
}

//push to back of deque
func (this *GOGPGlobalNamePrefixLFDeque) PushBack(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}

		this.d[this.tail] = v
		this.tail = this.next(this.tail)
		if this.tail == this.head { //tail catches up head, buffer full
			oldCap := this.Cap()
			d := this.d
			this.newBuf(oldCap * 2)
			h := copy(this.d, d[this.head:])
			t := copy(this.d[:h], d[:this.tail])
			this.head, this.tail = 0, int32(h+t)
		}
	}
	return
}

//pop front of deque
func (this *GOGPGlobalNamePrefixLFDeque) PopFront() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
		this.head = this.next(this.head)
	}
	return
}

//pop back of deque
func (this *GOGPGlobalNamePrefixLFDeque) PopBack() (back GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		this.tail = this.prev(this.tail)
		back = this.d[this.tail]
	}
	return
}

//data buffer size
func (this *GOGPGlobalNamePrefixLFDeque) Cap() int32 {
	return int32(len(this.d))
}

//size of deque
func (this *GOGPGlobalNamePrefixLFDeque) Size() (size int32) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = this.Cap() - (this.head - this.tail)
	}
	return
}

//if deque is empty
func (this *GOGPGlobalNamePrefixLFDeque) Empty() bool {
	return this.Size() == 0
}

//next buff
func (this *GOGPGlobalNamePrefixLFDeque) next(idx int32) (r int32) {
	if r = idx + 1; r >= this.Cap() {
		r = 0
	}
	return
}

//prev buff
func (this *GOGPGlobalNamePrefixLFDeque) prev(idx int32) (r int32) {
	if r = idx - 1; r < 0 {
		r = this.Cap() - 1
	}
	return
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
