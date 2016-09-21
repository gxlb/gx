//Package debug implements some useful funcs Bt() and Bts() to view backtrace infomation for debug
package debug

import (
	"bytes"
	"fmt"
	"runtime"
)

type tFmt uint32

const (
	BtMe     = 2
	BtFather = BtMe + 1
)

const (
	maxBtListLen      = 50
	fmtFileLine  tFmt = iota + 1
	fmtFunc
	fmtAll
)
const (
	strCaller   = "Caller:"
	strCurFrame = ""
)

//tell me current call frame
func Bt() string {
	return CallStackInfo(BtMe, fmtAll, strCurFrame)
}

//tell me current fileline
func BtFileLine() string {
	return CallStackInfo(BtMe, fmtFileLine, strCurFrame)
}

//tell me current func
func BtFunc() string {
	return CallStackInfo(BtMe, fmtFunc, strCurFrame)
}

//tell me caller call frame
func Caller(depth int) string {
	return CallStackInfo(BtFather+depth, fmtAll, strCaller)
}

//tell me caller fileline
func CallerFileLine(depth int) string {
	return CallStackInfo(BtFather+depth, fmtFileLine, strCaller)
}

//tell me caller func
func CallerFunc(depth int) string {
	return CallStackInfo(BtFather+depth, fmtFunc, strCaller)
}

//use a fast-way to build string
func Bts() (bts string) {
	var buf bytes.Buffer
	var __pcs [maxBtListLen]uintptr
	n := runtime.Callers(BtMe, __pcs[0:])
	for i := 0; i < n; i++ {
		_pc := __pcs[i]
		_f := runtime.FuncForPC(_pc)
		_file, _line := _f.FileLine(_pc)
		s := fmt.Sprintf("#%d Func{%s} File{%s : %d} pc{%0X}\n",
			i, _f.Name(), _file, _line, _pc)
		buf.WriteString(s)
	}

	return buf.String()
}

func CallStackInfo(nDepth int, _fmt tFmt, strhead string) (_info string) {
	__pc, __file, __line, __ok := runtime.Caller(nDepth)

	if __ok {
		__f := runtime.FuncForPC(__pc)
		switch _fmt {
		case fmtFileLine:
			_info = fmt.Sprintf("%s FileLine{%s : %d}",
				strhead, __file, __line)
		case fmtFunc:
			_info = fmt.Sprintf("%s Func{%s}", strhead, __f.Name())
		case fmtAll:
			fallthrough
		default:
			_info = fmt.Sprintf("%s Func{%s} FileLine{%s : %d} pc{%0X}",
				strhead, __f.Name(), __file, __line, __pc)
		}
	}
	return
}

func CallStackList(nDepth int) []string {
	var __pcs [maxBtListLen]uintptr
	n := runtime.Callers(nDepth, __pcs[0:])
	cs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		_pc := __pcs[i]
		_f := runtime.FuncForPC(_pc)
		_file, _line := _f.FileLine(_pc)
		s := fmt.Sprintf("Func{%s} FileLine{%s : %d} pc{%0X}",
			_f.Name(), _file, _line, _pc)
		cs = append(cs, s)
	}
	return cs
}
