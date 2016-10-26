package unsafe

import (
	"reflect"
	"unsafe"
)

type StringBytesT []byte
type UnsafeString string

//how to do this?
func (this *StringBytesT) Writeable() bool {
	return false
}

//return GoString's buffer slice(enable modify string)
func StringBytes(s string) StringBytesT {
	return *(*StringBytesT)(unsafe.Pointer(&s))
}

// convert b to string without copy
func BytesString(b []byte) UnsafeString {
	return *(*UnsafeString)(unsafe.Pointer(&b))
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
