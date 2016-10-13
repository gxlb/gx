//this file defines a template tree structure just like file system

package stl

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

import (
	"sort"
)

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files

type GOGPTreeNodeData int

func (this GOGPTreeNodeData) Less(o GOGPTreeNodeData) bool {
	return this < o
}

//
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//tree node
type GOGPTreeNamePrefixTreeNode struct {
	GOGPTreeNodeData
	children treeNodeSortSlice
}

//create a visitor
func (this *GOGPTreeNamePrefixTreeNode) Visitor() (v *GOGPTreeNamePrefixTreeNodeVisitor) {
	v = &GOGPTreeNamePrefixTreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *GOGPTreeNamePrefixTreeNode) All() (list []GOGPTreeNodeData) {
	list = append(list, this.GOGPTreeNodeData)
	for _, v := range this.children {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type GOGPTreeNamePrefixTreeNodeVisitor struct {
	node         *GOGPTreeNamePrefixTreeNode
	parents      []*GOGPTreeNamePrefixTreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) push(n *GOGPTreeNamePrefixTreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) pop() (n *GOGPTreeNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) tail() (n *GOGPTreeNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *GOGPTreeNamePrefixTreeNodeVisitor) top_right(n *GOGPTreeNamePrefixTreeNode) (p *GOGPTreeNamePrefixTreeNode) {
	if n != nil {
		l := len(n.children)
		for l > 0 {
			this.push(n, l-1)
			n = n.children[l-1]
			l = len(n.children)
		}
		p = n
	}
	return
}

//visit next node
func (this *GOGPTreeNamePrefixTreeNodeVisitor) Next() bool {
	if this.node != nil { //check if has any children
		if len(this.node.children) > 0 {
			this.push(this.node, 0)
			this.node = this.node.children[0]
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < len(p.children)-1 { //next brother
			bIdx++
			this.node = p.children[bIdx]
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	return this.node != nil
}

//visit previous node
func (this *GOGPTreeNamePrefixTreeNodeVisitor) Prev() bool {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		return this.node != nil
	}

	if this.node != nil { //check if has any children
		p, bIdx := this.tail()
		if bIdx > 0 {
			bIdx--
			this.update_tail(bIdx)
			this.node = this.top_right(p.children[bIdx])
		} else {
			this.node = p
			this.pop()
		}
	}
	return this.node != nil
}

//get node data
func (this *GOGPTreeNamePrefixTreeNodeVisitor) Get() *GOGPTreeNodeData {
	return &this.node.GOGPTreeNodeData
}

//for sort
type treeNodeSortSlice []*GOGPTreeNamePrefixTreeNode

func (this *treeNodeSortSlice) Sort() {
	sort.Sort(this)
}

//data
func (this *treeNodeSortSlice) D() []*GOGPTreeNamePrefixTreeNode {
	return *this
}

//push
func (this *treeNodeSortSlice) Push(v *GOGPTreeNamePrefixTreeNode) int {
	*this = append(*this, v)
	return this.Len()
}

func (this *treeNodeSortSlice) Pop() (r *GOGPTreeNamePrefixTreeNode) {
	if len(*this) > 0 {
		r = (*this)[len(*this)-1]
	}
	*this = (*this)[:len(*this)-1]
	return
}

//len
func (this *treeNodeSortSlice) Len() int {
	return len(this.D())
}

//sort by Hash decend,the larger one first
func (this *treeNodeSortSlice) Less(i, j int) bool {
	l, r := (*this)[i], (*this)[j]
	return l.Less(r.GOGPTreeNodeData)
}

//swap
func (this *treeNodeSortSlice) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile
