package unsafe

import (
	//"reflect"
	"unsafe"
)

//return GoString's buffer slice(enable modify string)
func StringBufer(s string) []byte {
	//b := (*[]byte(*reflect.SliceHeader(*reflect.StringHeader(unsafe.Pointer(&s)))))
	b := (*[]byte)(unsafe.Pointer(&s))
	return *b
}
