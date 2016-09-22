package fs

import (
	"fmt"
	"math"
	"testing"
)

func TestFilePath(t *testing.T) {
	path := "d:/base\\debug/app.go"
	fp := ToFilePath(path)
	fmt.Println(fp)
	fp.Set("c:/db\\ab.txt")
	fmt.Println(fp)

	var s FileSize = math.MaxUint64 - math.MaxUint64 + 2000000000000000 // - 17999999999999999999 // 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	fmt.Println(s)
	s.Set(1024)
	fmt.Println(s)
}
