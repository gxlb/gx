package gp

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
/*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is exactly not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

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

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var gGOGPGlobalNamePrefixDListGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixDListGbl.cmp = gGOGPGlobalNamePrefixDListGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//double-way cycle link list node
type GOGPGlobalNamePrefixDListNode struct {
	GOGPValueType
	prev, next *GOGPGlobalNamePrefixDListNode
}

func (this *GOGPGlobalNamePrefixDListNode) Get() *GOGPValueType {
	return &this.GOGPValueType
}

func (this *GOGPGlobalNamePrefixDListNode) Set(v GOGPValueType) (old GOGPValueType) {
	old, this.GOGPValueType = this.GOGPValueType, v
	return
}

func (this *GOGPGlobalNamePrefixDListNode) Next() (r *GOGPGlobalNamePrefixDListNode) {
	if this != nil {
		r = this.next
	}
	return
}

func (this *GOGPGlobalNamePrefixDListNode) Prev() (r *GOGPGlobalNamePrefixDListNode) {
	if this != nil {
		r = this.prev
	}
	return
}

//type GOGPListNamePrefixDListNodeVisitor struct {
//	node, head *GOGPListNamePrefixDListNode
//}

//func (this *GOGPListNamePrefixDListNodeVisitor) Next() bool {
//	return false
//}
//func (this *GOGPListNamePrefixDListNodeVisitor) Prev() bool {
//	return false
//}
//func (this *GOGPListNamePrefixDListNodeVisitor) Node() *GOGPListNamePrefixDListNode {
//	return nil
//}

//func (this *GOGPListNamePrefixDList) Visitor(node *GOGPListNamePrefixDListNode) *GOGPListNamePrefixDListNodeVisitor {
//	return nil
//}

//list object
type GOGPGlobalNamePrefixDList struct {
	head *GOGPGlobalNamePrefixDListNode
}

//new object
func NewGOGPGlobalNamePrefixDList() *GOGPGlobalNamePrefixDList {
	return &GOGPGlobalNamePrefixDList{}
}

//func (this *GOGPGlobalNamePrefixDList) Len() int {
//	return 0
//}

func (this *GOGPGlobalNamePrefixDList) Front() *GOGPGlobalNamePrefixDListNode {
	return this.head
}
func (this *GOGPGlobalNamePrefixDList) Back() (r *GOGPGlobalNamePrefixDListNode) {
	if this.head != nil {
		r = this.head.prev
	}
	return
}
func (this *GOGPGlobalNamePrefixDList) Clear() {
	this.head = nil
}
func (this *GOGPGlobalNamePrefixDList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixDListNode {
	n := this.PushBack(v)
	this.head = n
	return n
}
func (this *GOGPGlobalNamePrefixDList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixDListNode {
	n := &GOGPGlobalNamePrefixDListNode{GOGPValueType: v}
	if this.head != nil {
		this.head = n
		n.next = n
		n.prev = n
	} else {
		n.next = this.head
		n.prev = this.head.prev
	}
	return n
}
func (this *GOGPGlobalNamePrefixDList) PushFrontList(other *GOGPGlobalNamePrefixDList) {
	return
}
func (this *GOGPGlobalNamePrefixDList) PushBackList(other *GOGPGlobalNamePrefixDList) {
	return
}

func (this *GOGPGlobalNamePrefixDList) InsertBefore(v GOGPValueType, mark *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) InsertAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) Remove(node *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) MoveFront(node *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) MoveBck(node *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) MoveBefore(node, mark *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) MoveAfter(node, mark *GOGPGlobalNamePrefixDListNode) *GOGPGlobalNamePrefixDListNode {
	return nil
}
func (this *GOGPGlobalNamePrefixDList) Sort() {
	return
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
