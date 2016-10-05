//this file defines a template tree structure just like file system

<PACKAGE>

import (
	"sort"
)

//tree node
type <TREE_NAME_PREFIX>TreeNode struct {
	<TREE_NODE_DATA>
	Children <TREE_NAME_PREFIX>SortSlice
}

//create a visitor
func (this *<TREE_NAME_PREFIX>TreeNode) Visitor() (v *<TREE_NAME_PREFIX>TreeNodeVisitor) {
	v = &<TREE_NAME_PREFIX>TreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *<TREE_NAME_PREFIX>TreeNode) All() (list []<TREE_NODE_DATA>) {
	list = append(list, this.<TREE_NODE_DATA>)
	for _, v := range this.Children {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type <TREE_NAME_PREFIX>TreeNodeVisitor struct {
	node         *<TREE_NAME_PREFIX>TreeNode
	parents      []*<TREE_NAME_PREFIX>TreeNode
	brotherIdxes []int
	//visit order? this?child?brother?
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) push(n *<TREE_NAME_PREFIX>TreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) pop() (n *<TREE_NAME_PREFIX>TreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) tail() (n *<TREE_NAME_PREFIX>TreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) top_right(n *<TREE_NAME_PREFIX>TreeNode) (p *<TREE_NAME_PREFIX>TreeNode) {
	if n != nil {
		l := len(n.Children)
		for l > 0 {
			this.push(n, l-1)
			n = n.Children[l-1]
			l = len(n.Children)
		}
		p = n
	}
	return
}

//visit next node
func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) Next() bool {
	if this.node != nil { //check if has any children
		if len(this.node.Children) > 0 {
			this.push(this.node, 0)
			this.node = this.node.Children[0]
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < len(p.Children)-1 { //next brother
			bIdx++
			this.node = p.Children[bIdx]
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	return this.node != nil
}

//visit previous node
func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) Prev() bool {
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
			this.node = this.top_right(p.Children[bIdx])
		} else {
			this.node = p
			this.pop()
		}
	}
	return this.node != nil
}

//get node data
func (this *<TREE_NAME_PREFIX>TreeNodeVisitor) Get() *<TREE_NODE_DATA> {
	return &this.node.<TREE_NODE_DATA>
}

//for sort
type <TREE_NAME_PREFIX>SortSlice []*<TREE_NAME_PREFIX>TreeNode

func (this *<TREE_NAME_PREFIX>SortSlice) Sort() {
	sort.Sort(this)
}

//data
func (this *<TREE_NAME_PREFIX>SortSlice) D() []*<TREE_NAME_PREFIX>TreeNode {
	return *this
}

//push
func (this *<TREE_NAME_PREFIX>SortSlice) Push(v *<TREE_NAME_PREFIX>TreeNode) int {
	*this = append(*this, v)
	return this.Len()
}

func (this *<TREE_NAME_PREFIX>SortSlice) Pop() (r *<TREE_NAME_PREFIX>TreeNode) {
	if len(*this) > 0 {
		r = (*this)[len(*this)-1]
	}
	*this = (*this)[:len(*this)-1]
	return
}

//len
func (this *<TREE_NAME_PREFIX>SortSlice) Len() int {
	return len(this.D())
}

//sort by Hash decend,the larger one first
func (this *<TREE_NAME_PREFIX>SortSlice) Less(i, j int) bool {
	l, r := (*this)[i], (*this)[j]
	return l.Less(r.<TREE_NODE_DATA>)
}

//swap
func (this *<TREE_NAME_PREFIX>SortSlice) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

