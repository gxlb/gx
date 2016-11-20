//this file defines a template tree structure just like file system

package gp

//#GOGP_FILE_BEGIN
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

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



//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp,#GOGP_GPGCFG(GOGP_SectionSortSlice))
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/functorcmp)
//this file is used to //import by other gp files
//it cannot use independently, simulation C++ stl functors
//package gp

const (
	CMPLesser = iota //default
	CMPGreater
) //

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePrefix byte

const (
	CmpGOGPGlobalNamePrefixLesser  CmpGOGPGlobalNamePrefix = CMPLesser
	CmpGOGPGlobalNamePrefixGreater CmpGOGPGlobalNamePrefix = CMPGreater
)

//create cmp object by name
func CreateCmpGOGPGlobalNamePrefix(cmpName string) (r CmpGOGPGlobalNamePrefix) {
	r = CmpGOGPGlobalNamePrefixLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePrefix) F(left, right *GOGPGlobalNamePrefixTreeNode) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpGOGPGlobalNamePrefix) Lesser() CmpGOGPGlobalNamePrefix { return CMPLesser }

//Greater object
func (me CmpGOGPGlobalNamePrefix) Greater() CmpGOGPGlobalNamePrefix { return CMPGreater }

//show as string
func (me CmpGOGPGlobalNamePrefix) String() (s string) {
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
func (me CmpGOGPGlobalNamePrefix) CreateByBool(bigFirst bool) (r CmpGOGPGlobalNamePrefix) {
	if bigFirst {
		r = CMPGreater
	} else {
		r = CMPLesser
	}
	return
}

//create cmp object by name
func (me CmpGOGPGlobalNamePrefix) CreateByName(cmpName string) (r CmpGOGPGlobalNamePrefix) {
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
func (me CmpGOGPGlobalNamePrefix) less(left, right *GOGPGlobalNamePrefixTreeNode) (ok bool) {

	ok = left.Less(right)

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePrefix) great(left, right *GOGPGlobalNamePrefixTreeNode) (ok bool) {

	ok = right.Less(left)

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)



//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/sort_slice,#GOGP_GPGCFG(GOGP_SectionSortSlice))
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/sort_slice)
//this file define a template type for sort

//package gp

//import "sort"

var gGOGPGlobalNamePrefixSortSliceGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixSortSliceGbl.cmp = gGOGPGlobalNamePrefixSortSliceGbl.cmp.CreateByName("")
}

//new sort object
func NewGOGPGlobalNamePrefixSortSlice(capacity int) *GOGPGlobalNamePrefixSortSlice {
	p := &GOGPGlobalNamePrefixSortSlice{}
	p.Init(capacity)
	return p
}

//sort slice
type GOGPGlobalNamePrefixSortSlice struct {
	d []*GOGPGlobalNamePrefixTreeNode
}

//init
func (this *GOGPGlobalNamePrefixSortSlice) Init(capacity int) {
	this.d = make([]*GOGPGlobalNamePrefixTreeNode, 0, capacity)
}

//sort
func (this *GOGPGlobalNamePrefixSortSlice) Sort() {
	sort.Sort(this)
}

//data buffer
func (this *GOGPGlobalNamePrefixSortSlice) Buffer() []*GOGPGlobalNamePrefixTreeNode {
	return this.d
}

//push
func (this *GOGPGlobalNamePrefixSortSlice) Push(v *GOGPGlobalNamePrefixTreeNode) int {
	this.d = append(this.d, v)
	return this.Len()
}

//insert
func (this *GOGPGlobalNamePrefixSortSlice) Insert(v *GOGPGlobalNamePrefixTreeNode, idx int) int {
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
func (this *GOGPGlobalNamePrefixSortSlice) Remove(idx int) (r *GOGPGlobalNamePrefixTreeNode, ok bool) {
	if r, ok = this.Get(idx); ok {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], right...)
	}
	return
}

//pop
func (this *GOGPGlobalNamePrefixSortSlice) Pop() (r *GOGPGlobalNamePrefixTreeNode, ok bool) {
	if ok = len(this.d) > 0; ok {
		r = (this.d)[len(this.d)-1]
	}
	this.d = (this.d)[:len(this.d)-1]
	return
}

//get
func (this *GOGPGlobalNamePrefixSortSlice) Get(idx int) (r *GOGPGlobalNamePrefixTreeNode, ok bool) {
	if ok = idx >= 0 && idx < this.Len(); ok {
		r = this.d[idx]
	}
	return
}

