//this file define a template type for sort

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

import (
	"sort"
)

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPSliceElem int

func (this GOGPSliceElem) Less(o GOGPSliceElem) bool {
	return this < o
}

//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//for sort
type GOGPNamePrefixSortSlice []GOGPSliceElem

func (this *GOGPNamePrefixSortSlice) Sort() {
	sort.Sort(this)
}

//data
func (this *GOGPNamePrefixSortSlice) Slice() []GOGPSliceElem {
	return *this
}

//push
func (this *GOGPNamePrefixSortSlice) Push(v GOGPSliceElem) int {
	*this = append(*this, v)
	return this.Len()
}

func (this *GOGPNamePrefixSortSlice) Pop() (r GOGPSliceElem) {
	if len(*this) > 0 {
		r = (*this)[len(*this)-1]
	}
	*this = (*this)[:len(*this)-1]
	return
}

//len
func (this *GOGPNamePrefixSortSlice) Len() int {
	return len(this.Slice())
}

//sort by Hash decend,the larger one first
func (this *GOGPNamePrefixSortSlice) Less(i, j int) (ok bool) {
	l, r := (*this)[i], (*this)[j]
	//#GOGP_IFDEF GOGP_HasLess
	ok = l.Less(r)
	//#GOGP_ELSE
	ok = l < r
	//#GOGP_ENDIF
	return
}

//swap
func (this *GOGPNamePrefixSortSlice) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
