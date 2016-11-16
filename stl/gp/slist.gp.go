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

func (this *GOGPGlobalNamePrefixSListNode) Tail() (r *GOGPGlobalNamePrefixSListNode, size int) {
	if this != nil {
		for r, size = this, 1; r != nil && r.next != nil; r, size = r.next, size+1 { //do nothing body
		}
	}
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

//single-way link list object
type GOGPListNamePrefixSList struct {
	head GOGPGlobalNamePrefixSListNode //head is a dummy node, not a pionter
	//#GOGP_IFDEF GOGP_HasTail
	tail *GOGPGlobalNamePrefixSListNode //
	//#GOGP_ENDIF

	//GOGP_HasTail
	//bug: ring?
}

//new object
func NewGOGPStackNamePrefixSList() *GOGPListNamePrefixSList {
	return &GOGPListNamePrefixSList{}
}

//func (this *GOGPListNamePrefixSList) Len() int {
//	return 0
//}

//func (this *GOGPListNamePrefixSList) Visitor(node *GOGPGlobalNamePrefixSListNode) *GOGPTreeNamePrefixRBTreeNodeVisitor {
//	return nil
//}

func (this *GOGPListNamePrefixSList) Front() *GOGPGlobalNamePrefixSListNode {
	return this.head.next
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) Back() *GOGPGlobalNamePrefixSListNode {
	return this.tail
} //
//#GOGP_ENDIF

func (this *GOGPListNamePrefixSList) Clear() {
	this.head.next = nil
	//#GOGP_IFDEF GOGP_HasTail
	this.tail = nil //
	//#GOGP_ENDIF
}

func (this *GOGPListNamePrefixSList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: this.head.next}
	this.head.next = n
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
		this.head.next = n
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
	t, _ = other.head.next.Tail()
	//#GOGP_ENDIF
	if t != nil {
		t.next = this.head.next
		this.head.next = other.head.next
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
			this.tail.next = other.head.next
			this.tail = other.tail
		} else {
			this.head.next, this.tail = other.head.next, other.tail
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
		n = &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: mark.next}
		mark.next = n
	}
	return
}

func (this *GOGPListNamePrefixSList) Remove(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil {
		for b := &this.head; b != nil; b = b.next {
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
		for b := this.head.next; b != nil; b = b.next {
			if b.next == node {
				b.next = node.next
				node.next = this.head.next
				this.head.next = node
				r = node
				//#GOGP_IFDEF GOGP_HasTail
				if this.tail == node {
					this.tail = b
				} //
				//#GOGP_ENDIF
				break
			}
		}
	}
	return
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPListNamePrefixSList) MoveBack(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil && node.next != nil {
		for b := this.head.next; b != nil; b = b.next {
			if b.next == node {
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

func (this *GOGPListNamePrefixSList) MoveAfter(node, mark *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil && mark != nil && node != mark {
		for b := &this.head; b != nil; b = b.next { //bug:tail?
			if b.next == node {
				if b != mark {
					b.next = node.next
					node.next = mark.next
					mark.next = node
					r = node
					//#GOGP_IFDEF GOGP_HasTail
					if this.tail == node {
						this.tail = b
					} //
					//#GOGP_ENDIF
				}
				break
			}
		}
	}
	return
}

//merge sort
func (this *GOGPListNamePrefixSList) Sort() {
	head := this.head.next
	this.Clear()
	this.head.next = this.insertSort(head)
	//#GOGP_IFDEF GOGP_HasTail
	this.tail, _ = this.head.next.Tail() //
	//#GOGP_ENDIF
	return
}

func (this *GOGPListNamePrefixSList) mid(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	if head == nil || head.next == nil {
		return head
	}
	slow := head
	for fast := head; fast != nil && fast.next != nil; fast, slow = fast.next.next, slow.next { //do nothing body
	}
	return slow
}

func (this *GOGPListNamePrefixSList) merge(a, b *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	var (
		head GOGPGlobalNamePrefixSListNode
		tail *GOGPGlobalNamePrefixSListNode
	)
	tail = &head
	for a != nil && b != nil {
		if gGOGPGlobalNamePrefixSListGbl.cmp.F(a.GOGPValueType, b.GOGPValueType) {
			tail.next = a
			tail = tail.next
			a = a.next
		} else {
			tail.next = b
			tail = tail.next
			b = b.next
		}
	}
	if a != nil {
		tail.next = a
		tail, _ = a.Tail()
	} else if b != nil {
		tail.next = b
		tail, _ = b.Tail()
	}
	tail.next = nil

	return head.next
}

func (this *GOGPListNamePrefixSList) mergeSort(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	if head == nil || head.next == nil {
		return head
	} else if mid := this.mid(head); mid != nil {
		midNext := mid.next
		mid.next = nil
		return this.merge(this.mergeSort(mid), this.mergeSort(midNext))
	}
	return nil
}

func (this *GOGPListNamePrefixSList) PushOrderly(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: nil}
	return this.InsertOrderly(n)
}

func (this *GOGPListNamePrefixSList) InsertOrderly(node *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	p := &this.head
	for ; p.next != nil; p = p.next {
		if n := p.next; gGOGPGlobalNamePrefixSListGbl.cmp.F(node.GOGPValueType, n.GOGPValueType) {
			break
		}
	}
	node.next = p.next
	p.next = node
	//#GOGP_IFDEF GOGP_HasTail
	if this.tail == nil || this.tail == p {
		this.tail = node
	} //
	//#GOGP_ENDIF
	return node
}

func (this *GOGPListNamePrefixSList) insertSort(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
	if head == nil || head.next == nil {
		return head
	}

	var h GOGPListNamePrefixSList
	for p := head; p != nil; {
		q := p.next
		p.next = nil
		h.InsertOrderly(p)
		p = q
	}
	return h.head.next
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
