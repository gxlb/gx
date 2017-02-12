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

//push to front of deque, maybe block when busy
//func (this *GOGPGlobalNamePrefixRingBuffer) PushFront(v GOGPValueType) (ok bool) { return }

//push to back of deque, maybe block when busy
func (this *GOGPGlobalNamePrefixRingBuffer) Push(v GOGPValueType) (ok bool) {
	tail_r := atomic.AddUint64(&this.tail_r, 1)
	idx := tail_r & this.mask
	this.d[idx] = v
	for old := tail_r - 1; !atomic.CompareAndSwapUint64(&this.tail, old, tail_r); { //body do nothing
	}
	return
}

//pop front of deque
func (this *GOGPGlobalNamePrefixRingBuffer) Pop() (val GOGPValueType, ok bool) {
	head_r := atomic.AddUint64(&this.head_r, 1)
	idx := head_r & this.mask
	val, ok = this.d[idx], true
	for old := head_r - 1; !atomic.CompareAndSwapUint64(&this.head, old, head_r); { //body do nothing
	}
	return
}

//pop back of deque
//func (this *GOGPGlobalNamePrefixRingBuffer) PopBack() (back GOGPValueType, ok bool) { return }

//data buffer size
func (this *GOGPGlobalNamePrefixRingBuffer) Cap() uint32 {
	return uint32(len(this.d))
}

////size of deque
//func (this *GOGPGlobalNamePrefixRingBuffer) Size() (size uint32) {
//	head, tail, _ := this.headTail()
//	if tail >= head {
//		size = tail - head
//	} else {
//		size = this.Cap() - (head - tail)
//	}
//	return
//}

//Busy returns how many times Push when busy
func (this *GOGPGlobalNamePrefixRingBuffer) Busy() uint64 {
	return atomic.LoadUint64(&this.busy)
}

////if deque is empty
//func (this *GOGPGlobalNamePrefixRingBuffer) Empty() bool {
//	return this.Size() == 0
//}

//#GOGP_FILE_END
