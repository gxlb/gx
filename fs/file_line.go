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

//the caller's file/line info
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
func RelateGoPath(file_path string) string {
	return RelatePath(file_path, GoPath())
}

//path related to root
func RelatePath(file_path, root string) string {
	return strings.TrimPrefix(file_path, root)
}

//to format "\\" with "/"
func FormatPath(filepath string) string {
	return FilePath(filepath).String()
}
