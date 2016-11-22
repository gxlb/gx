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
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

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

type GOGPGlobalNamePrefixSListNodeVisitor struct {
	node, head *GOGPGlobalNamePrefixSListNode
}

func (this *GOGPGlobalNamePrefixSListNodeVisitor) Next() (ok bool) {
	if this.node == nil {
		if ok = this.head != nil; ok {
			this.node = this.head
		}
	} else {
		this.node = this.node.next
		ok = this.node != nil
	}
	return false
}

func (this *GOGPGlobalNamePrefixSListNodeVisitor) Get() *GOGPGlobalNamePrefixSListNode {
	return this.node
}

//single-way link list object
type GOGPGlobalNamePrefixSList struct {
	head GOGPGlobalNamePrefixSListNode //head is a dummy node, not a pionter
	//#GOGP_IFDEF GOGP_HasTail
	tail *GOGPGlobalNamePrefixSListNode //
	//#GOGP_ENDIF

	//GOGP_HasTail
	//bug: ring?
}

//new object
func NewGOGPGlobalNamePrefixSList() *GOGPGlobalNamePrefixSList {
	return &GOGPGlobalNamePrefixSList{}
}

func (this *GOGPGlobalNamePrefixSList) Empty() bool {
	return this.head.next != nil
}

//func (this *GOGPGlobalNamePrefixSList) Len() int {
//	return 0
//}

func (this *GOGPGlobalNamePrefixSList) Visitor() *GOGPGlobalNamePrefixSListNodeVisitor {
	v := &GOGPGlobalNamePrefixSListNodeVisitor{node: nil, head: this.head.next}
	return v
}

func (this *GOGPGlobalNamePrefixSList) Front() *GOGPGlobalNamePrefixSListNode {
	return this.head.next
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPGlobalNamePrefixSList) Back() *GOGPGlobalNamePrefixSListNode {
	return this.tail
} //
//#GOGP_ENDIF

func (this *GOGPGlobalNamePrefixSList) Clear() {
	this.head.next = nil
	//#GOGP_IFDEF GOGP_HasTail
	this.tail = nil //
	//#GOGP_ENDIF
}

func (this *GOGPGlobalNamePrefixSList) PushFront(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: this.head.next}
	this.head.next = n
	//#GOGP_IFDEF GOGP_HasTail
	if this.tail == nil {
		this.tail = n
	} //
	//#GOGP_ENDIF
	return n
}

func (this *GOGPGlobalNamePrefixSList) PopFront() (v GOGPValueType, ok bool) {
	if n := this.Remove(this.Front()); n != nil {
		v, ok = n.Get(), true
	}
	return
}

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPGlobalNamePrefixSList) PopBack() (v GOGPValueType, ok bool) {
	if n := this.Remove(this.Back()); n != nil {
		v, ok = n.Get(), true
	}
	return
} //
//#GOGP_ENDIF

