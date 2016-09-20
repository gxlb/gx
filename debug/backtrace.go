//Package debug implements some useful funcs Bt() and Bts() to view backtrace infomation for debug
package debug

import (
	"bytes"
	"fmt"
	"runtime"
)

type t_fmt uint32

const (
	BT_ME     = 2
	BT_FATHER = BT_ME + 1
)

const (
	max_bt_list_len       = 50
	fmt_fileline    t_fmt = iota + 1
	fmt_func
	fmt_all
)
const (
	str_caller    = "Caller:"
	str_cur_frame = ""
)

//tell me current call frame
func Bt() string {
	return CallStackInfo(BT_ME, fmt_all, str_cur_frame)
}

//tell me current fileline
func BtFileLine() string {
	return CallStackInfo(BT_ME, fmt_fileline, str_cur_frame)
}

//tell me current func
func BtFunc() string {
	return CallStackInfo(BT_ME, fmt_func, str_cur_frame)
}

//tell me caller call frame
func Caller(_depth int) string {
	return CallStackInfo(BT_FATHER+_depth, fmt_all, str_caller)
}

//tell me caller fileline
func CallerFileLine(_depth int) string {
	return CallStackInfo(BT_FATHER+_depth, fmt_fileline, str_caller)
}

//tell me caller func
func CallerFunc(_depth int) string {
	return CallStackInfo(BT_FATHER+_depth, fmt_func, str_caller)
}

//use a fast-way to build string
func Bts() (bts string) {
	var buf bytes.Buffer
	var __pcs [max_bt_list_len]uintptr
	n := runtime.Callers(BT_ME, __pcs[0:])
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

func CallStackInfo(nDepth int, _fmt t_fmt, strhead string) (_info string) {
	__pc, __file, __line, __ok := runtime.Caller(nDepth)

	if __ok {
		__f := runtime.FuncForPC(__pc)
		switch _fmt {
		case fmt_fileline:
			_info = fmt.Sprintf("%s FileLine{%s : %d}",
				strhead, __file, __line)
		case fmt_func:
			_info = fmt.Sprintf("%s Func{%s}", strhead, __f.Name())
		case fmt_all:
			fallthrough
		default:
			_info = fmt.Sprintf("%s Func{%s} FileLine{%s : %d} pc{%0X}",
				strhead, __f.Name(), __file, __line, __pc)
		}
	}
	return
}

func CallStackList(nDepth int) []string {
	var __pcs [max_bt_list_len]uintptr
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
