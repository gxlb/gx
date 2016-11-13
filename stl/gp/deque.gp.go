package gp

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
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
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//#GOGP_IFDEF GOGP_Show
import show_bytes "bytes" //#GOGP_ENDIF

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/fakedef)
//these defines is used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/fakedef)

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//deque object
type GOGPGlobalNamePrefixDeque struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	//head point to the first elem  available for read
	//tail point to the first space available for write
	head int
	tail int
	d    []GOGPValueType
}

//new object
func NewGOGPDequeNamePrefixDeque(bufSize int) *GOGPGlobalNamePrefixDeque {
	r := &GOGPGlobalNamePrefixDeque{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPGlobalNamePrefixDeque) Init(bufSize int) {
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
func (this *GOGPGlobalNamePrefixDeque) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPValueType, bufSize, bufSize) //the same cap and len
	}
}

//clear all deque data
func (this *GOGPGlobalNamePrefixDeque) Clear() {
	this.head, this.tail = 0, 0
}

//push to front of deque
func (this *GOGPGlobalNamePrefixDeque) PushFront(v GOGPValueType) (ok bool) {
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
			this.head, this.tail = 0, h+t
		}
	}
	return
}

//push to back of deque
func (this *GOGPGlobalNamePrefixDeque) PushBack(v GOGPValueType) (ok bool) {
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
			this.head, this.tail = 0, h+t
		}
	}
	return
}

//pop front of deque
func (this *GOGPGlobalNamePrefixDeque) PopFront() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
		this.head = this.next(this.head)
	}
	return
}

//pop back of deque
func (this *GOGPGlobalNamePrefixDeque) PopBack() (back GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		this.tail = this.prev(this.tail)
		back = this.d[this.tail]
	}
	return
}

//front data
func (this *GOGPGlobalNamePrefixDeque) Front() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
	}
	return
}

//back data
func (this *GOGPGlobalNamePrefixDeque) Back() (back GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		t := this.prev(this.tail)
		back = this.d[t]
	}
	return
}

//shrink data buffer if necessary
func (this *GOGPGlobalNamePrefixDeque) Shrink() (ok bool) {
	oldCap := this.Cap()
	oldSize := this.Size()
	if ok := oldCap > 8 && oldCap >= 3*oldSize; ok { //leave at least 8 elem space
		d := this.d
		this.newBuf(oldCap / 2)
		if this.tail >= this.head {
			copy(this.d, d[this.head:this.tail])
			this.tail -= this.head
			this.head = 0
		} else {
			h := copy(this.d, d[this.head:])
			t := copy(this.d[:h], d[:this.tail])
			this.head, this.tail = 0, h+t
		}
	}
	return
}

//func (this *GOGPDequeNamePrefixDeque) Sort() {}

//data buffer size
func (this *GOGPGlobalNamePrefixDeque) Cap() int {
	return len(this.d)
}

//size of deque
func (this *GOGPGlobalNamePrefixDeque) Size() (size int) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = this.Cap() - (this.head - this.tail)
	}
	return
}

//if deque is empty
func (this *GOGPGlobalNamePrefixDeque) Empty() bool {
	return this.Size() == 0
}

//next buff
func (this *GOGPGlobalNamePrefixDeque) next(idx int) (r int) {
	if r = idx + 1; r >= this.Cap() {
		r = 0
	}
	return
}

//prev buff
func (this *GOGPGlobalNamePrefixDeque) prev(idx int) (r int) {
	if r = idx - 1; r < 0 {
		r = this.Cap() - 1
	}
	return
}

//#GOGP_IFDEF GOGP_Show
//show
func (this *GOGPGlobalNamePrefixDeque) Show() string {
	var b show_bytes.Buffer
	b.WriteByte('[')
	for i := this.head; i != this.tail; i++ {
		if i >= this.Cap() {
			i = 0
		}
		v := this.d[i]
		b.WriteString(v.Show())
		b.WriteByte(',')
	}
	if this.Size() > 0 {
		b.Truncate(b.Len() - 1) //remove last ','
	}
	b.WriteByte(']')
	return b.String()
} //#GOGP_ENDIF //GOGP_Show

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
