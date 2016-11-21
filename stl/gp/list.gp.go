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





//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
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
func (me CmpGOGPGlobalNamePrefix) F(left, right GOGPValueType) (ok bool) {
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
func (me CmpGOGPGlobalNamePrefix) less(left, right GOGPValueType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePrefix) great(left, right GOGPValueType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

var gGOGPGlobalNamePrefixListGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixListGbl.cmp = gGOGPGlobalNamePrefixListGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//double-way cycle link list node
type GOGPGlobalNamePrefixListNode struct {
	GOGPValueType
	prev, next *GOGPGlobalNamePrefixListNode
}

func (this *GOGPGlobalNamePrefixListNode) Get() GOGPValueType {
	return this.GOGPValueType
}

func (this *GOGPGlobalNamePrefixListNode) Set(v GOGPValueType) (old GOGPValueType) {
	old, this.GOGPValueType = this.GOGPValueType, v
	return
}

func (this *GOGPGlobalNamePrefixListNode) Next() (r *GOGPGlobalNamePrefixListNode) {
	if this != nil {
		r = this.next
	}
	return
}

func (this *GOGPGlobalNamePrefixListNode) Prev() (r *GOGPGlobalNamePrefixListNode) {
	if this != nil {
		r = this.prev
	}
	return
}

type GOGGlobalNamePrefixListNodeVisitor struct {
	node, head *GOGPGlobalNamePrefixListNode
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Reset() {
	this.node = nil
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Next() (ok bool) {
	if this.node == nil {
		if ok = this.head != nil; ok {
			this.node = this.head
		}
	} else {
		this.node = this.node.next
		ok = this.node != this.head
	}
	return
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Prev() (ok bool) {
	if this.node == nil {
		if ok = this.head != nil; ok {
			this.node = this.head.prev
		}
	} else {
		this.node = this.node.prev
		ok = this.node != this.head.prev
	}
	return
}

func (this *GOGGlobalNamePrefixListNodeVisitor) Get() *GOGPGlobalNamePrefixListNode {
	return this.node
}

func (this *GOGPGlobalNamePrefixList) Visitor(node *GOGPGlobalNamePrefixListNode) *GOGGlobalNamePrefixListNodeVisitor {
	n := &GOGGlobalNamePrefixListNodeVisitor{node: nil, head: this.head}
	return n
}

//list object
type GOGPGlobalNamePrefixList struct {
	head *GOGPGlobalNamePrefixListNode
}

//new object
func NewGOGPGlobalNamePrefixList() *GOGPGlobalNamePrefixList {
	return &GOGPGlobalNamePrefixList{}
}

//func (this *GOGPGlobalNamePrefixList) Len() int {
//	return 0
//}

func (this *GOGPGlobalNamePrefixList) Front() *GOGPGlobalNamePrefixListNode {
	return this.head
}
func (this *GOGPGlobalNamePrefixList) Back() (r *GOGPGlobalNamePrefixListNode) {
	if this.head != nil {
		r = this.head.prev
	}
	return
}

func (this *GOGPGlobalNamePrefixList) Clear() {
	this.head = nil
}

func (this *GOGPGlobalNamePrefixList) RotateForward() {
	if this.head != nil {
		this.head = this.head.next
	}
}
func (this *GOGPGlobalNamePrefixList) RotateBackward() {
	if this.head != nil {
		this.head = this.head.prev
	}
}

func (this *GOGPGlobalNamePrefixList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixListNode {
	n := this.PushBack(v)
	this.RotateBackward()
	return n
}

func (this *GOGPGlobalNamePrefixList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixListNode {
	n := &GOGPGlobalNamePrefixListNode{GOGPValueType: v}
	if this.head != nil {
		this.head, n.next, n.prev = n, n, n
	} else {
		n.next = this.head
		n.prev = this.head.prev
	}
	return n
}

func (this *GOGPGlobalNamePrefixList) PopFront() (v GOGPValueType, ok bool) {
	if n := this.Remove(this.Front()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *GOGPGlobalNamePrefixList) PopBack() (v GOGPValueType, ok bool) {
	if n := this.Remove(this.Back()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

func (this *GOGPGlobalNamePrefixList) PushFrontList(other *GOGPGlobalNamePrefixList) {
	this.PushBackList(other)
	this.RotateBackward()
}

func (this *GOGPGlobalNamePrefixList) PushBackList(other *GOGPGlobalNamePrefixList) {
	if other.head != nil {
		if this.head == nil {
			this.head = other.head
		} else {
			myback, oback := this.Back(), other.Back()
			myback.next = other.head
			oback.next = this.head
			other.head.prev = myback
			this.head.prev = oback
		}
	}
}

func (this *GOGPGlobalNamePrefixList) PushBefore(v GOGPValueType, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixListNode{GOGPValueType: v}
		n.next = mark
		n.prev = mark.prev
		mark.prev = n
	}
	return
}

func (this *GOGPGlobalNamePrefixList) PushAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixListNode{GOGPValueType: v}
		n.next = mark.next
		n.prev = mark
		mark.next = n
	}
	return
}

func (this *GOGPGlobalNamePrefixList) Remove(node *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if node != nil && node.next != nil && node.prev != nil {
		n = node
		if node.next == node {
			this.head = nil
		} else if node == this.head {
			this.head = node.next
		}
		node.next.prev = node.prev
		node.prev.next = node.next
		node.next, node.prev = nil, nil
	}
	return
}

func (this *GOGPGlobalNamePrefixList) Reachable(node, dest *GOGPGlobalNamePrefixListNode) (ok bool) {
	if ok = (node == dest) && node != nil; !ok && node != nil && dest != nil {
		for p := node; p != nil && p != node; p = p.next {
			if ok = (p == dest); ok {
				break
			}
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixList) IsValidNode(node *GOGPGlobalNamePrefixListNode) bool {
	return this.Reachable(this.head, node)
}

func (this *GOGPGlobalNamePrefixList) MoveFront(node *GOGPGlobalNamePrefixListNode) (r *GOGPGlobalNamePrefixListNode) {
	if r = this.MoveBack(node); r != nil {
		this.RotateBackward()
	}
	return
}

func (this *GOGPGlobalNamePrefixList) MoveBack(node *GOGPGlobalNamePrefixListNode) (r *GOGPGlobalNamePrefixListNode) {
	if node != nil && node.next != nil && node.prev != nil && node.next != node { //bug:node is back?
		node.next.prev = node.prev
		node.prev.next = node.next
		node.next = this.head
		node.prev = this.head.prev
		this.head.prev.next = node
		this.head.prev = node
		r = node
	}
	return
}

func (this *GOGPGlobalNamePrefixList) MoveBefore(node, mark *GOGPGlobalNamePrefixListNode) (r *GOGPGlobalNamePrefixListNode) {
	if node != nil && mark != nil && node != mark && node.next != mark {
		if r = this.Remove(node); r != nil {
			node.next = mark.prev.next
			node.prev = mark.prev
			mark.prev.next = node
			mark.prev = node
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixList) MoveAfter(node, mark *GOGPGlobalNamePrefixListNode) (r *GOGPGlobalNamePrefixListNode) {
	if node != nil && mark != nil && node != mark && mark.next != node {
		if r = this.Remove(node); r != nil {
			node.next = mark.next
			node.prev = mark
			mark.next.prev = node
			mark.next = node
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixList) Reverse() {}

func (this *GOGPGlobalNamePrefixList) Sort() {
	return
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
