//rb-tree

package gp

//#GOGP_FILE_BEGIN

//import here...

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

type GOGPGlobalNamePrefixRBTreeNodeData struct {
}

//list node
type GOGPGlobalNamePrefixRBTreeNode struct {
	val         GOGPGlobalNamePrefixRBTreeNodeData
	left, right *GOGPGlobalNamePrefixRBTreeNode
}

func (this *GOGPGlobalNamePrefixRBTreeNode) Get() *GOGPGlobalNamePrefixRBTreeNodeData {
	return &this.val
}

type GOGPTreeNamePrefixRBTreeNodeVisitor struct {
	node, head *GOGPGlobalNamePrefixRBTreeNode
}

func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Next() bool {
	return false
}
func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Prev() bool {
	return false
}
func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Node() *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

//list object
type GOGPTreeNamePrefixRBTree struct {
	head *GOGPGlobalNamePrefixRBTreeNode
}

//new object
func NewGOGPStackNamePrefixDList() *GOGPTreeNamePrefixRBTree {
	return &GOGPTreeNamePrefixRBTree{}
}

func (this *GOGPTreeNamePrefixRBTree) Size() int {
	return 0
}
func (this *GOGPTreeNamePrefixRBTree) Visitor(node *GOGPGlobalNamePrefixRBTreeNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Begin() *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) End() *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Clear() *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Insert(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Delete(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Remove(n *GOGPGlobalNamePrefixRBTreeNode) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) LowerBound(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) UpperBound(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}
func (this *GOGPTreeNamePrefixRBTree) Find(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

//#GOGP_FILE_END
