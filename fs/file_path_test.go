package fs

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestFilePath(t *testing.T) {
	//	_path := "d:/base\\debug/app.go"
	//	fp := FilePath(_path)
	//	fmt.Println(fp)
	//	fp.Set("c:/db\\ab.txt")
	//	fmt.Println(fp)

	//	var s FileSize = math.MaxUint64 - math.MaxUint64 + 2000000000000000 // - 17999999999999999999 // 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	//	fmt.Println(s)
	//	s.Set(1024)
	//	fmt.Println(s)

	//	fmt.Println(FileLine())
	//	fmt.Println(Func())
	//	fmt.Println(GoPath())
	//	f, _ := FileLine()
	//	fmt.Println(f)
	//	fmt.Println(RelateGoPath(f))
	//	fmt.Println(RelatePath(f, "E:/dev"))
	//	fmt.Println(filepath.Clean("E:/dev/../app/hello\\x.txt"))
	p := FilePath("E:/temp")
	//p := FilePath("//localhost/share/share_dir/share_file_in_share_dir.txt")
	//fmt.Println(p.Rename("share_file_in_share_dir2.txt"))
	//fmt.Println(p.Truncate(20))
	//	fmt.Println(p.Statistic())
	//	fmt.Println(FilePath(WorkPath()).RelateGoPath())
	//	fmt.Println(WorkPath())
	//	fmt.Println(FilePath("E:/dev/gocode/trunk/src/github.com/vipally/gx/fs/x/h.txt").RelateWorkPath())
	//	fmt.Println(WorkPath())
	if n, e := p.Hash(NONE, ""); e == nil {
		l := n.All()
		for _, v := range l {
			fmt.Println(v.Hash, v.Path, v.Size)
		}
		//		{
		//			v := n.Visitor()
		//			for v.Next() {
		//				g := v.Get()
		//				fmt.Println(g.Hash, g.Path, g.Size)
		//			}
		//		}
		fmt.Println("visiting...")
		{
			v := n.Visitor()
			for v.Prev() {
				g := v.Get()
				fmt.Println(g.Hash, g.Path, g.Size)
			}
		}

	} else {
		fmt.Println(e)
	}

	//	fmt.Println(os.Getegid(), os.Geteuid(), os.Getpid(), os.Getpagesize())
	//	fmt.Println(os.Hostname())
	//	p.Statistic()

	//	dirs := []string{"e:/", "//localhost/share/", "\\\\localhost\\share\\", "\\\\localhost\\share\\share_dir\\xx.txt\\yy\\zz\\aa.txt", "c:/;d:/;e:"}
	//	for _, v := range dirs {
	//		fp := FilePath(v)
	//		fmt.Println(v, fp, fp.StringSys())
	//		fmt.Println(fp.SplitList())
	//		list := fp.SplitAll()
	//		fmt.Println(list, Joins(list...))
	//		fmt.Println(fp.Join("a/b/hello.txt"))
	//	}

}

func TestFileLock(t *testing.T) {
	_path := "e:/url.txt"
	h, err := syscall.CreateFile(syscall.StringToUTF16Ptr(_path),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
		0)
	if err != nil {
		fmt.Println("e0", err)
		return
	}
	f := os.NewFile(uintptr(h), _path)
	defer func() {
		if err != nil {
			f.Close()
		}
	}()
	f0, e0 := os.OpenFile(_path, os.O_CREATE, os.ModePerm)
	f1, e1 := os.Open(_path)
	f.Close()
	f2, e2 := os.Open(_path)
	if e0 == nil {
		f0.Close()
	} else {
		fmt.Println("e0", e0)
	}
	if e1 == nil {
		f1.Close()
	} else {
		fmt.Println("e1", e1)
	}
	if e2 == nil {
		f2.Close()
	} else {
		fmt.Println("e2", e2)
	}
}
