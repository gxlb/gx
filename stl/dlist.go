//double-way cycle link list

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

type GOGPDListData int

//
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//list node
type GOGPListNamePrefixDListNode struct {
	GOGPDListData
	prev, next *GOGPListNamePrefixDListNode
}

func (this *GOGPListNamePrefixDListNode) Get() *GOGPDListData {
	return &this.GOGPDListData
}

type GOGPListNamePrefixDListNodeVisitor struct {
	node, head *GOGPListNamePrefixDListNode
}

func (this *GOGPListNamePrefixDListNodeVisitor) Next() bool {
	return false
}
func (this *GOGPListNamePrefixDListNodeVisitor) Prev() bool {
	return false
}
func (this *GOGPListNamePrefixDListNodeVisitor) Node() *GOGPListNamePrefixDListNode {
	return nil
}

//list object
type GOGPListNamePrefixDList struct {
	head *GOGPListNamePrefixDListNode
}

//new object
func NewGOGPStackNamePrefixDList() *GOGPListNamePrefixDList {
	return &GOGPListNamePrefixDList{}
}

func (this *GOGPListNamePrefixDList) Len() int {
	return 0
}
func (this *GOGPListNamePrefixDList) Visitor(node *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNodeVisitor {
	return nil
}
func (this *GOGPListNamePrefixDList) Front() *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) Back() *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) Clear() *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) PushFront(v GOGPDListData) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) PushBack(v GOGPDListData) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) PushFrontList(other *GOGPListNamePrefixDList) {
	return
}
func (this *GOGPListNamePrefixDList) PushBackList(other *GOGPListNamePrefixDList) {
	return
}

func (this *GOGPListNamePrefixDList) InsertBefore(v GOGPDListData, mark *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) InsertAfter(v GOGPDListData, mark *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) Remove(node *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) MoveFront(node *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) MoveBck(node *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) MoveBefore(node, mark *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) MoveAfter(node, mark *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNode {
	return nil
}
func (this *GOGPListNamePrefixDList) Sort() {
	return
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
