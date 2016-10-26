package unsafe_test

import (
	"fmt"
	"testing"

	"github.com/vipally/gx/unsafe"
)

func TestString(t *testing.T) {
	s := ""
	for i := 10; i < 100; i++ {
		p := new(complex128)
		fmt.Println(i, p)
	}

	var bb []byte
	fmt.Println(bb)
	fmt.Println(unsafe.BytesPointer(bb))
	fmt.Println(unsafe.StringBytes(s))
	s = "hello"

	b := []byte{117, 118}
	s2 := fmt.Sprintf("hhh")
	bs := unsafe.StringBytes(s2)
	fmt.Println(unsafe.StringPointer(s))
	fmt.Println(unsafe.StringPointer(s2))
	fmt.Println(unsafe.BytesPointer(b))
	bs[0] = 'a'
	fmt.Println(string(bs))
}
