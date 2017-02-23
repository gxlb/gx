package gp

//#GOGP_FILE_BEGIN 1

import (
	"sync/atomic"
)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)

//#GOGP_ONCE
const (
	defaultRingBufferCap = 8
) //
//#GOGP_END_ONCE

//refer http://lmax-exchange.github.io/disruptor/
//RingBuffer object
//Thread-safe, Use atomic algorithm to avoid Mutex lock
type GOGPGlobalNamePrefixRingBuffer struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	//head point to the first elem  available for read
	//tail point to the first space available for write
	//headtail=tai<<32 + head
	head, head_r uint64 //head_r is reverve for head
	tail, tail_r uint64
	busy         uint64
	mask         uint64
	d            []GOGPValueType
}

//init
func (this *GOGPGlobalNamePrefixRingBuffer) Init(bufSize int) {
	if nil == this.d {
		if bufSize <= defaultRingBufferCap {
			bufSize = defaultRingBufferCap //default buffer size
		}
		sizeRndUp := 1
		for sizeRndUp := 1; sizeRndUp < bufSize; sizeRndUp *= 2 { //body do nothing
		}
		this.mask = uint64(sizeRndUp - 1)
		this.newBuf(sizeRndUp)
	}
	//atomic.StoreUint64(&this.head_r, 1)
	//atomic.StoreUint64(&this.tail_r, 1)
	//this.Clear()
	return
}

//create new buffer
func (this *GOGPGlobalNamePrefixRingBuffer) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPValueType, bufSize, bufSize) //the same cap and len
	}
}

//push to back of deque, will block when full
func (this *GOGPGlobalNamePrefixRingBuffer) MustPush(v GOGPValueType) (ok bool) {
	tail_r := atomic.AddUint64(&this.tail_r, 1)
	old := tail_r - 1
	for tail_r > atomic.LoadUint64(&this.head)+this.mask+1 { //wait while full, body do nothing
	}

	idx := old & this.mask
	this.d[idx], ok = v, true
	for !atomic.CompareAndSwapUint64(&this.tail, old, tail_r) { //body do nothing
	}
	return
}

//pop front of deque, will block when empty
func (this *GOGPGlobalNamePrefixRingBuffer) MustPop() (val GOGPValueType, ok bool) {
	head_r := atomic.AddUint64(&this.head_r, 1)
	old := head_r - 1
	for old >= atomic.LoadUint64(&this.tail) { //wait while empty, body do nothing
	}

	idx := old & this.mask
	val, ok = this.d[idx], true
	for !atomic.CompareAndSwapUint64(&this.head, old, head_r) { //body do nothing
	}
	return
}

//push to back of deque, will return false when full
func (this *GOGPGlobalNamePrefixRingBuffer) Push(v GOGPValueType) (ok bool) {
	tail_r := atomic.AddUint64(&this.tail_r, 1)
	idx := tail_r & this.mask
	this.d[idx] = v
	for old := tail_r - 1; !atomic.CompareAndSwapUint64(&this.tail, old, tail_r); { //body do nothing
	}
	return
}

//pop front of deque, will return false when empty
func (this *GOGPGlobalNamePrefixRingBuffer) Pop() (val GOGPValueType, ok bool) {
	head_r := atomic.AddUint64(&this.head_r, 1)
	idx := head_r & this.mask
	val, ok = this.d[idx], true
	for old := head_r - 1; !atomic.CompareAndSwapUint64(&this.head, old, head_r); { //body do nothing
	}
	return
}

//data buffer size
func (this *GOGPGlobalNamePrefixRingBuffer) Cap() uint32 {
	return uint32(len(this.d))
}

//size of queue
func (this *GOGPGlobalNamePrefixRingBuffer) Size() (size int) {
	return int(atomic.LoadUint64(&this.tail) - atomic.LoadUint64(&this.head))
}

//Busy returns how many times Push when busy
func (this *GOGPGlobalNamePrefixRingBuffer) Busy() uint64 {
	return atomic.LoadUint64(&this.busy)
}

//if queue is empty
func (this *GOGPGlobalNamePrefixRingBuffer) Empty() bool {
	return this.Size() == 0
}

//#GOGP_FILE_END
