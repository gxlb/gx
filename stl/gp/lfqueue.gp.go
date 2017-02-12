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

import (
	"sync/atomic"
	"unsafe"
)

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

//list node
type GOGPGlobalNamePrefixLFQueueNode struct {
	val  GOGPValueType
	next unsafe.Pointer
}

//single-way link list object
type GOGPGlobalNamePrefixLFQueue struct {
	head GOGPGlobalNamePrefixLFQueueNode //head is a dummy node, not a pionter
	//tail unsafe.Pointer
	size int32
}

func (this *GOGPGlobalNamePrefixLFQueue) makeNode(v GOGPValueType) (n *GOGPGlobalNamePrefixLFQueueNode) {
	n = &GOGPGlobalNamePrefixLFQueueNode{}
	n.val = v
	n.next = nil
	return
}

func (this *GOGPGlobalNamePrefixLFQueue) PushFront(v GOGPValueType) bool {
	n := this.makeNode(v)
	p := unsafe.Pointer(n)
	for {
		n.next = atomic.LoadPointer(&this.head.next)
		if atomic.CompareAndSwapPointer(&this.head.next, n.next, p) {
			break
		}
	}
	//	if atomic.LoadPointer(&this.tail) == nil {
	//		atomic.StorePointer(&this.tail, p)
	//	}
	atomic.AddInt32(&this.size, 1)
	return true
}

//func (this *GOGPGlobalNamePrefixLFQueue) PushBack(v GOGPValueType) bool {
//	n := this.makeNode(v)
//	p := unsafe.Pointer(n)
//	for {
//		t := atomic.LoadPointer(&this.tail)
//		if t == nil {
//			atomic.CompareAndSwapPointer(&this.tail, nil, p)
//			if atomic.CompareAndSwapPointer(&this.head.next, nil, p) {
//				break
//			}
//		} else {
//			tt := (*GOGPGlobalNamePrefixLFQueueNode)(t)
//			if atomic.CompareAndSwapPointer(&this.tail, t, p) {
//				tt.next = p
//				break
//			}
//		}
//	}
//	atomic.AddInt32(&this.size, 1)
//	return true
//}

func (this *GOGPGlobalNamePrefixLFQueue) PopFront() (v GOGPValueType, ok bool) {
	for {
		t := atomic.LoadPointer(&this.head.next)
		if t == nil {
			break
		} else {
			n := (*GOGPGlobalNamePrefixLFQueueNode)(t)
			next := n.next
			if atomic.CompareAndSwapPointer(&this.head.next, t, next) {
				v, ok = n.val, true
				break
			}
		}
	}
	atomic.AddInt32(&this.size, -1)
	return
}

func (this *GOGPGlobalNamePrefixLFQueue) Clear() {
	atomic.StorePointer(&this.head.next, nil)
	//atomic.StorePointer(&this.tail, nil)
}

//func (this *GOGPGlobalNamePrefixLFQueue) PopBack() (v GOGPValueType, ok bool)  { return }

func (this *GOGPGlobalNamePrefixLFQueue) Size() int {
	return int(atomic.LoadInt32(&this.size))
}

func (this *GOGPGlobalNamePrefixLFQueue) Empty() bool {
	return atomic.LoadPointer(&this.head.next) == nil
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
