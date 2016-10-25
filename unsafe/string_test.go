package unsafe_test

import (
	"fmt"
	"testing"

	"github.com/vipally/gx/unsafe"
)

func TestString(t *testing.T) {
	s := "hello"
	b := unsafe.StringBufer(s)
	b[0] = 'G'
	fmt.Println(s)
}
