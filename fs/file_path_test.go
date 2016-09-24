package fs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	//"fmt"
	//"math"
	//"path/filepath"
	//"path"
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
	//p := FilePath("e:\\")
	//fmt.Println(p)
	//p.Statistic()

	dirs := []string{"e://", "\\localhost\\temp\\"}
	for _, v := range dirs {
		dirPath := path.Clean(v)
		dirFilePath := filepath.Clean(v)
		fmt.Printf("path.Clean([%s])=[%s]\n", v, dirPath)
		fmt.Printf("filepath.Clean([%s])=[%s]\n", v, dirFilePath)
		filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if info.IsDir() {
					fmt.Printf("walk [%s] name=[%s] size=%d\n", path, info.Name(), info.Size())
				}
			} else {
				fmt.Println(err)
			}
			return nil
		})
	}

}
