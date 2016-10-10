package convert

import (
	"fmt"
	//"reflect"
	"testing"
	"unsafe"
)

func TestConvert(t *testing.T) {
	type T struct {
		a int
		b int
		c int
	}

	type SliceHeader struct {
		addr uintptr
		len  int
		cap  int
	}

	tt := &T{a: 1, b: 2, c: 3}
	p := unsafe.Sizeof(*tt)
	fmt.Println(int(p))

	sl := &SliceHeader{
		addr: uintptr(unsafe.Pointer(tt)),
		len:  int(p),
		cap:  int(p),
	}

	b := *(*[]byte)(unsafe.Pointer(sl))
	fmt.Println(len(b))
	fmt.Println(b)

	b[0] = 7
	b[4] = 5
	b[8] = 8

	fmt.Println(tt)
	fmt.Println(b)
	//	type reflect.StringHeader struct {
	//		Data uintptr
	//		Len  int
	//	}
	//	type reflect.SliceHeader struct {
	//		Data uintptr
	//		Len  int
	//		Cap  int
	//	}
}
