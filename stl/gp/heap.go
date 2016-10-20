package gp

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
///*   <----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
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

func (this GOGPHeapElem) Less(o GOGPHeapElem) bool {
	return this < o
}

func (me GOGPHeapElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//stack object
type GOGPHeapNamePrefixHeap struct {
	d      []GOGPHeapElem
	limitN int
	maxTop bool
}

//new object
func NewGOGPHeapNamePrefixHeap(capacity int) (r *GOGPHeapNamePrefixHeap) {
	r = &GOGPHeapNamePrefixHeap{}
	r.Init(capacity, 0, false)
	return
}

//initialize
func (this *GOGPHeapNamePrefixHeap) Init(capacity, limitN int, max bool) {
	if cap(this.d) != capacity {
		this.d = make([]GOGPHeapElem, 0, capacity)
	}
	this.d = this.d[:0]

	this.limitN = limitN
	this.maxTop = max
}

//push
func (this *GOGPHeapNamePrefixHeap) Push(v GOGPHeapElem) {
	this.d = append(this.d, v)
	idx := this.Size() - 1
	parent := this.parent(idx)
	for ; idx > 0; idx, parent = parent, this.parent(parent) {
		if this.cmp(idx, parent) {
			this.swap(idx, parent)
		} else {
			break
		}
	}
}

//pop
func (this *GOGPHeapNamePrefixHeap) Pop() (top GOGPHeapElem, ok bool) {
	if top, ok = this.Top(); ok {
		this.swap(0, this.Size()-1)
		for i, l, r := 0, 1, 2; ; l, r = this.lchild(i), this.rchild(i) {
			if l == r {
			}
		}
	}
	return
}

func (this *GOGPHeapNamePrefixHeap) Top() (top GOGPHeapElem, ok bool) {
	if ok = !this.Empty(); ok {
		top = this.d[0]
	}
	return
}

func (this *GOGPHeapNamePrefixHeap) cmp(i, j int) (ok bool) {
	l, r := this.d[i], this.d[j]
	if this.maxTop {
		//#if GOGP_HasLess
		ok = r.Less(l)
		//#else
		ok = r < l
		//#endif
	} else {
		//#if GOGP_HasLess
		ok = l.Less(r)
		//#else
		ok = l < r
		//#endif
	}

	return
}
func (this *GOGPHeapNamePrefixHeap) swap(i, j int) {
	this.d[i], this.d[j] = this.d[j], this.d[i]
}

func (this *GOGPHeapNamePrefixHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (this *GOGPHeapNamePrefixHeap) lchild(idx int) int {
	return 2*idx + 1
}

func (this *GOGPHeapNamePrefixHeap) rchild(idx int) int {
	return 2*idx + 2
}

//size
func (this *GOGPHeapNamePrefixHeap) Size() int {
	return len(this.d)
}

func (this *GOGPHeapNamePrefixHeap) Empty() bool {
	return this.Size() == 0
}

func (this *GOGPHeapNamePrefixHeap) MakeHeap() {}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
