//fibonacci-heap

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

//import here...

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
import (
	dumy_fmt "fmt"
)

type GOGPFibHeapElem int

func (me GOGPFibHeapElem) Show() string {
	return dumy_fmt.Sprintf("%d", me)
} //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//heap node
type GOGPHeapNamePrefixFibheapNode struct {
	GOGPHeapElem
	degree        int
	marked        bool //if delete first child
	left, right   *GOGPHeapNamePrefixFibheapNode
	parent, child *GOGPHeapNamePrefixFibheapNode
}

//heap
type GOGPHeapNamePrefixFibheap struct {
	nodes     int
	maxDegree int
	min       *GOGPHeapNamePrefixFibheapNode
	cons      **GOGPHeapNamePrefixFibheapNode
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
