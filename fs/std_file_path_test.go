package fs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func __TestStdFilePath(t *testing.T) {

	dirs := []string{"e:/", "//localhost/share/", "\\\\localhost\\share\\", "https://github.com/vipally/../allydale/"}
	for _, v := range dirs {
		dirPath := path.Clean(v)
		dirFilePath := filepath.Clean(v)
		fmt.Printf("path.Clean(%s)=%s\n", v, dirPath)
		fmt.Printf("filepath.Clean(%s)=%s\n", v, dirFilePath)
		if e := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if info.IsDir() {
					//fmt.Printf("walk %s name=%s size=%d\n", path, info.Name(), info.Size())
				}
			} else {
				fmt.Println(err)
			}
			return err
		}); e != nil {
			//panic(e)
		}
	}

}
