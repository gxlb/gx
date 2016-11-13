//single-way link list

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

var gGOGPGlobalNamePrefixSListGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixSListGbl.cmp = gGOGPGlobalNamePrefixSListGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//list node
type GOGPGlobalNamePrefixSListNode struct {
	GOGPValueType
	next *GOGPGlobalNamePrefixSListNode
}

func (this *GOGPGlobalNamePrefixSListNode) Get() *GOGPValueType {
	return &this.GOGPValueType
}

type GOGPTreeNamePrefixRBTreeNodeVisitor struct {
	node, head *GOGPGlobalNamePrefixSListNode
}

func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Next() bool {
	return false
}

func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Node() *GOGPGlobalNamePrefixSListNode {
	return nil
}

//list object
type GOGPListNamePrefixSList struct {
	head, tail *GOGPGlobalNamePrefixSListNode
}

//new object
func NewGOGPStackNamePrefixSList() *GOGPListNamePrefixSList {
	return &GOGPListNamePrefixSList{}
}

func (this *GOGPListNamePrefixSList) Len() int {
	return 0
}
func (this *GOGPListNamePrefixSList) Visitor(node *GOGPGlobalNamePrefixSListNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
	return nil
}
func (this *GOGPListNamePrefixSList) Front() *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Back() *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Clear() *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) PushFrontList(other *GOGPListNamePrefixSList) {
	return
}
func (this *GOGPListNamePrefixSList) PushBackList(other *GOGPListNamePrefixSList) {
	return
}

func (this *GOGPListNamePrefixSList) InsertBefore(v GOGPValueType, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) InsertAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Remove(node *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveFront(node *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveBck(node *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveBefore(node, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) MoveAfter(node, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Sort() {
	return
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
