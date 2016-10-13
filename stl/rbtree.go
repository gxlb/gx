//rb-tree

package stl

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
///*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
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
//these defines is used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//

type GOGPRBTreeData int

func (this GOGPRBTreeData) Less(o GOGPRBTreeData) bool {
	return this < o
}

//
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//list node
type GOGPTreeNamePrefixRBTreeNode struct {
	GOGPRBTreeData
	left, right *GOGPTreeNamePrefixRBTreeNode
}

func (this *GOGPTreeNamePrefixRBTreeNode) Get() *GOGPRBTreeData {
	return &this.GOGPRBTreeData
}

type GOGPTreeNamePrefixRBTreeNodeVisitor struct {
	node, head *GOGPTreeNamePrefixRBTreeNode
}

func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Next() bool {
	return false
}
func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Prev() bool {
	return false
}
func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Node() *GOGPTreeNamePrefixRBTreeNode {
	return nil
}

//list object
type GOGPTreeNamePrefixRBTree struct {
	head *GOGPTreeNamePrefixRBTreeNode
}

//new object
func NewGOGPStackNamePrefixDList() *GOGPTreeNamePrefixRBTree {
	return &GOGPTreeNamePrefixRBTree{}
}

func (this *GOGPTreeNamePrefixRBTree) Size() int {
	return 0
}
func (this *GOGPTreeNamePrefixRBTree) Visitor(node *GOGPTreeNamePrefixRBTreeNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Begin() *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) End() *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Clear() *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Insert(d GOGPRBTreeData) *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Delete(d GOGPRBTreeData) *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Remove(n *GOGPTreeNamePrefixRBTreeNode) *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) LowerBound(d GOGPRBTreeData) *GOGPTreeNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) UpperBound(d GOGPRBTreeData) *GOGPTreeNamePrefixRBTreeNode {
	return nil
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
