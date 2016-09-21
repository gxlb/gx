//Package errors supply some way to send and predefines string error by only passing it's ID
package errors

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/vipally/gx/consts"
	"github.com/vipally/gx/debug"
)

const (
	MaxRef = 100000
)

type ErrorId int
type WithIdError interface {
	error
	Id() ErrorId
}

type ErrMgr struct {
	id2Obj map[ErrorId]*ErrorDef
}
type ErrorDef struct {
	content string
	name    string
	id      ErrorId
	ref     int
}
type ErrorObjRaw struct {
	id ErrorId
}
type ErrorObj struct {
	ErrorObjRaw
	para []interface{}
}

//func (this *ErrorObj) Clone(para ...interface{}) *ErrorObj {
//	return &ErrorObj{id: this.ErrorObjRaw.id, para: para}
//}

//-------------------------------------------------------------------
func init() {
	g_id2err = []string{"unknown error"}
}

const (
	//unknown error id
	ERRID_UNKNOWN = errId(g_id_min)
)

var (
	ERRID_STD, _ = Reg("standard error")
)

var (
	//registered error mapping
	g_id2err      = make([]string, 0, g_init_err)
	g_id2err_lock sync.RWMutex
	g_id_cur_max  errId = g_id_invalid
)

const (
	g_init_err   = 32
	g_id_invalid = g_id_min - 1
	g_id_min     = 1001
	g_id_max     = 9999
)

type WithIdErr interface {
	error
	Id() errId
}

//new std error
func NewError(_fmt string, _para ...interface{}) error {
	return fmt.Errorf(_fmt, _para...)
}

//func Clear() {
//	g_id2err_lock.Lock()
//	defer g_id2err_lock.Unlock()
//	g_id2err = nil
//}

func Reg(_err_name string) (id errId, err error) {
	g_id2err_lock.Lock()
	defer g_id2err_lock.Unlock()
	//idx := g_id_cur_max - g_id_min
	//id = errId(g_id_min + len(g_id2err) + 1)
	if g_id_cur_max < g_id_max {
		g_id2err = append(g_id2err, _err_name)
		g_id_cur_max++
		id = g_id_cur_max
	} else {
		id = g_id_invalid
		err = fmt.Errorf("[Register] error id out of range[%d,%d] from %s",
			g_id_min, g_id_max, debug.Caller(0))
	}
	return
}

func ShowAll() string {
	var buf bytes.Buffer
	s := fmt.Sprintf("[errors] Count:%d", g_id_cur_max-g_id_min+1)
	buf.WriteString(s)
	for id := errId(g_id_min); id <= g_id_cur_max; id++ {
		buf.WriteString(consts.NewLine)
		buf.WriteString(id.String())
	}
	return buf.String()
}

type errId uint32 //never generate by other package
func (cp errId) Get() (_name string, err error) {
	g_id2err_lock.RLock()
	defer g_id2err_lock.RUnlock()
	if err = cp.valid(); err == nil {
		idx := cp - g_id_min
		_name = g_id2err[idx]
	} else {
		_name = consts.EmptyStr
	}
	return
}
func (cp errId) valid() (err error) {
	if ok := cp >= g_id_min && cp <= g_id_cur_max; !ok {
		err = fmt.Errorf("invalid error id:%d", cp)
	}
	return
}
func (cp errId) String() string {
	s, err := cp.Get()
	if err != nil {
		return err.Error()
	}
	s = fmt.Sprintf("error#%d (%s)", cp, s)
	return s
}

func StdErr(e error) (err WithIdErr) {
	return Newf(ERRID_STD, e.Error())
}
