package unsafe

import (
	"reflect"
	"unsafe"
)

var (
	//dynamic data address
	firstPointer = uintptr(unsafe.Pointer(new(int))) >> 7
)

type Bytes []byte
type String string

//how to do this?
func (me Bytes) Writeable() bool {
	addr := uintptr(BytesPointer(me))
	return addr < firstPointer
}

//check if a string's buffer is writeable
func (me String) Writeable() bool {
	p := uintptr(StringPointer(string(me)))
	return p < firstPointer
}

//return GoString's buffer slice(enable modify string)
func StringBytes(s string) Bytes {
	return *(*Bytes)(unsafe.Pointer(&s))
}

// convert b to string without copy
func BytesString(b []byte) String {
	return *(*String)(unsafe.Pointer(&b))
}

// returns &s[0], which is not allowed in go
func StringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func BytesPointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}
