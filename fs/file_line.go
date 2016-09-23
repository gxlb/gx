package fs

import (
	"runtime"
	"strings"
)

const (
	thisFile      = "github.com/vipally/gx/fs/file_line.go"
	myCallerFrame = 1 //1 means my caller it self
)

var (
	gGoPath = goPath()
)

//the caller's file/line info
func FileLine() (file string, line int) {
	if _, __file, __line, __ok := runtime.Caller(myCallerFrame); __ok {
		file, line = __file, __line
	}
	return
}

//the caller's func path
func Func() (f string) {
	if __pc, _, _, __ok := runtime.Caller(myCallerFrame); __ok {
		__f := runtime.FuncForPC(__pc)
		f = __f.Name()
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

//path related to GoPath
func RelateGoPath(file_path string) string {
	return RelatePath(file_path, gGoPath)
}

//path related to Root
func RelatePath(file_path, root string) string {
	return strings.TrimPrefix(file_path, root)
}
