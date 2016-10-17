package gp

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
///*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is exactly not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//import here...

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

import (
	dumy_fmt "fmt"
)

type GOGPDequeElem int

func (me GOGPDequeElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//deque object
type GOGPDequeNamePrefixDeque struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1) = 0
	//if tail<head, data is [head, end, 0, tail)
	head int
	tail int
	d    []GOGPDequeElem
}

//new object
func NewGOGPDequeNamePrefixDeque(bufSize int) *GOGPDequeNamePrefixDeque {
	r := &GOGPDequeNamePrefixDeque{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPDequeNamePrefixDeque) Init(bufSize int) {
	if nil == this.d {
		if bufSize <= 0 {
			bufSize = 8 //default buffer size
		}
		this.newBuf(bufSize)
	}
	this.Clear()
	return
}

func (this *GOGPDequeNamePrefixDeque) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPDequeElem, bufSize, bufSize) //the same cap and len
	}
}

//clear
func (this *GOGPDequeNamePrefixDeque) Clear() {
	this.head, this.tail = 0, 0
}

//push to front of deque
func (this *GOGPDequeNamePrefixDeque) PushFront(v GOGPDequeElem) (ok bool) {
	//	if ok = true; ok {
	//		if nil == this.d { //init if needed
	//			this.Init(-1)
	//		}
	//		t := this.head - 1
	//		if t < 0 {
	//			t = this.Cap() - 1
	//		}
	//		this.d[t] = v
	//		if this.tail >= this.head { //normal order [head,tail)
	//			if t == this.tail { //buffer full
	//			} else {
	//			}
	//		} else { //this.tail < this.head, abnormal order [head, end, 0, tail)
	//		}
	//	}
	return
}

//push to back of deque
func (this *GOGPDequeNamePrefixDeque) PushBack(v GOGPDequeElem) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}
		if this.tail >= this.head { //normal order [head,tail)
			if this.tail < this.Cap() {
				this.d[this.tail] = v
				this.tail++
				if this.tail >= this.Cap() {
					if this.head >= 1 {
						this.tail = 0
					} else { //increase buffer
						t := this.d
						this.newBuf(len(t) * 2)
						copy(this.d, t)
					}
				}
			}
		} else { //this.tail < this.head, abnormal order [head, end, 0, tail)
			this.d[this.tail] = v
			if this.tail++; this.tail >= this.head { //buffer full,copy data as normal order, and of cause, the buffer will change
				oldSize := this.Cap()
				t := this.d
				this.newBuf(oldSize * 2)
				copy(this.d, t[:this.head])
				copy(this.d[:this.head], t[:this.tail])
				this.head, this.tail = 0, oldSize
			}
		}
	}
	return
}

//pop front of deque
func (this *GOGPDequeNamePrefixDeque) PopFront() (front GOGPDequeElem, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
		if this.head++; this.head >= this.Cap() && this.head != this.tail {
			this.head = 0
		}
	}
	return
}

//pop back of deque
func (this *GOGPDequeNamePrefixDeque) PopBack() (back GOGPDequeElem, ok bool) {
	if ok = this.head != this.tail; ok {
		t := this.tail - 1
		if t < 0 {
			t = this.Cap() - 1
		}
		back = this.d[t]
		this.tail = t
	}
	return
}

//front data
func (this *GOGPDequeNamePrefixDeque) Front() (front GOGPDequeElem, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
	}
	return
}

//back data
func (this *GOGPDequeNamePrefixDeque) Back() (back GOGPDequeElem, ok bool) {
	if ok = this.head != this.tail; ok {
		t := this.tail - 1
		if t < 0 {
			t = this.Cap() - 1
		}
		back = this.d[t]
	}
	return
}

//cap
func (this *GOGPDequeNamePrefixDeque) Cap() int {
	return len(this.d)
}

//size of deque
func (this *GOGPDequeNamePrefixDeque) Size() (size int) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = this.Cap() - this.head + this.tail
	}
	return
}

//if deque is empty
func (this *GOGPDequeNamePrefixDeque) Empty() bool {
	return this.Size() == 0
}

////show
//func (this *GOGPDequeNamePrefixDeque) Show() string {
//	var b show_bytes.Buffer
//	b.WriteByte('[')
//	for i := this.head; i != this.tail; i++ {
//		if i >= this.Cap() {
//			i = 0
//		}
//		v = this.v[i]
//		b.WriteString(v.Show())
//		b.WriteByte(',')
//	}
//	if this.Depth() > 0 {
//		b.Truncate(b.Len() - 1) //remove last ','
//	}
//	b.WriteByte(']')
//	return b.String()
//}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
