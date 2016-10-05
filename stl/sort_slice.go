//this file define a template type for sort

package stl

//import (
//	"sort"
//)

////GOGPDummyDefineComment/*
////these defines will never exists in real go files
//type GOGPSliceElem int

//func (this GOGPSliceElem) Less(o GOGPSliceElem) bool {
//	return this < o
//}

//GOGPDummyDefineComment*/

////for sort
//type SortSliceGOGPNameSuffix []GOGPSliceElem

//func (this *SortSliceGOGPNameSuffix) Sort() {
//	sort.Sort(this)
//}

////data
//func (this *SortSliceGOGPNameSuffix) D() []GOGPSliceElem {
//	return *this
//}

////push
//func (this *SortSliceGOGPNameSuffix) Push(v GOGPSliceElem) int {
//	*this = append(*this, v)
//	return this.Len()
//}

//func (this *SortSliceGOGPNameSuffix) Pop() (r GOGPSliceElem) {
//	if len(*this) > 0 {
//		r = (*this)[len(*this)-1]
//	}
//	*this = (*this)[:len(*this)-1]
//	return
//}

////len
//func (this *SortSliceGOGPNameSuffix) Len() int {
//	return len(this.D())
//}

////sort by Hash decend,the larger one first
//func (this *SortSliceGOGPNameSuffix) Less(i, j int) bool {
//	l, r := (*this)[i], (*this)[j]
//	return l.Less(r)
//}

////swap
//func (this *SortSliceGOGPNameSuffix) Swap(i, j int) {
//	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
//}
