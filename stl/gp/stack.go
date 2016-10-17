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

type GOGPStackElem int

func (me GOGPStackElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//stack object
type GOGPStackNamePrefixStack struct {
	d []GOGPStackElem
}

//new object
func NewGOGPStackNamePrefixStack(bufSize int) *GOGPStackNamePrefixStack {
	r := &GOGPStackNamePrefixStack{}
	r.Init(bufSize)
	return r
}

//init
func (this *GOGPStackNamePrefixStack) Init(bufSize int) {
	if nil == this.d {
		if -1 == bufSize {
			bufSize = 10
		}
		if bufSize > 0 {
			this.d = make([]GOGPStackElem, 0, bufSize)
		}
	}
	this.d = this.d[:0]
	return
}

//clear
func (this *GOGPStackNamePrefixStack) Clear() {
	this.Init(-1)
}

//push v to top of stack
func (this *GOGPStackNamePrefixStack) Push(v GOGPStackElem) (ok bool) {
	if ok = true; ok {
		this.d = append(this.d, v)
	}
	return
}

//pop top of stack
func (this *GOGPStackNamePrefixStack) Pop() (top GOGPStackElem, ok bool) {
	if top, ok = this.Top(); ok {
		this.d = this.d[:this.Size()-1]
	}
	return
}

//get top of stack
func (this *GOGPStackNamePrefixStack) Top() (top GOGPStackElem, ok bool) {
	d := this.Size()
	if d > 0 {
		top = this.d[d-1]
		ok = true
	}
	return

}

//size of stack
func (this *GOGPStackNamePrefixStack) Size() int {
	return len(this.d)
}

//is stack is empty
func (this *GOGPStackNamePrefixStack) Empty() bool {
	return this.Size() == 0
}

//GOGPShowComment//show
//GOGPShowCommentfunc (this *GOGPStackNamePrefixStack) Show() string {
//GOGPShowComment	var b show_bytes.Buffer
//GOGPShowComment	b.WriteByte('[')
//GOGPShowComment	for _, v := range this.d {
//GOGPShowComment		b.WriteString(v.Show())
//GOGPShowComment		b.WriteByte(',')
//GOGPShowComment	}
//GOGPShowComment	if this.Depth() > 0 {
//GOGPShowComment		b.Truncate(b.Len() - 1) //remove last ','
//GOGPShowComment	}
//GOGPShowComment	b.WriteByte(']')
//GOGPShowComment	return b.String()
//GOGPShowComment}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
