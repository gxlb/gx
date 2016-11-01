//this file defines a template tree structure just like file system

package gp

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
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
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//#GOGP_IGNORE_BEGIN
import "sort" //#GOGP_IGNORE_END

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/fakedef)
//these defines is used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/fakedef)



//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/functorcmp)
//this file is used to //import by other gp files
//it cannot use independently, simulation C++ stl functors
//package gp

//#GOGP_ONCE
const (
	CMPLesser = iota //default
	CMPGreater
) //
//#GOGP_ONCE_END

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePart byte

const (
	CmpGOGPGlobalNamePartLesser  CmpGOGPGlobalNamePart = CMPLesser
	CmpGOGPGlobalNamePartGreater CmpGOGPGlobalNamePart = CMPGreater
)

//create cmp object by name
func CreateCmpGOGPGlobalNamePart(cmpName string) (r CmpGOGPGlobalNamePart) {
	r = CmpGOGPGlobalNamePartLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePart) F(left, right GOGPValueType) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpGOGPGlobalNamePart) Lesser() CmpGOGPGlobalNamePart { return CMPLesser }

//Greater object
func (me CmpGOGPGlobalNamePart) Greater() CmpGOGPGlobalNamePart { return CMPGreater }

//show as string
func (me CmpGOGPGlobalNamePart) String() (s string) {
	switch me {
	case CMPLesser:
		s = "Lesser"
	case CMPGreater:
		s = "Greater"
	default:
		s = "error cmp value"
	}
	return
}

//create by bool
func (me CmpGOGPGlobalNamePart) CreateByBool(bigFirst bool) (r CmpGOGPGlobalNamePart) {
	if bigFirst {
		r = CMPGreater
	} else {
		r = CMPLesser
	}
	return
}

//create cmp object by name
func (me CmpGOGPGlobalNamePart) CreateByName(cmpName string) (r CmpGOGPGlobalNamePart) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = CMPLesser
	case "Greater":
		r = CMPGreater
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//lesser operation
func (me CmpGOGPGlobalNamePart) less(left, right GOGPValueType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePart) great(left, right GOGPValueType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)



//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/sort_slice)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/sort_slice)
//this file define a template type for sort

//package gp

//import "sort"

//func init() {
//	gGOGPGlobalNamePrefixSortSliceGbl.cmp = gGOGPGlobalNamePrefixSortSliceGbl.cmp.CreateByName("")
//}

//var gGOGPGlobalNamePrefixSortSliceGbl struct {
//	cmp CmpGOGPGlobalNamePart
//}

//new sort object
func NewGOGPGlobalNamePrefixSortSlice(capacity int, bigFirst bool) *GOGPGlobalNamePrefixSortSlice {
	p := &GOGPGlobalNamePrefixSortSlice{}
	p.Init(capacity, bigFirst)
	return p
}

//sort slice
type GOGPGlobalNamePrefixSortSlice struct {
	d   []GOGPValueType
	cmp CmpGOGPGlobalNamePart
}

//init
func (this *GOGPGlobalNamePrefixSortSlice) Init(capacity int, bigFirst bool) {
	this.d = make([]GOGPValueType, 0, capacity)
	this.cmp = this.cmp.CreateByBool(bigFirst)
}

//sort
func (this *GOGPGlobalNamePrefixSortSlice) Sort() {
	sort.Sort(this)
}

//data buffer
func (this *GOGPGlobalNamePrefixSortSlice) Buffer() []GOGPValueType {
	return this.d
}

//push
func (this *GOGPGlobalNamePrefixSortSlice) Push(v GOGPValueType) int {
	this.d = append(this.d, v)
	return this.Len()
}

//insert
func (this *GOGPGlobalNamePrefixSortSlice) Insert(v GOGPValueType, idx int) int {
	if idx >= 0 && idx < this.Len() {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], v)
		this.d = append(this.d, right...)
	} else {
		this.d = append(this.d, v)
	}
	return this.Len()
}

//remove
func (this *GOGPGlobalNamePrefixSortSlice) Remove(idx int) (ok bool) {
	if ok = idx >= 0 && idx < this.Len(); ok {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], right...)
	}
	return
}

//pop
func (this *GOGPGlobalNamePrefixSortSlice) Pop() (r GOGPValueType, ok bool) {
	if ok = len(this.d) > 0; ok {
		r = (this.d)[len(this.d)-1]
	}
	this.d = (this.d)[:len(this.d)-1]
	return
}

//len
func (this *GOGPGlobalNamePrefixSortSlice) Len() int {
	return len(this.d)
}

//sort by Hash decend,the larger one first
func (this *GOGPGlobalNamePrefixSortSlice) Less(i, j int) (ok bool) {
	l, r := (this.d)[i], (this.d)[j]
	return this.cmp.F(l, r)
}

