package fs

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	thisFile      = "github.com/vipally/gx/fs/file_line.go"
	myCallerFrame = 1 //1 means my caller it self
)

var (
	gGoPath   = goPath()
	gWorkPath = workPath()
)

//the caller's file/line info
func FileLine() (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(myCallerFrame); __ok {
		file, line = __file, __line
	}
	return
}

//the caller's file/line relate GoPath
func RelateFileLine() (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(myCallerFrame); __ok {
		file, line = RelateGoPath(__file), __line
	}
	return
}

//the caller's func path
func Func() (f string) {
	if __pc, _, __line, __ok := runtime.Caller(myCallerFrame); __ok {
		__f := runtime.FuncForPC(__pc)
		f = fmt.Sprintf("%s : %d", __f.Name(), __line)
	}
	return
}

func goPath() (p string) {
	f, _ := FileLine()
	p = strings.TrimSuffix(f, thisFile)
	return
}

//GoPath
func GoPath() string {
	return gGoPath
}

func workPath() (p string) {
	if dir, err := os.Getwd(); err == nil {
		p = dir
	} else {
		panic(err)
	}
	return
}

//working path
func WorkPath() (dir string) {
	return gWorkPath
}

//path related to GoPath
func RelateGoPath(filePath string) string {
	return FilePath(filePath).RelateGoPath().String()
}

//path related to WorkPath
func RelateWorkPath(filePath string) string {
	return FilePath(filePath).RelateWorkPath().String()
}

//path related to root
func RelatePath(filePath, root string) string {
	return FilePath(filePath).Relate(root).String()
}

//temp dir
func TempDir() string {
	return FormatPath(os.TempDir())
}

func Chdir(filePath string) (err error) {
	if err = os.Chdir(FormatPath(filePath)); err == nil {
		gWorkPath = workPath() //refresh this lib
	}
	return
}

//to format "\\" with "/"
func FormatPath(filePath string) string {
	return FilePath(filePath).String()
}
