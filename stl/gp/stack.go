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

//#GOGP_IFDEF GOGP_HasShow
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

//stack object
type GOGPGlobalNamePrefixStack struct {
	d []GOGPValueType
}

//new object
func NewGOGPStackNamePrefixStack(bufSize int) *GOGPGlobalNamePrefixStack {
	r := &GOGPGlobalNamePrefixStack{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPGlobalNamePrefixStack) Init(bufSize int) {
	if nil == this.d {
		if -1 == bufSize {
			bufSize = 10
		}
		if bufSize > 0 {
			this.d = make([]GOGPValueType, 0, bufSize)
		}
	}
	this.d = this.d[:0]
	return
}

//clear
func (this *GOGPGlobalNamePrefixStack) Clear() {
	this.Init(-1)
}

//push v to top of stack
func (this *GOGPGlobalNamePrefixStack) Push(v GOGPValueType) (ok bool) {
	if ok = true; ok {
		this.d = append(this.d, v)
	}
	return
}

//pop top of stack
func (this *GOGPGlobalNamePrefixStack) Pop() (top GOGPValueType, ok bool) {
	if top, ok = this.Top(); ok {
		this.d = this.d[:this.Size()-1]
	}
	return
}

//get top of stack
func (this *GOGPGlobalNamePrefixStack) Top() (top GOGPValueType, ok bool) {
	d := this.Size()
	if d > 0 {
		top = this.d[d-1]
		ok = true
	}
	return

}

//size of stack
func (this *GOGPGlobalNamePrefixStack) Size() int {
	return len(this.d)
}

//is stack is empty
func (this *GOGPGlobalNamePrefixStack) Empty() bool {
	return this.Size() == 0
}

//#GOGP_IFDEF GOGP_HasShow
//show
func (this *GOGPGlobalNamePrefixStack) Show() string {
	var b show_bytes.Buffer
	b.WriteByte('[')
	for _, v := range this.d {
		b.WriteString(v.Show())
		b.WriteByte(',')
	}
	if this.Size() > 0 {
		b.Truncate(b.Len() - 1) //remove last ','
	}
	b.WriteByte(']')
	return b.String()
} //
//#GOGP_ENDIF

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
