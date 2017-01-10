//rb-tree

package gp

//#GOGP_FILE_BEGIN 1
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
///*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPKeyType int                             //
func (this GOGPKeyType) Less(o GOGPKeyType) bool { return this < o }
func (this GOGPKeyType) Show() string            { return "" } //

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gogp/lib/fakedef)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
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
func (me CmpGOGPGlobalNamePrefix) F(left, right GOGPKeyType) (ok bool) {
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
func (me CmpGOGPGlobalNamePrefix) less(left, right GOGPKeyType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePrefix) great(left, right GOGPKeyType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

//var gGOGPGlobalNamePrefixRBTreeGbl struct {
//	cmp CmpGOGPGlobalNamePrefix
//}

//func init() {
//	gGOGPGlobalNamePrefixRBTreeGbl.cmp = gGOGPGlobalNamePrefixRBTreeGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
//}

//#GOGP_ONCE
type ColorType int8

const (
	RBTREE_RED ColorType = iota //default
	RBTREE_BLACK
) //
//#GOGP_END_ONCE

type GOGPGlobalNamePrefixRBTree struct {
	header GOGPGlobalNamePrefixRBTreeNode
	size   int
	cmp    CmpGOGPGlobalNamePrefix
}

func (this *GOGPGlobalNamePrefixRBTree) root() *GOGPGlobalNamePrefixRBTreeNode {
	return this.header.parent
}
func (this *GOGPGlobalNamePrefixRBTree) topLeft() *GOGPGlobalNamePrefixRBTreeNode {
	return this.header.left
}
func (this *GOGPGlobalNamePrefixRBTree) topRight() *GOGPGlobalNamePrefixRBTreeNode {
	return this.header.right
}

type GOGPGlobalNamePrefixRBTreeNodeData struct {
	key GOGPKeyType
	//#GOGP_IFDEF VALUE_TYPE
	val GOGPValueType
	//#GOGP_ENDIF
}

func (this *GOGPGlobalNamePrefixRBTreeNodeData) Key() GOGPKeyType {
	return this.key
}

//#GOGP_IFDEF VALUE_TYPE
func (this *GOGPGlobalNamePrefixRBTreeNodeData) Value() GOGPValueType {
	return this.val
} //
//#GOGP_ENDIF

//tree node
type GOGPGlobalNamePrefixRBTreeNode struct {
	val                 GOGPGlobalNamePrefixRBTreeNodeData
	left, right, parent *GOGPGlobalNamePrefixRBTreeNode
	color               ColorType
}

//func (this *GOGPGlobalNamePrefixRBTree) Root() *GOGPGlobalNamePrefixRBTreeNode { return this.root() }

func (this *GOGPGlobalNamePrefixRBTreeNode) Get() *GOGPGlobalNamePrefixRBTreeNodeData {
	return &this.val
}

type GOGPGlobalNamePrefixRBTreeNodeVisitor struct {
	node, root *GOGPGlobalNamePrefixRBTreeNode
}

func (this *GOGPGlobalNamePrefixRBTreeNodeVisitor) Next() bool {
	return false
}

func (this *GOGPGlobalNamePrefixRBTreeNodeVisitor) Prev() bool {
	return false
}

func (this *GOGPGlobalNamePrefixRBTreeNodeVisitor) Get() *GOGPGlobalNamePrefixRBTreeNode {
	return this.node
}

func (this *GOGPGlobalNamePrefixRBTreeNode) rebalence(root **GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil && *root == this {
	}
}

func (this *GOGPGlobalNamePrefixRBTreeNode) topLeft() (n *GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil {
		for n = this; n.left != nil; n = n.left { //body do nothing
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixRBTreeNode) topRight() (n *GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil {
		for n = this; n.right != nil; n = n.right { //body do nothing
		}
	}
	return
}

//next node
func (this *GOGPGlobalNamePrefixRBTreeNode) next() (n *GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil {
		if this.right != nil {
			n = this.right.topLeft()
		} else {
			x, y := this, this.parent
			for x == y.right {
				x, y = y, y.parent
			}
			if x.right != y { //x is not header
				x = y
			}
			n = x
		}
	}
	return
}

//prev node
func (this *GOGPGlobalNamePrefixRBTreeNode) prev() (n *GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil {
		if this.color == RBTREE_RED && this.parent == this { //this is header
			n = this.right
		} else if this.left != nil {
			n = this.left.topRight()
		} else {
			x, y := this, this.parent
			for x == y.left {
				x, y = y, y.parent
			}
			n = y
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixRBTreeNode) rotateLeft(root **GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil && *root == this {
		y := this.right
		if this.right = y.left; y.left != nil {
			y.left.parent = this
		}
		y.parent = this.parent
		*root = y
		y.left = this
		this.parent = y
	}
}

func (this *GOGPGlobalNamePrefixRBTreeNode) rotateRight(root **GOGPGlobalNamePrefixRBTreeNode) {
	if this != nil && *root == this {
		y := this.left
		if this.left = y.right; y.right != nil {
			y.right.parent = this
		}
		y.parent = this.parent
		*root = y
		y.right = this
		this.parent = y
	}
}

//new object
func NewGOGPGlobalNamePrefixRBTree() *GOGPGlobalNamePrefixRBTree {
	return &GOGPGlobalNamePrefixRBTree{}
}

func (this *GOGPGlobalNamePrefixRBTree) Init(bigFirst bool) {
	this.header.parent = nil
	this.header.left = nil
	this.header.right = nil
	this.header.color = RBTREE_RED
	this.size = 0
	this.cmp.CreateByBool(bigFirst)
}

func (this *GOGPGlobalNamePrefixRBTree) insertUnique(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	x, y, comp := this.root(), &this.header, true
	for x != nil {
		y = x
		comp = this.cmp.F(d.key, x.val.key)
		if comp {
			x = x.left
		} else {
			x = x.right
		}
	}
	//todo
	return this.insert(x, y, d)
}

func (this *GOGPGlobalNamePrefixRBTree) insertEqual(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	x, y := this.root(), &this.header
	for x != nil {
		y = x
		if this.cmp.F(d.key, x.val.key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	return this.insert(x, y, d)
}

func (this *GOGPGlobalNamePrefixRBTree) insert(x, y *GOGPGlobalNamePrefixRBTreeNode, d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) Size() int {
	return this.size
}

func (this *GOGPGlobalNamePrefixRBTree) Empty() bool {
	return this.root() == nil
}

func (this *GOGPGlobalNamePrefixRBTree) Visitor(node *GOGPGlobalNamePrefixRBTreeNode) *GOGPGlobalNamePrefixRBTreeNodeVisitor {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) Begin() *GOGPGlobalNamePrefixRBTreeNode {
	return this.topLeft()
}

func (this *GOGPGlobalNamePrefixRBTree) End() *GOGPGlobalNamePrefixRBTreeNode {
	return &this.header
}

func (this *GOGPGlobalNamePrefixRBTree) Clear() {
	this.Init(this.cmp == CMPGreater)
}

func (this *GOGPGlobalNamePrefixRBTree) Insert(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) Remove(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) Erase(n *GOGPGlobalNamePrefixRBTreeNode) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) LowerBound(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) UpperBound(d GOGPGlobalNamePrefixRBTreeNodeData) *GOGPGlobalNamePrefixRBTreeNode {
	return nil
}

func (this *GOGPGlobalNamePrefixRBTree) Find(key GOGPKeyType) (r *GOGPGlobalNamePrefixRBTreeNode) {
	var y *GOGPGlobalNamePrefixRBTreeNode = nil
	for n := this.root(); n != nil; {
		if !this.cmp.F(n.val.key, key) {
			y, n = n, n.left
		} else {
			n = n.right
		}
	}
	if y != nil && !this.cmp.F(key, y.val.key) {
		r = y
	}
	return
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