//must get
func (this *GOGPGlobalNamePrefixSortSlice) MustGet(idx int) (r *GOGPGlobalNamePrefixTreeNode) {
	ok := false
	if r, ok = this.Get(idx); !ok {
		panic(idx)
	}
	return
}

//len
func (this *GOGPGlobalNamePrefixSortSlice) Len() int {
	return len(this.d)
}

//sort by Hash decend,the larger one first
func (this *GOGPGlobalNamePrefixSortSlice) Less(i, j int) (ok bool) {
	l, r := (this.d)[i], (this.d)[j]
	return gGOGPGlobalNamePrefixSortSliceGbl.cmp.F(l, r)
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

//var gGOGPGlobalNamePrefixTreeGbl struct {
//	cmp CmpGOGPGlobalNamePrefix
//}
//func init() {
//	gGOGPGlobalNamePrefixTreeGbl.cmp = gGOGPGlobalNamePrefixTreeGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
//}

//tree strture
type GOGPGlobalNamePrefixTree struct {
	root *GOGPGlobalNamePrefixTreeNode
}

//new container
func NewGOGPGlobalNamePrefixTree() *GOGPGlobalNamePrefixTree {
	p := &GOGPGlobalNamePrefixTree{}
	return p
}

//tree node
type GOGPGlobalNamePrefixTreeNode struct {
	GOGPValueType
	children GOGPGlobalNamePrefixSortSlice
}

func (this *GOGPGlobalNamePrefixTreeNode) Less(right *GOGPGlobalNamePrefixTreeNode) (ok bool) {
	//#GOGP_IFDEF GOGP_HasCmpFunc
	ok = this.GOGPValueType.Less(right.GOGPValueType)
	//#GOGP_ELSE
	ok = this.GOGPValueType < right.GOGPValueType
	//#GOGP_ENDIF
	return
}

func (this *GOGPGlobalNamePrefixTreeNode) SortChildren() {
	this.children.Sort()
}

func (this *GOGPGlobalNamePrefixTreeNode) Children() []*GOGPGlobalNamePrefixTreeNode {
	return this.children.Buffer()
}

//add a child
func (this *GOGPGlobalNamePrefixTreeNode) AddChild(v GOGPValueType, idx int) (child *GOGPGlobalNamePrefixTreeNode) {
	n := &GOGPGlobalNamePrefixTreeNode{GOGPValueType: v}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *GOGPGlobalNamePrefixTreeNode) AddChildNode(node *GOGPGlobalNamePrefixTreeNode, idx int) (child *GOGPGlobalNamePrefixTreeNode) {
	this.children.Insert(node, idx)
	return node
}

//cound of children
func (this *GOGPGlobalNamePrefixTreeNode) NumChildren() int {
	return this.children.Len()
}

//get child
func (this *GOGPGlobalNamePrefixTreeNode) GetChild(idx int) (child *GOGPGlobalNamePrefixTreeNode, ok bool) {
	child, ok = this.children.Get(idx)
	return
}

//remove child
func (this *GOGPGlobalNamePrefixTreeNode) RemoveChild(idx int) (child *GOGPGlobalNamePrefixTreeNode, ok bool) {
	child, ok = this.children.Remove(idx)
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
	for _, v := range this.children.Buffer() {
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
		l := n.children.Len()
		for l > 0 {
			this.push(n, l-1)
			n = n.children.MustGet(l - 1)
			l = n.children.Len()
		}
		p = n
	}
	return
}

//visit next node
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Next() (ok bool) {
	if this.node != nil { //check if has any children
		if this.node.children.Len() > 0 {
			this.push(this.node, 0)
			this.node = this.node.children.MustGet(0)
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < p.children.Len()-1 { //next brother
			bIdx++
			this.node = p.children.MustGet(bIdx)
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//visit previous node
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Prev() (ok bool) {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		if ok = this.node != nil; ok {
			//do nothing
		}
		return
	}

	if this.node != nil { //check if has any children
		p, bIdx := this.tail()
		if bIdx > 0 {
			bIdx--
			this.update_tail(bIdx)
			this.node = this.top_right(p.children.MustGet(bIdx))
		} else {
			this.node = p
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//get node data
func (this *GOGPGlobalNamePrefixTreeNodeVisitor) Get() (data *GOGPValueType) {
	if nil != this.node {
		data = &this.node.GOGPValueType
	}
	return
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
