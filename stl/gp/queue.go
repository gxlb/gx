package gp

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
/*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
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

type GOGPQueueElem int

func (me GOGPQueueElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//stack object
type GOGPQueueNamePrefixQueue struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1) = 0
	head int
	tail int
	d    []GOGPQueueElem
}

//new object
func NewGOGPQueueNamePrefixQueue(bufSize int) *GOGPQueueNamePrefixQueue {
	r := &GOGPQueueNamePrefixQueue{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPQueueNamePrefixQueue) Init(bufSize int) {
	if nil == this.d {
		if bufSize <= 0 {
			bufSize = 8 //default buffer size
		}
		this.newBuf(bufSize)
	}
	this.Clear()
	return
}

func (this *GOGPQueueNamePrefixQueue) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPQueueElem, bufSize, bufSize) //the same cap and len
	}
}

//clear
func (this *GOGPQueueNamePrefixQueue) Clear() {
	this.head, this.tail = 0, 0
}

//push to head of queue
func (this *GOGPQueueNamePrefixQueue) Push(v GOGPQueueElem) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}
		if this.tail >= this.head { //normal order [head,tail)
			if this.tail < len(this.d) {
				this.d[this.tail] = v
				this.tail++
				if this.tail >= len(this.d) {
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
				oldSize := len(this.d)
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

//pop front of queue
func (this *GOGPQueueNamePrefixQueue) Pop() (front GOGPQueueElem, ok bool) {
	if this.head != this.tail {
		ok, front = true, this.d[this.head]
		if this.head++; this.head >= len(this.d) && this.head != this.tail {
			this.head = 0
		}
	}
	return
}

//size of queue
func (this *GOGPQueueNamePrefixQueue) Size() (size int) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = len(this.d) - this.head + this.tail
	}
	return
}

//if queye is empty
func (this *GOGPQueueNamePrefixQueue) Empty() bool {
	return this.Size() == 0
}

////show
//func (this *GOGPQueueNamePrefixQueue) Show() string {
//	var b show_bytes.Buffer
//	b.WriteByte('[')
//	for i := this.head; i != this.tail; i++ {
//		if i >= len(this.d) {
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
