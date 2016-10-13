package stl

import (
	"fmt"
	"testing"
	"time"
	//"regexp"

	//_ "github.com/vipally/gogp/auto"
)

func init() {
	//	s := []GOGPSliceElem{6, 7, 8, 3, 4, 6}
	//	//MakeSlice(s).Sort()
	//	ss := GOGPTreeNamePrefixSortSlice(s)
	//	ss.Sort()
	//	ss.Pop()
	//	for _, v := range ss.D() {
	//		fmt.Println(v)
	//	}
}

func Fibonacci(n uint64) uint64 {
	start := time.Now()

	x := n
	if n > 1 {
		x = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cost := time.Now().Sub(start)
	fmt.Println(n, x, cost)
	return x

}

func TestA1(t *testing.T) {
	n := uint64(30)

	Fibonacci(n)

}
