package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	//"time"

	xreflect "github.com/vipally/gx/reflect"
)

type I int

func TestNew(t *testing.T) {
	f := 1
	n := xreflect.New(f)
	fmt.Println(&f, n, reflect.ValueOf(n))
	var (
		i  I    = 1
		j  I    = 2
		b  byte = 1
		u8 rune = 2
	)
	fmt.Println(i + j)
	fmt.Println(reflect.TypeOf(i).PkgPath())
	fmt.Println(xreflect.IsSameType(b, u8))
	//fmt.Println(xreflect.IsCmpableInnterType(f))
	fmt.Println(xreflect.IsSameType(uint8(0), byte(0)))
	fmt.Println(xreflect.IsSameType(rune(0), int32(0)))
}

func Benchmark_TypeCheck(b *testing.B) {
	b.N = 2100000000
	for ii := 0; ii < 100000000; ii++ {
		f := 1
		var i I = 1
		xreflect.IsSameType(f, i)
	}
}
