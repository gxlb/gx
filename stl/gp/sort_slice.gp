//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Nov 13 2016 21:51:46]
// Generate from:
//   [github.com/vipally/gx/stl/gp/sort_slice.gp.go]
//   [github.com/vipally/gx/stl/gp/gp.gpg] [GOGP_REVERSE_sort_slice]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////
//#GOGP_IGNORE_END

//this file define a template type for sort

<PACKAGE>

import "sort"

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)

var g<GLOBAL_NAME_PREFIX>SortSliceGbl struct {
	cmp Cmp<GLOBAL_NAME_PREFIX>
}

func init() {
	g<GLOBAL_NAME_PREFIX>SortSliceGbl.cmp = g<GLOBAL_NAME_PREFIX>SortSliceGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//new sort object
func New<GLOBAL_NAME_PREFIX>SortSlice(capacity int) *<GLOBAL_NAME_PREFIX>SortSlice {
	p := &<GLOBAL_NAME_PREFIX>SortSlice{}
	p.Init(capacity)
	return p
}

//sort slice
type <GLOBAL_NAME_PREFIX>SortSlice struct {
	d []<VALUE_TYPE>
}

//init
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Init(capacity int) {
	this.d = make([]<VALUE_TYPE>, 0, capacity)
}

//sort
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Sort() {
	sort.Sort(this)
}

//data buffer
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Buffer() []<VALUE_TYPE> {
	return this.d
}

//push
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Push(v <VALUE_TYPE>) int {
	this.d = append(this.d, v)
	return this.Len()
}

//insert
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Insert(v <VALUE_TYPE>, idx int) int {
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
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Remove(idx int) (r <VALUE_TYPE>, ok bool) {
	if r, ok = this.Get(idx); ok {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], right...)
	}
	return
}

//pop
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Pop() (r <VALUE_TYPE>, ok bool) {
	if ok = len(this.d) > 0; ok {
		r = (this.d)[len(this.d)-1]
	}
	this.d = (this.d)[:len(this.d)-1]
	return
}

//get
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Get(idx int) (r <VALUE_TYPE>, ok bool) {
	if ok = idx >= 0 && idx < this.Len(); ok {
		r = this.d[idx]
	}
	return
}

//must get
func (this *<GLOBAL_NAME_PREFIX>SortSlice) MustGet(idx int) (r <VALUE_TYPE>) {
	ok := false
	if r, ok = this.Get(idx); !ok {
		panic(idx)
	}
	return
}

//len
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Len() int {
	return len(this.d)
}

//sort by Hash decend,the larger one first
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Less(i, j int) (ok bool) {
	l, r := (this.d)[i], (this.d)[j]
	return g<GLOBAL_NAME_PREFIX>SortSliceGbl.cmp.F(l, r)
}

//swap
func (this *<GLOBAL_NAME_PREFIX>SortSlice) Swap(i, j int) {
	(this.d)[i], (this.d)[j] = (this.d)[j], (this.d)[i]
}

