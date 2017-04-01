///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sat Apr 01 2017 22:48:09]
// Generate from:
//   [github.com/vipally/gx/stl/gp/sort_slice.gp]
//   [github.com/vipally/gx/fs/tree.gpg] [_tree_sort_slice]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt :
// Version : 3.0.0.final
//
///////////////////////////////////////////////////////////////////

//this file define a template type for sort

package fs

import "sort"

////////////////////////////////////////////////////////////////////////////////

var gFileHashSortSliceGbl struct {
	cmp CmpFileHash
}

func init() {
	gFileHashSortSliceGbl.cmp = gFileHashSortSliceGbl.cmp.CreateByName("")
}

//new sort object
func NewFileHashSortSlice(capacity int) *FileHashSortSlice {
	p := &FileHashSortSlice{}
	p.Init(capacity)
	return p
}

//sort slice
type FileHashSortSlice struct {
	d []*FileHashTreeNode
}

//init
func (this *FileHashSortSlice) Init(capacity int) {
	this.d = make([]*FileHashTreeNode, 0, capacity)
}

//sort
func (this *FileHashSortSlice) Sort() {
	sort.Sort(this)
}

//data buffer
func (this *FileHashSortSlice) Buffer() []*FileHashTreeNode {
	return this.d
}

//push
func (this *FileHashSortSlice) Push(v *FileHashTreeNode) int {
	this.d = append(this.d, v)
	return this.Len()
}

//insert
func (this *FileHashSortSlice) Insert(v *FileHashTreeNode, idx int) int {
	if idx >= 0 && idx < this.Len() {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], v)
		this.d = append(this.d, right...)
	} else {
		this.d = append(this.d, v)
	}
	return this.Len()
}

//remove
func (this *FileHashSortSlice) Remove(idx int) (r *FileHashTreeNode, ok bool) {
	if r, ok = this.Get(idx); ok {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], right...)
	}
	return
}

//pop
func (this *FileHashSortSlice) Pop() (r *FileHashTreeNode, ok bool) {
	if ok = len(this.d) > 0; ok {
		r = (this.d)[len(this.d)-1]
	}
	this.d = (this.d)[:len(this.d)-1]
	return
}

//get
func (this *FileHashSortSlice) Get(idx int) (r *FileHashTreeNode, ok bool) {
	if ok = idx >= 0 && idx < this.Len(); ok {
		r = this.d[idx]
	}
	return
}

//must get
func (this *FileHashSortSlice) MustGet(idx int) (r *FileHashTreeNode) {
	ok := false
	if r, ok = this.Get(idx); !ok {
		panic(idx)
	}
	return
}

//len
func (this *FileHashSortSlice) Len() int {
	return len(this.d)
}

//sort by Hash decend,the larger one first
func (this *FileHashSortSlice) Less(i, j int) (ok bool) {
	l, r := (this.d)[i], (this.d)[j]
	return gFileHashSortSliceGbl.cmp.F(l, r)
}

//swap
func (this *FileHashSortSlice) Swap(i, j int) {
	(this.d)[i], (this.d)[j] = (this.d)[j], (this.d)[i]
}
