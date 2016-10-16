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

//GOGPShowCommentimport show_bytes "bytes"

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
type GOGPStackNamePrefixStack []GOGPStackElem

//new object
func NewGOGPStackNamePrefixStack() *GOGPStackNamePrefixStack {
	return &GOGPStackNamePrefixStack{}
}

//push
func (this *GOGPStackNamePrefixStack) Push(v GOGPStackElem) {
	*this = append(*this, v)
}

//pop
func (this *GOGPStackNamePrefixStack) Pop() (top GOGPStackElem, ok bool) {
	if top, ok = this.Top(); ok {
		*this = (*this)[:this.Depth()-1]
	}
	return
}

//top
func (this *GOGPStackNamePrefixStack) Top() (top GOGPStackElem, ok bool) {
	if this.Depth() > 0 {
		top = (*this)[this.Depth()-1]
		ok = true
	}
	return

}

//depth
func (this *GOGPStackNamePrefixStack) Depth() int {
	return len(*this)
}

//GOGPShowComment//show
//GOGPShowCommentfunc (this *GOGPStackNamePrefixStack) Show() string {
//GOGPShowComment	var b show_bytes.Buffer
//GOGPShowComment	b.WriteByte('[')
//GOGPShowComment	for _, v := range *this {
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
