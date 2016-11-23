package stl

import (
	"fmt"
	"testing"
	//"time"
	//"regexp"

	//_ "github.com/vipally/gogp/auto"
)

func init() {
	//	s := []GOGPSliceElem{6, 7, 8, 3, 4, 6}
	//	//MakeSlice(s).Sort()
	//	ss := GOGPTreeNamePrefixSortSlice(s)
	//	ss.Sort()
	//	ss.Pop()
	//	for _, v := range ss.D() {
	//		fmt.Println(v)
	//	}
}

func TestHeap(t *testing.T) {
	a := []int{4, 5, 6, 2, 9, 8, 4, 7}
	h := NewIntHeap(0, 6, false)
	for _, v := range a {
		h.Push(v)
	}
	for !h.Empty() {
		v, _ := h.Pop()
		fmt.Printf("%d ", v)
	}
}
