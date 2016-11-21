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

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/fakedef)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/functorcmp)
//this file is used to //import by other gp files
//it cannot use independently, simulation C++ stl functors

//package gp

const (
	CMPLesser = iota //default
	CMPGreater
) //

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePrefix byte

const (
	CmpGOGPGlobalNamePrefixLesser  CmpGOGPGlobalNamePrefix = CMPLesser
	CmpGOGPGlobalNamePrefixGreater CmpGOGPGlobalNamePrefix = CMPGreater
)

//create cmp object by name
func CreateCmpGOGPGlobalNamePrefix(cmpName string) (r CmpGOGPGlobalNamePrefix) {
	r = CmpGOGPGlobalNamePrefixLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePrefix) F(left, right GOGPValueType) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpGOGPGlobalNamePrefix) Lesser() CmpGOGPGlobalNamePrefix { return CMPLesser }

//Greater object
func (me CmpGOGPGlobalNamePrefix) Greater() CmpGOGPGlobalNamePrefix { return CMPGreater }

//show as string
func (me CmpGOGPGlobalNamePrefix) String() (s string) {
	switch me {
	case CMPLesser:
		s = "Lesser"
	case CMPGreater:
		s = "Greater"
	default:
		s = "error cmp value"
	}
	return
}

//create by bool
func (me CmpGOGPGlobalNamePrefix) CreateByBool(bigFirst bool) (r CmpGOGPGlobalNamePrefix) {
	if bigFirst {
		r = CMPGreater
	} else {
		r = CMPLesser
	}
	return
}

//create cmp object by name
func (me CmpGOGPGlobalNamePrefix) CreateByName(cmpName string) (r CmpGOGPGlobalNamePrefix) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = CMPLesser
	case "Greater":
		r = CMPGreater
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//lesser operation
func (me CmpGOGPGlobalNamePrefix) less(left, right GOGPValueType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePrefix) great(left, right GOGPValueType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

//queue object
type GOGPGlobalNamePrefixQueue struct {
	//real data is [head,tail)
	//buffer d is cycle, that is to say, next(len(d)-1)=0, prev(0)=len(d)-1
	//so if tail<head, data is [head, end, 0, tail)
	head int
	tail int
	d    []GOGPValueType
}

//new object
func NewGOGPQueueNamePrefixQueue(bufSize int) *GOGPGlobalNamePrefixQueue {
	r := &GOGPGlobalNamePrefixQueue{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPGlobalNamePrefixQueue) Init(bufSize int) {
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
func (this *GOGPGlobalNamePrefixQueue) newBuf(bufSize int) {
	if bufSize > 0 {
		this.d = make([]GOGPValueType, bufSize, bufSize) //the same cap and len
	}
}

//clear
func (this *GOGPGlobalNamePrefixQueue) Clear() {
	this.head, this.tail = 0, 0
}

//push to back of queue
func (this *GOGPGlobalNamePrefixQueue) Push(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		if nil == this.d { //init if needed
			this.Init(-1)
		}
		this.d[this.tail] = v
		if this.tail++; this.tail >= this.Cap() {
			this.tail = 0
			if this.tail == this.head { //tail catch up head, buffer full
				oldCap := this.Cap()
				d := this.d
				this.newBuf(oldCap * 2)
				h := copy(this.d, d[this.head:])
				t := copy(this.d[:h], d[:this.tail])
				this.head, this.tail = 0, h+t
			}
		}
	}
	return
}

//pop front of queue
func (this *GOGPGlobalNamePrefixQueue) Pop() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
		if this.head++; this.head >= this.Cap() && this.head != this.tail {
			this.head = 0
		}
	}
	return
}

//front data
func (this *GOGPGlobalNamePrefixQueue) Front() (front GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		front = this.d[this.head]
	}
	return
}

//back data
func (this *GOGPGlobalNamePrefixQueue) Back() (back GOGPValueType, ok bool) {
	if ok = this.head != this.tail; ok {
		t := this.tail - 1
		if t < 0 {
			t = this.Cap() - 1
		}
		back = this.d[t]
	}
	return
}

//shrink data buffer if necessary
func (this *GOGPGlobalNamePrefixQueue) Shrink() (ok bool) {
	oldCap := this.Cap()
	oldSize := this.Size()
	if ok := oldCap > 8 && oldCap >= 3*oldSize; ok { //leave at least 8 elem space
		d := this.d
		this.newBuf(oldSize / 2)
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
func (this *GOGPGlobalNamePrefixQueue) Cap() int {
	return len(this.d)
}

//size of queue
func (this *GOGPGlobalNamePrefixQueue) Size() (size int) {
	if this.tail >= this.head {
		size = this.tail - this.head
	} else {
		size = this.Cap() - this.head + this.tail
	}
	return
}

//if queue is empty
func (this *GOGPGlobalNamePrefixQueue) Empty() bool {
	return this.Size() == 0
}

////show
//func (this *GOGPQueueNamePrefixQueue) Show() string {
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

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