//#GOGP_IFDEF GOGP_HasTail
func (this *GOGPGlobalNamePrefixSList) PushBack(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
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

func (this *GOGPGlobalNamePrefixSList) PushFrontList(other *GOGPGlobalNamePrefixSList) {
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
func (this *GOGPGlobalNamePrefixSList) PushBackList(other *GOGPGlobalNamePrefixSList) {
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

func (this *GOGPGlobalNamePrefixSList) PushAfter(v GOGPValueType, mark *GOGPGlobalNamePrefixSListNode) (n *GOGPGlobalNamePrefixSListNode) {
	if mark != nil {
		n = &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: mark.next}
		n = this.InsertAfter(n, mark)
	}
	return
}

func (this *GOGPGlobalNamePrefixSList) InsertAfter(node, mark *GOGPGlobalNamePrefixSListNode) (n *GOGPGlobalNamePrefixSListNode) {
	if mark != nil && node != nil && node != mark {
		node.next = mark.next
		mark.next = node
		n = node
	}
	return
}

func (this *GOGPGlobalNamePrefixSList) InsertFront(node *GOGPGlobalNamePrefixSListNode) (n *GOGPGlobalNamePrefixSListNode) {
	n = this.InsertAfter(node, &this.head)
	return
}

func (this *GOGPGlobalNamePrefixSList) RemoveFront() (n *GOGPGlobalNamePrefixSListNode) {
	n = this.Remove(this.head.next)
	return
}

func (this *GOGPGlobalNamePrefixSList) Remove(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
	if node != nil {
		for b := &this.head; b.next != nil; b = b.next {
			if b.next == node {
				b.next = node.next
				node.next = nil
				r = node
				break
			}
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixSList) MoveFront(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
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
func (this *GOGPGlobalNamePrefixSList) MoveBack(node *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
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

func (this *GOGPGlobalNamePrefixSList) MoveAfter(node, mark *GOGPGlobalNamePrefixSListNode) (r *GOGPGlobalNamePrefixSListNode) {
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

func (this *GOGPGlobalNamePrefixSList) Reverse() {
	p := this.head.next
	this.Clear()
	//#GOGP_IFDEF GOGP_HasTail
	this.tail = p //
	//#GOGP_ENDIF
	for p != nil {
		q := p.next
		this.InsertFront(p)
		p = q
	}
}

func (this *GOGPGlobalNamePrefixSList) HasCycle(head *GOGPGlobalNamePrefixSListNode) (ok bool) {
	for slow, fast := head, head; fast != nil; fast, slow = fast.next.next, slow.next { //do nothing body
		fast = fast.next
		if ok = fast != nil; !ok {
			break
		}
		fast = fast.next
		slow = slow.next
		if ok = fast == slow; ok {
			break
		}
	}
	return
}

//func (this *GOGPGlobalNamePrefixSList) FindCycleStart() *GOGPGlobalNamePrefixSListNode { return nil }

//func (this *GOGPGlobalNamePrefixDList) Find(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
//	return nil
//}

func (this *GOGPGlobalNamePrefixSList) Reachable(node, dest *GOGPGlobalNamePrefixSListNode) (ok bool) {
	if ok = (node == dest) && node != nil; !ok && node != nil && dest != nil {
		for p := node; p != nil && p != node; p = p.next {
			if ok = (p == dest); ok {
				break
			}
		}
	}
	return
}

func (this *GOGPGlobalNamePrefixSList) IsValidNode(node *GOGPGlobalNamePrefixSListNode) bool {
	return this.Reachable(this.head.next, node)
}

//merge sort
func (this *GOGPGlobalNamePrefixSList) Sort() {
	if nil == this.head.next || nil == this.head.next.next { //0 or 1 element, no need to sort
		return
	}
	this.mergeSort()
}

//STL's merge sort algorithm for list
func (this *GOGPGlobalNamePrefixSList) mergeSort() {
	var (
		hand, unsorted GOGPGlobalNamePrefixSList
		binList        [64]GOGPGlobalNamePrefixSList //save temp list that len=2^i
		nFilledBin     = 0
	)

	for unsorted = *this; !unsorted.Empty(); {
		hand.InsertFront(unsorted.RemoveFront())
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
	//#GOGP_IFDEF GOGP_HasTail
	this.tail, _ = this.head.next.Tail() //
	//#GOGP_ENDIF
}

//merge other sorted-list  to this sorted-list
func (this *GOGPGlobalNamePrefixSList) merge(other *GOGPGlobalNamePrefixSList) {
	p, po := &this.head, other.Front()
	for p.next != nil && po != nil {
		pn := p.next
		if gGOGPGlobalNamePrefixSListGbl.cmp.F(po.GOGPValueType, pn.GOGPValueType) {
			n := other.RemoveFront()
			po = other.Front()
			p = this.InsertAfter(n, p)
		} else {
			p = p.next
		}
	}
	if po != nil {
		p.next = po
		other.Clear()
		//#GOGP_IFDEF GOGP_HasTail
		this.tail = other.tail //
		//#GOGP_ENDIF
	}
}

//func (this *GOGPGlobalNamePrefixSList) merge(a, b *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	if a == nil {
//		return b
//	} else if b == nil {
//		return a
//	}
//	var (
//		head GOGPGlobalNamePrefixSListNode
//		tail *GOGPGlobalNamePrefixSListNode
//	)
//	for tail = &head; a != nil && b != nil; {
//		if gGOGPGlobalNamePrefixSListGbl.cmp.F(a.GOGPValueType, b.GOGPValueType) {
//			tail.next = a
//			tail = tail.next
//			a = a.next
//		} else {
//			tail.next = b
//			tail = tail.next
//			b = b.next
//		}
//	}
//	if a != nil {
//		tail.next = a
//		tail, _ = a.Tail()
//	} else if b != nil {
//		tail.next = b
//		tail, _ = b.Tail()
//	}
//	tail.next = nil

//	return head.next
//}

//func (this *GOGPGlobalNamePrefixSList) mid(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	if head == nil || head.next == nil {
//		return head
//	}
//	slow := head
//	for fast := head; fast != nil && fast.next != nil; fast, slow = fast.next.next, slow.next { //do nothing body
//	}
//	return slow
//}

//func (this *GOGPGlobalNamePrefixSList) mergeSort(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	if head == nil || head.next == nil {
//		return head
//	} else if mid := this.mid(head); mid != nil {
//		midNext := mid.next
//		mid.next = nil
//		return this.merge(this.mergeSort(mid), this.mergeSort(midNext))
//	}
//	return nil
//}

//func (this *GOGPGlobalNamePrefixSList) PushOrderly(v GOGPValueType) *GOGPGlobalNamePrefixSListNode {
//	n := &GOGPGlobalNamePrefixSListNode{GOGPValueType: v, next: nil}
//	return this.InsertOrderly(n)
//}

//func (this *GOGPGlobalNamePrefixSList) InsertOrderly(node *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	p := &this.head
//	for ; p.next != nil; p = p.next {
//		if n := p.next; gGOGPGlobalNamePrefixSListGbl.cmp.F(node.GOGPValueType, n.GOGPValueType) {
//			break
//		}
//	}
//	node.next = p.next
//	p.next = node
//	//#GOGP_IFDEF GOGP_HasTail
//	if this.tail == nil || this.tail == p {
//		this.tail = node
//	} //
//	//#GOGP_ENDIF
//	return node
//}

//func (this *GOGPGlobalNamePrefixSList) insertSort(head *GOGPGlobalNamePrefixSListNode) *GOGPGlobalNamePrefixSListNode {
//	if head == nil || head.next == nil {
//		return head
//	}

//	var h GOGPGlobalNamePrefixSList
//	for p := head; p != nil; {
//		q := p.next
//		p.next = nil
//		h.InsertOrderly(p)
//		p = q
//	}
//	return h.head.next
//}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
