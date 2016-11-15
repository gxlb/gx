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

func (this *GOGPGlobalNamePrefixSListNode) Get() GOGPValueType {
	return this.GOGPValueType
}

func (this *GOGPGlobalNamePrefixSListNode) Set(v GOGPValueType) (old GOGPValueType) {
	old, this.GOGPValueType = this.GOGPValueType, v
	return
}

func (this *GOGPGlobalNamePrefixSListNode) Next() (r *GOGPGlobalNamePrefixSListNode) {
	if this != nil {
		r = this.next
	}
	return
}

//type GOGPTreeNamePrefixRBTreeNodeVisitor struct {
//	node, head *GOGPGlobalNamePrefixSListNode
//}

//func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Next() bool {
//	return false
//}

//func (this *GOGPTreeNamePrefixRBTreeNodeVisitor) Node() *GOGPGlobalNamePrefixSListNode {
//	return nil
//}

//list object
type GOGPListNamePrefixSList struct {
	head *GOGPGlobalNamePrefixSListNode
	//#GOGP_IFDEF GOGP_HasTail
	tail *GOGPGlobalNamePrefixSListNode //
	//#GOGP_ENDIF

	//GOGP_Cycle
	//GOGP_HasTail
	//GOGP_DummyHead
}

//new object
func NewGOGPStackNamePrefixSList() *GOGPListNamePrefixSList {
	return &GOGPListNamePrefixSList{}
}

func (this *GOGPListNamePrefixSList) Len() int {
	return 0
}

//func (this *GOGPListNamePrefixSList) Visitor(node *GOGPGlobalNamePrefixSListNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
//	return nil
//}

func (this *GOGPListNamePrefixSList) Front() *GOGPGlobalNamePrefixSListNode {
	return this.head
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) Back() *GOGPGlobalNamePrefixSListNode {
	return this.tail
} //
//#GOGP_ENDIF

func (this *GOGPListNamePrefixSList) Clear() {
	this.head = nil
	//#GOGP_IFDEF GOGP_HasTail
	this.tail = nil //
	//#GOGP_ENDIF
}
func (this *GOGPListNamePrefixSList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v}
	n.next = this.head
	this.head = n
	//#GOGP_IFDEF GOGP_HasTail
	if this.tail == nil {
		this.tail = n
	} //
	//#GOGP_ENDIF
	return n
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v}
	if this.tail != nil {
		this.tail.next = n
	} else {
		this.head = n
	}
	this.tail = n
	return n
} //
//#GOGP_ENDIF

func (this *GOGPListNamePrefixSList) PushFrontList(other *GOGPListNamePrefixSList) {
	var t *GOGPGlobalNamePrefixSListNode
	//#GOGP_IFDEF GOGP_HasTail
	t = other.tail
	//#GOGP_ELSE
	for t = other.head; t != nil && t.next != nil; t = t.next {
	}
	//#GOGP_ENDIF
	if t != nil {
		t.next = this.head
		this.head = other.head
		//#GOGP_IFDEF GOGP_HasTail
		if this.tail == nil {
			this.tail = t
		} //
		//#GOGP_ENDIF
	}
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) PushBackList(other *GOGPListNamePrefixSList) {
	if other.tail != nil {
		if this.tail != nil {
			this.tail.next = other.head
			this.tail = other.tail
		} else {
			this.head, this.tail = other.head, other.tail
		}
		other.Clear()
	}
} //
//#GOGP_ENDIF

//func (this *GOGPListNamePrefixSList) InsertBefore(v GOGPValueType, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	return nil
//}

func (this *GOGPListNamePrefixSList) InsertAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixSListNode) (n *GOGPGlobalNamePrefixSListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixSListNode{GOGPValueType: v}
		n.next = mark.next
		mark.next = n
	}
	return
}
func (this *GOGPListNamePrefixSList) Remove(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil {
		for b := this.head; b != nil; b = b.next {
			if b.next == node {
				b.next = node.next
				r = node
				break
			}
		}
	}
	return
}
func (this *GOGPListNamePrefixSList) MoveFront(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil {
		for b := this.head; b != nil; b = b.next {
			if b.next == node {
				b.next = node.next
				node.next = this.head
				this.head = node
				r = node
				break
			} else if b == node {
				r = node
				break
			}
		}
	}
	return
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) MoveBack(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil {
		for b := this.head; b != nil; b = b.next {
			if b.next == node { //bug:node==head?
				b.next = node.next
				node.next = nil
				this.tail.next = node
				this.tail = node
				r = node
				break
			}
		}
	}
	return
} //
//#GOGP_ENDIF

//func (this *GOGPListNamePrefixSList) MoveBefore(node, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	return nil
//}

func (this *GOGPListNamePrefixSList) MoveAfter(node, mark *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	return nil
}
func (this *GOGPListNamePrefixSList) Sort() {
	return
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
