package gp

//#GOGP_FILE_BEGIN 1

//import (
//	"sync/atomic"
//)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)

//list node
type GOGPGlobalNamePrefixLFQueueNode struct {
	val  GOGPValueType
	next *GOGPGlobalNamePrefixLFQueueNode
}

//single-way link list object
type GOGPGlobalNamePrefixLFQueue struct {
	head       GOGPGlobalNamePrefixLFQueueNode //head is a dummy node, not a pionter
	tail, free *GOGPGlobalNamePrefixLFQueueNode
}

func (this *GOGPGlobalNamePrefixLFQueue) PushFront(v GOGPValueType) bool { return false }

//#GOGP_FILE_END
