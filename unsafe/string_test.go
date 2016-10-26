package unsafe_test

import (
	"fmt"
	"testing"

	"github.com/vipally/gx/unsafe"
)

//test result
//String to bytes:
//0	s=     	ptr(v)=0x51bd70    	ptr(StringBytes(v)=0x51bd70    	ptr([]byte(v)=0xc042021c58
//1	s=     	ptr(v)=0x51bd70    	ptr(StringBytes(v)=0x51bd70    	ptr([]byte(v)=0xc042021c58
//2	s=hello	ptr(v)=0x51c2fa    	ptr(StringBytes(v)=0x51c2fa    	ptr([]byte(v)=0xc042021c58
//3	s=hello	ptr(v)=0x51c2fa    	ptr(StringBytes(v)=0x51c2fa    	ptr([]byte(v)=0xc042021c58
//4	s=     	ptr(v)=<nil>       	ptr(StringBytes(v)=<nil>       	ptr([]byte(v)=0xc042021c58
//5	s=     	ptr(v)=<nil>       	ptr(StringBytes(v)=<nil>       	ptr([]byte(v)=0xc042021c58
//6	s=xello	ptr(v)=0xc0420444b5	ptr(StringBytes(v)=0xc0420444b5	ptr([]byte(v)=0xc042021c58
//7	s=xello	ptr(v)=0xc0420444ba	ptr(StringBytes(v)=0xc0420444ba	ptr([]byte(v)=0xc042021c58
//Bytes to string:
//0	s=     	ptr(v)=0x5c38b8    	ptr(StringBytes(v)=0x5c38b8    	ptr(string(v)=<nil>
//1	s=hello	ptr(v)=0xc0420445e0	ptr(StringBytes(v)=0xc0420445e0	ptr(string(v)=0xc042021c38
//Benchmark_Normal-4   	1000000000	         0.87 ns/op
//Benchmark_Direct-4   	2000000000	         0.24 ns/op

func TestPointer(t *testing.T) {
	s := []string{
		"",
		"",
		"hello",
		"hello",
		fmt.Sprintf(""),
		fmt.Sprintf(""),
		fmt.Sprintf("hello"),
		fmt.Sprintf("hello"),
	}
	fmt.Println("String to bytes:")
	for i, v := range s {
		b := unsafe.StringBytes(v)
		b2 := []byte(v)
		if b.Writeable() {
			b[0] = 'x'
		}
		fmt.Printf("%d\ts=%5s\tptr(v)=%-12v\tptr(StringBytes(v)=%-12v\tptr([]byte(v)=%-12v\n",
			i, v, unsafe.StringPointer(v), b.Pointer(), unsafe.BytesPointer(b2))
	}

	b := [][]byte{
		[]byte{},
		[]byte{'h', 'e', 'l', 'l', 'o'},
	}
	fmt.Println("Bytes to string:")
	for i, v := range b {
		s1 := unsafe.BytesString(v)
		s2 := string(v)
		fmt.Printf("%d\ts=%5s\tptr(v)=%-12v\tptr(StringBytes(v)=%-12v\tptr(string(v)=%-12v\n",
			i, s1, unsafe.BytesPointer(v), s1.Pointer(), unsafe.StringPointer(s2))
	}

}

const N = 3000000

func Benchmark_Normal(b *testing.B) {
	for i := 1; i < N; i++ {
		s := fmt.Sprintf("12345678901234567890123456789012345678901234567890")
		bb := []byte(s)
		bb[0] = 'x'
		s = string(bb)
		s = s
	}
}
func Benchmark_Direct(b *testing.B) {
	for i := 1; i < N; i++ {
		s := fmt.Sprintf("12345678901234567890123456789012345678901234567890")
		bb := unsafe.StringBytes(s)
		bb[0] = 'x'
		s = s
	}
}
