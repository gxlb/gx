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
	d      []GOGPHeapElem //data buffer
	limitN int            //if limitN>0, heap size must<=limitN
	maxTop bool           //if top is max value
}

//push heap value
func (this *GOGPHeapNamePrefixHeap) PushHeap(d []GOGPHeapElem, v GOGPHeapElem) []GOGPHeapElem {
	d = append(d, v)
	idx := len(d) - 1
	parent := this.parent(idx)
	for ; idx > 0; idx, parent = parent, this.parent(parent) {
		if this.cmpV(d[idx], d[parent]) {
			d[idx], d[parent] = d[parent], d[idx]
		} else {
			break
		}
	}
	return d
}

//pop heap top
func (this *GOGPHeapNamePrefixHeap) PopHeap(d []GOGPHeapElem) (h []GOGPHeapElem, top GOGPHeapElem, ok bool) {
	l := len(d)
	if ok = l > 0; ok {
		top = d[0]
		if l > 1 {
			d[0], d[l-1] = d[l-1], d[0]
			this.adjust(d[:l-1], 0, d[0])
		}
		h = d[:l-1]
	}
	return
}

//adjust heap to select a proper hole to set v
func (this *GOGPHeapNamePrefixHeap) adjust(d []GOGPHeapElem, hole int, v GOGPHeapElem) {}

//make d as h heap
func (this *GOGPHeapNamePrefixHeap) MakeHeap(d []GOGPHeapElem) {
	if l := len(d); l > 1 {
		for i := 0; i < l; i++ {
			this.PushHeap(d[:i], d[i])
		}
	}
}

//reverse order of d
func (this *GOGPHeapNamePrefixHeap) Reverse(d []GOGPHeapElem) {
	l := len(d) - 1
	m := l >> 1
	for i := 0; i <= m; i++ {
		d[i], d[l-i] = d[l-i], d[i]
	}
}

//sort slice use heap algorithm
func (this *GOGPHeapNamePrefixHeap) SortHeap(d []GOGPHeapElem) []GOGPHeapElem {
	this.MakeHeap(d)
	for t := d; len(t) > 1; {
		t, _, _ = this.PopHeap(t)
	}
	return d
}

//new object
func NewGOGPHeapNamePrefixHeap(capacity int, maxTop bool) (r *GOGPHeapNamePrefixHeap) {
	r = &GOGPHeapNamePrefixHeap{}
	r.Init(capacity, 0, maxTop)
	return
}

//initialize
func (this *GOGPHeapNamePrefixHeap) Init(capacity, limitN int, maxTop bool) {
	if cap(this.d) != capacity {
		this.d = make([]GOGPHeapElem, 0, capacity)
	}
	this.d = this.d[:0]

	this.limitN = limitN
	this.maxTop = maxTop
}

//heap slice
func (this *GOGPHeapNamePrefixHeap) D() []GOGPHeapElem {
	return this.d
}

//push
func (this *GOGPHeapNamePrefixHeap) Push(v GOGPHeapElem) {
	if this.limitN > 0 && this.Size() >= this.limitN {
		if top, ok := this.Top(); ok && this.cmpV(top, v) {
			this.Pop()
		}
	}
	this.d = this.PushHeap(this.d, v)
}

//pop
func (this *GOGPHeapNamePrefixHeap) Pop() (top GOGPHeapElem, ok bool) {
	this.d, top, ok = this.PopHeap(this.d)
	return
}

//heap top
func (this *GOGPHeapNamePrefixHeap) Top() (top GOGPHeapElem, ok bool) {
	if ok = !this.Empty(); ok {
		top = this.d[0]
	}
	return
}

//cmpare value
func (this *GOGPHeapNamePrefixHeap) cmpV(c, p GOGPHeapElem) (ok bool) {
	//#if GOGP_HasLess==true
	if this.maxTop {
		ok = !p.Less(c)
	} else {
		ok = !c.Less(p)
	}
	//#else
	if this.maxTop {
		ok = !(p < c)
	} else {
		ok = !(c < p)
	}
	//#endif
	return
}

//compare index
//func (this *GOGPHeapNamePrefixHeap) cmp(c, p int) (ok bool) {
//	this.cmpV(this.d[c], this.d[p])
//	return
//}

//func (this *GOGPHeapNamePrefixHeap) swap(i, j int) {
//	this.d[i], this.d[j] = this.d[j], this.d[i]
//}

//get parent index
func (this *GOGPHeapNamePrefixHeap) parent(idx int) int {
	return (idx - 1) / 2
}

//get left child index
func (this *GOGPHeapNamePrefixHeap) lchild(idx int) int {
	return 2*idx + 1
}

//get right child index
func (this *GOGPHeapNamePrefixHeap) rchild(idx int) int {
	return 2*idx + 2
}

//size
func (this *GOGPHeapNamePrefixHeap) Size() int {
	return len(this.d)
}

//empty
func (this *GOGPHeapNamePrefixHeap) Empty() bool {
	return this.Size() == 0
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
