//single-way link list

package stl

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

//import here...

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//

type GOGPSListData int

//
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//list node
type GOGPListNamePrefixSListNode struct {
	GOGPSListData
	next *GOGPListNamePrefixSListNode
}

func (this *GOGPListNamePrefixSListNode) Get() *GOGPSListData {
	return &this.GOGPSListData
}

type GOGPListNamePrefixSListNodeVisitor struct {
	node, head *GOGPListNamePrefixSListNode
}

func (this *GOGPListNamePrefixSListNodeVisitor) Next() bool {
	return false
}

func (this *GOGPListNamePrefixSListNodeVisitor) Node() *GOGPListNamePrefixSListNode {
	return nil
}

//list object
type GOGPListNamePrefixSList struct {
	head, tail *GOGPListNamePrefixSListNode
}

//new object
func NewGOGPStackNamePrefixSList() *GOGPListNamePrefixSList {
	return &GOGPListNamePrefixSList{}
}

func (this *GOGPListNamePrefixSList) Len() int {
	return 0
}
func (this *GOGPListNamePrefixSList) Visitor(node *GOGPListNamePrefixSListNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
	return nil
}
func (this *GOGPListNamePrefixSList) Front() *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Back() *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Clear() *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushFront(v GOGPSListData) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushBack(v GOGPSListData) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushFrontList(other *GOGPListNamePrefixSList) {
	return
}
func (this *GOGPListNamePrefixSList) PushBackList(other *GOGPListNamePrefixSList) {
	return
}

func (this *GOGPListNamePrefixSList) InsertBefore(v GOGPSListData, mark *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) InsertAfter(v GOGPSListData, mark *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Remove(node *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveFront(node *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveBck(node *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveBefore(node, mark *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveAfter(node, mark *GOGPListNamePrefixSListNode) *GOGPListNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Sort() {
	return
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
