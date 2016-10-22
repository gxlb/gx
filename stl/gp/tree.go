//this file defines a template tree structure just like file system

package gp

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//
//
///*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
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

type GOGPValueType int

func (this GOGPValueType) Less(o GOGPValueType) bool {
	return this < o
}

//
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

//tree node
type GOGPContainerNamePrefixTree struct {
	root *GOGPContainerNamePrefixTreeNode
}

//new container
func NewGOGPContainerNamePrefixTree() *GOGPContainerNamePrefixTree {
	return &GOGPContainerNamePrefixTree{}
}

//tree node
type GOGPContainerNamePrefixTreeNode struct {
	GOGPValueType
	children _GOGPContainerNamePrefixTreeNodeSortSlice
}

func (this *GOGPContainerNamePrefixTreeNode) SortChildren() {
	this.children.Sort()
}

func (this *GOGPContainerNamePrefixTreeNode) Children() []*GOGPContainerNamePrefixTreeNode {
	return this.children.Buffer()
}

//add a child
func (this *GOGPContainerNamePrefixTreeNode) AddChild(v GOGPValueType, idx int) (child *GOGPContainerNamePrefixTreeNode) {
	n := &GOGPContainerNamePrefixTreeNode{GOGPValueType: v, children: nil}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *GOGPContainerNamePrefixTreeNode) AddChildNode(node *GOGPContainerNamePrefixTreeNode, idx int) (child *GOGPContainerNamePrefixTreeNode) {
	if idx >= 0 && idx < len(this.children) {
		right := this.children[idx+1:]
		this.children = append(this.children[:idx], node)
		this.children = append(this.children, right...)
	} else {
		this.children = append(this.children, node)
	}
	return node
}

//cound of children
func (this *GOGPContainerNamePrefixTreeNode) NumChildren() int {
	return len(this.children)
}

//get child
func (this *GOGPContainerNamePrefixTreeNode) GetChild(idx int) (child *GOGPContainerNamePrefixTreeNode, ok bool) {
	if ok = idx >= 0 && idx < len(this.children); ok {
		child = this.children[idx]
	}
	return
}

//remove child
func (this *GOGPContainerNamePrefixTreeNode) RemoveChild(idx int) (child *GOGPContainerNamePrefixTreeNode, ok bool) {
	if child, ok = this.GetChild(idx); ok {
		this.children = append(this.children[:idx], this.children[idx+1:]...)
	}
	return
}

//create a visitor
func (this *GOGPContainerNamePrefixTreeNode) Visitor() (v *GOGPContainerNamePrefixTreeNodeVisitor) {
	v = &GOGPContainerNamePrefixTreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *GOGPContainerNamePrefixTreeNode) All() (list []GOGPValueType) {
	list = append(list, this.GOGPValueType)
	for _, v := range this.children {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type GOGPContainerNamePrefixTreeNodeVisitor struct {
	node         *GOGPContainerNamePrefixTreeNode
	parents      []*GOGPContainerNamePrefixTreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) push(n *GOGPContainerNamePrefixTreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) pop() (n *GOGPContainerNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) tail() (n *GOGPContainerNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *GOGPContainerNamePrefixTreeNodeVisitor) top_right(n *GOGPContainerNamePrefixTreeNode) (p *GOGPContainerNamePrefixTreeNode) {
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
func (this *GOGPContainerNamePrefixTreeNodeVisitor) Next() (data *GOGPValueType, ok bool) {
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
	if ok = this.node != nil; ok {
		data = this.Get()
	}
	return
}

//visit previous node
func (this *GOGPContainerNamePrefixTreeNodeVisitor) Prev() (data *GOGPValueType, ok bool) {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		if ok = this.node != nil; ok {
			data = this.Get()
		}
		return
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
	if ok = this.node != nil; ok {
		data = this.Get()
	}
	return
}

//get node data
func (this *GOGPContainerNamePrefixTreeNodeVisitor) Get() *GOGPValueType {
	return &this.node.GOGPValueType
}

//for sort
type _GOGPContainerNamePrefixTreeNodeSortSlice []*GOGPContainerNamePrefixTreeNode

func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Sort() {
	sort.Sort(this)
}

//data
func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Buffer() []*GOGPContainerNamePrefixTreeNode {
	return *this
}

//push
func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Push(v *GOGPContainerNamePrefixTreeNode) int {
	*this = append(*this, v)
	return this.Len()
}

func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Pop() (r *GOGPContainerNamePrefixTreeNode) {
	if len(*this) > 0 {
		r = (*this)[len(*this)-1]
	}
	*this = (*this)[:len(*this)-1]
	return
}

//len
func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Len() int {
	return len(this.Buffer())
}

//sort by Hash decend,the larger one first
func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Less(i, j int) (ok bool) {
	l, r := (*this)[i], (*this)[j]

	//#if GOGP_HasLess
	ok = l.Less(r.GOGPValueType)
	//#else
	ok = l.GOGPValueType < r.GOGPValueType
	//#endif

	return
}

//swap
func (this *_GOGPContainerNamePrefixTreeNodeSortSlice) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile
