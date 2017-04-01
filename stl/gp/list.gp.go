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

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END ///require end from(github.com/vipally/gogp/lib/fakedef)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/vipally/gx/stl/gp/functorcmp)
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

//#GOGP_IGNORE_END ///require end from(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

var gGOGPGlobalNamePrefixListGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixListGbl.cmp = gGOGPGlobalNamePrefixListGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//double-way cycle link list node
type GOGPGlobalNamePrefixListNode struct {
	val        GOGPValueType
	prev, next *GOGPGlobalNamePrefixListNode
}

func (this *GOGPGlobalNamePrefixListNode) Get() GOGPValueType {
	return this.val
}

func (this *GOGPGlobalNamePrefixListNode) Set(v GOGPValueType) (old GOGPValueType) {
	old, this.val = this.val, v
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

type GOGPGlobalNamePrefixListNodeVisitor struct {
	node, head *GOGPGlobalNamePrefixListNode
}

func (this *GOGPGlobalNamePrefixListNodeVisitor) Reset() {
	this.node = nil
}

func (this *GOGPGlobalNamePrefixListNodeVisitor) Next() (ok bool) {
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

func (this *GOGPGlobalNamePrefixListNodeVisitor) Prev() (ok bool) {
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

func (this *GOGPGlobalNamePrefixListNodeVisitor) Get() *GOGPGlobalNamePrefixListNode {
	return this.node
}

func (this *GOGPGlobalNamePrefixList) Visitor() *GOGPGlobalNamePrefixListNodeVisitor {
	n := &GOGPGlobalNamePrefixListNodeVisitor{node: nil, head: this.head}
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
	n := &GOGPGlobalNamePrefixListNode{val: v}
	return this.InsertFront(n)
}

func (this *GOGPGlobalNamePrefixList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixListNode {
	n := &GOGPGlobalNamePrefixListNode{val: v}
	return this.InsertBack(n)
}

func (this *GOGPGlobalNamePrefixList) InsertFront(node *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if n = this.InsertBack(node); n != nil {
		this.RotateBackward()
	}
	return
}

func (this *GOGPGlobalNamePrefixList) InsertBack(node *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if n = node; n != nil {
		if this.head == nil {
			this.head, n.next, n.prev = n, n, n
		} else {
			n.next = this.head
			n.prev = this.head.prev
			this.head.prev.next = n
			this.head.prev = n
		}
	}
	return
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

func (this *GOGPGlobalNamePrefixList) InsertFrontList(other *GOGPGlobalNamePrefixList) (ok bool) {
	rotate := !this.Empty()
	if ok = this.InsertBackList(other); ok && rotate {
		this.RotateBackward()
	}
	return
}

func (this *GOGPGlobalNamePrefixList) InsertBackList(other *GOGPGlobalNamePrefixList) (ok bool) {
	if ok = !other.Empty(); ok {
		if this.Empty() {
			this.head = other.head
		} else {
			myback, oback := this.Back(), other.Back()
			myback.next = other.head
			oback.next = this.head
			other.head.prev = myback
			this.head.prev = oback
		}
		other.Clear()
	}
	return
}

func (this *GOGPGlobalNamePrefixList) PushBefore(v GOGPValueType, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixListNode{val: v}
		n = this.InsertBefore(n, mark)
	}
	return
}

func (this *GOGPGlobalNamePrefixList) PushAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixListNode{val: v}
		n = this.InsertAfter(n, mark)
	}
	return
}

func (this *GOGPGlobalNamePrefixList) InsertBefore(node, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if n = node; node != nil && mark != nil && node != mark {
		n.next = mark
		n.prev = mark.prev
		mark.prev.next = n
		mark.prev = n
		if this.head == mark {
			this.RotateBackward()
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixList) InsertAfter(node, mark *GOGPGlobalNamePrefixListNode) (n *GOGPGlobalNamePrefixListNode) {
	if n = node; node != nil && mark != nil {
		n.next = mark.next
		n.prev = mark
		mark.next = n
	}
	return
}

func (this *GOGPGlobalNamePrefixList) RemoveFront() (n *GOGPGlobalNamePrefixListNode) {
	return this.Remove(this.Front())
}

func (this *GOGPGlobalNamePrefixList) RemoveBack() (n *GOGPGlobalNamePrefixListNode) {
	return this.Remove(this.Back())
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

func (this *GOGPGlobalNamePrefixList) Empty() bool {
	return this.head == nil
}

func (this *GOGPGlobalNamePrefixList) Reverse() {
	v := this.Visitor()
	this.Clear()
	for v.Next() {
		this.InsertFront(v.Get())
	}
}

func (this *GOGPGlobalNamePrefixList) Sort() {
	this.mergeSort()
}

//STL's merge sort algorithm for list
func (this *GOGPGlobalNamePrefixList) mergeSort() {
	if nil == this.head || this.head == this.head.next { //0 or 1 element, no need to sort
		return
	}

	var (
		hand       GOGPGlobalNamePrefixList
		binList    [64]GOGPGlobalNamePrefixList //save temp list that len=2^i
		nFilledBin = 0
	)

	for !this.Empty() {
		hand.InsertFront(this.RemoveFront())
		i := 0
		for ; i < nFilledBin && !binList[i].Empty(); i++ {
			binList[i].merge(&hand)
			hand, binList[i] = binList[i], hand
		}
		hand, binList[i] = binList[i], hand
		if i == nFilledBin {
			nFilledBin++
		}
	}

	for i := 1; i < nFilledBin; i++ {
		binList[i].merge(&binList[i-1])
	}

	*this = binList[nFilledBin-1]
}

//merge other sorted-list  to this sorted-list
func (this *GOGPGlobalNamePrefixList) merge(other *GOGPGlobalNamePrefixList) {
	if this.Empty() || other.Empty() {
		this.InsertBackList(other)
		return
	}

	p, po := this.Front(), other.Front()
	for p != nil && po != nil {
		if gGOGPGlobalNamePrefixListGbl.cmp.F(po.val, p.val) {
			n := other.RemoveFront()
			po = other.Front()
			this.InsertBefore(n, p)
		} else {
			if p = p.next; p == this.Front() {
				p = nil
			}
		}
	}
	if po != nil {
		this.InsertBackList(other)
	}
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
