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

//import...

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

import (
	dumy_fmt "fmt"
)

type GOGPHeapElem int

func (me GOGPHeapElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//stack object
type GOGPHeapNamePrefixHeap []GOGPHeapElem

//new object
func NewGOGPHeapNamePrefixHeap() *GOGPHeapNamePrefixHeap {
	return &GOGPHeapNamePrefixHeap{}
}

//push
func (this *GOGPHeapNamePrefixHeap) Push(v GOGPHeapElem) {
	*this = append(*this, v)
}

//pop
func (this *GOGPHeapNamePrefixHeap) Pop() (top GOGPHeapElem, ok bool) {
	if top, ok = this.Top(); ok {
		*this = (*this)[:this.Depth()-1]
	}
	return
}

//top
func (this *GOGPHeapNamePrefixHeap) Sort() {
	if this.Depth() > 0 {
		top = (*this)[this.Depth()-1]
		ok = true
	}
	return

}

//depth
func (this *GOGPHeapNamePrefixHeap) Depth() int {
	return len(*this)
}

func (this *GOGPHeapNamePrefixHeap) MakeHeap() {}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
