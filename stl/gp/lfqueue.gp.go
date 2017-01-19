package gp

//#GOGP_FILE_BEGIN 1

import (
	"sync/atomic"
	"unsafe"
)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)

//list node
type GOGPGlobalNamePrefixLFQueueNode struct {
	val  GOGPValueType
	next unsafe.Pointer
}

//single-way link list object
type GOGPGlobalNamePrefixLFQueue struct {
	head GOGPGlobalNamePrefixLFQueueNode //head is a dummy node, not a pionter
	tail unsafe.Pointer
	free unsafe.Pointer
	size int32
}

func (this *GOGPGlobalNamePrefixLFQueue) makeNode(v GOGPValueType) (n *GOGPGlobalNamePrefixLFQueueNode) {
	for {
		p := atomic.LoadPointer(&this.free)
		if p == nil {
			n = &GOGPGlobalNamePrefixLFQueueNode{}
			break
		}

		n = (*GOGPGlobalNamePrefixLFQueueNode)(p)
		next := n.next
		if atomic.CompareAndSwapPointer(&this.free, p, next) {
			break
		}
	}
	n.val = v
	n.next = nil
	return
}

func (this *GOGPGlobalNamePrefixLFQueue) PushFront(v GOGPValueType) bool {
	n := this.makeNode(v)
	p := unsafe.Pointer(n)
	if atomic.LoadPointer(&this.tail) == nil {
		atomic.StorePointer(&this.tail, p)
	}
	for {
		n.next = atomic.LoadPointer(&this.head.next)
		if atomic.CompareAndSwapPointer(&this.head.next, n.next, p) {
			break
		}
	}
	atomic.AddInt32(&this.size, 1)
	return true
}

func (this *GOGPGlobalNamePrefixLFQueue) PushBack(v GOGPValueType) bool {
	//	n := this.makeNode(v)
	//	for {
	//		t := atomic.LoadPointer(&this.tail)
	//		if t == nil {
	//		}
	//	}
	return true
}

func (this *GOGPGlobalNamePrefixLFQueue) PopFront() (v GOGPValueType, ok bool) {
	return
}

func (this *GOGPGlobalNamePrefixLFQueue) ClearFree() {

}

//func (this *GOGPGlobalNamePrefixLFQueue) PopBack() (v GOGPValueType, ok bool)  { return }

func (this *GOGPGlobalNamePrefixLFQueue) Size() int {
	return int(atomic.LoadInt32(&this.size))
}

func (this *GOGPGlobalNamePrefixLFQueue) Empty() bool {
	return atomic.LoadPointer(&this.head.next) == nil
}

//#GOGP_FILE_END
