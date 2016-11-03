package unsafe

import (
	"reflect"
	"unsafe"
)

var (
	//dynamic data address
	firstPointer = uintptr(unsafe.Pointer(new(int)))
)

type Pointer unsafe.Pointer
type Bytes []byte
type String string

//invalid receiver type Pointer
//func (me Pointer) Writeable() bool {
//	addr := uintptr(me)
//	return addr >= firstPointer
//}

func Writeable(ptr Pointer) bool {
	addr := uintptr(ptr)
	return addr >= firstPointer
}

//how to do this?
func (me Bytes) Writeable() bool {
	addr := uintptr(BytesPointer(me))
	return addr >= firstPointer
}

func (me Bytes) Pointer() Pointer {
	return BytesPointer(me)
}

//check if a string's buffer is writeable
func (me String) Writeable() bool {
	p := uintptr(StringPointer(string(me)))
	return p >= firstPointer
}

func (me String) Pointer() Pointer {
	return StringPointer(string(me))
}

//return GoString's buffer slice(enable modify string)
func StringBytes(s string) Bytes {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*Bytes)(unsafe.Pointer(&bh))
}

// convert b to string without copy
func BytesString(b []byte) String {
	return *(*String)(unsafe.Pointer(&b))
}

// returns &s[0], which is not allowed in go
func StringPointer(s string) Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func BytesPointer(b []byte) Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return Pointer(p.Data)
}