//swap
func (this *GOGPGlobalNamePrefixSortSlice) Swap(i, j int) {
	(this.d)[i], (this.d)[j] = (this.d)[j], (this.d)[i]
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/sort_slice)

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
//
//add dummy defines here..........
//
//
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

func init() {
	gGOGPGlobalNamePrefixTreeGbl.cmp = gGOGPGlobalNamePrefixTreeGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

var gGOGPGlobalNamePrefixTreeGbl struct {
	cmp CmpGOGPGlobalNamePart
}

//tree strture
type GOGPGlobalNamePrefixTree struct {
	cmp  CmpGOGPGlobalNamePart
	root *GOGPGlobalNamePrefixTreeNode
}

//new container
func NewGOGPGlobalNamePrefixTree(bigFirst bool) *GOGPGlobalNamePrefixTree {
	p := &GOGPGlobalNamePrefixTree{cmp: gGOGPGlobalNamePrefixTreeGbl.cmp.CreateByBool(bigFirst)}
	return p
}

//tree node
type GOGPGlobalNamePrefixTreeNode struct {
	GOGPValueType
	children GOGPGlobalNamePrefixSortSlice
}

func (this *GOGPGlobalNamePrefixTreeNode) SortChildren() {
	this.children.Sort()
}

func (this *GOGPGlobalNamePrefixTreeNode) Children() []*GOGPGlobalNamePrefixTreeNode {
	return this.children.Buffer()
}

//add a child
func (this *GOGPGlobalNamePrefixTreeNode) AddChild(v GOGPValueType, idx int) (child *GOGPGlobalNamePrefixTreeNode) {
	n := &GOGPGlobalNamePrefixTreeNode{GOGPValueType: v, children: nil}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *GOGPGlobalNamePrefixTreeNode) AddChildNode(node *GOGPGlobalNamePrefixTreeNode, idx int) (child *GOGPGlobalNamePrefixTreeNode) {
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
func (this *GOGPGlobalNamePrefixTreeNode) NumChildren() int {
	return len(this.children)
}

//get child
func (this *GOGPGlobalNamePrefixTreeNode) GetChild(idx int) (child *GOGPGlobalNamePrefixTreeNode, ok bool) {
	if ok = idx >= 0 && idx < len(this.children); ok {
		child = this.children[idx]
	}
	return
}

//remove child
func (this *GOGPGlobalNamePrefixTreeNode) RemoveChild(idx int) (child *GOGPGlobalNamePrefixTreeNode, ok bool) {
	if child, ok = this.GetChild(idx); ok {
		this.children = append(this.children[:idx], this.children[idx+1:]...)
	}
	return
}

//create a visitor
func (this *GOGPGlobalNamePrefixTreeNode) Visitor() (v *GOGPGlobalNamePrefixTreeNodeVisitor) {
	v = &GOGPGlobalNamePrefixTreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *GOGPGlobalNamePrefixTreeNode) All() (list []GOGPValueType) {
	list = append(list, this.GOGPValueType)
	for _, v := range this.children {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type GOGPGlobalNamePrefixTreeNodeVisitor struct {
	node         *GOGPGlobalNamePrefixTreeNode
	parents      []*GOGPGlobalNamePrefixTreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) push(n *GOGPGlobalNamePrefixTreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) pop() (n *GOGPGlobalNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) tail() (n *GOGPGlobalNamePrefixTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *GOGPGlobalNamePrefixTreeNodeVisitor) top_right(n *GOGPGlobalNamePrefixTreeNode) (p *GOGPGlobalNamePrefixTreeNode) {
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
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Next() (data *GOGPValueType, ok bool) {
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
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Prev() (data *GOGPValueType, ok bool) {
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
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Get() *GOGPValueType {
	return &this.node.GOGPValueType
}

////for sort
//type __GOGPGlobalNamePrefixTreeNodeSortSlice []*GOGPGlobalNamePrefixTreeNode

//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Sort() {
//	sort.Sort(this)
//}

////data
//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Buffer() []*GOGPGlobalNamePrefixTreeNode {
//	return *this
//}

////push
//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Push(v *GOGPGlobalNamePrefixTreeNode) int {
//	*this = append(*this, v)
//	return this.Len()
//}

//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Pop() (r *GOGPGlobalNamePrefixTreeNode) {
//	if len(*this) > 0 {
//		r = (*this)[len(*this)-1]
//	}
//	*this = (*this)[:len(*this)-1]
//	return
//}

////len
//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Len() int {
//	return len(this.Buffer())
//}

////sort by Hash decend,the larger one first
//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Less(i, j int) (ok bool) {
//	l, r := (*this)[i], (*this)[j]
//	return gGOGPGlobalNamePrefixTreeGbl.cmp.F(l.GOGPValueType, r.GOGPValueType)
//}

////swap
//func (this *__GOGPGlobalNamePrefixTreeNodeSortSlice) Swap(i, j int) {
//	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
//}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile
