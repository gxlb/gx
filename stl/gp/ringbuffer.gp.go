package gp

//#GOGP_FILE_BEGIN 1

//import (
//	"sync/atomic"
//)

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
	head, head_r uint64
	tail, tail_r uint64
	busy         uint64
	d            []GOGPValueType
}

//init
func (this *GOGPGlobalNamePrefixRingBuffer) Init(bufSize int) {
	if nil == this.d {
		if bufSize <= defaultRingBufferCap {
			bufSize = defaultRingBufferCap //default buffer size
		}
		for i := 1; i < bufSize; i *= 2 {
		}
		//this.newBuf(bufSize)
	}
	//this.Clear()
	return
}

//create new buffer
func (this *GOGPGlobalNamePrefixRingBuffer) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPValueType, bufSize, bufSize) //the same cap and len
	}
}

//#GOGP_FILE_END
