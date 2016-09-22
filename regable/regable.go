//Assuming that all of the acts are registered in the initialization phase is complete
//so do not use lock at regers

//Package regable supply some ways of regist strings,bits variables,then operate then with ids
package regable

import (
	"bytes"
	"fmt"

	acst "vipally.gmail.com/basic/consts"
	aerr "vipally.gmail.com/basic/errors"
)

const (
	g_invalid_id   uint32 = 0
	g_max_id_bits         = 32
	g_real_id_bits        = 24
	g_get_real_id         = (1<<g_max_id_bits - 1) >> (g_max_id_bits - g_real_id_bits)
	g_max_reger_id        = 1<<(g_max_id_bits-g_real_id_bits) - 1
	g_max_real_id         = 1<<g_real_id_bits - 1
	g_no_name             = "no name"
)

var (
	g_reg_locked              = false
	errid_max_reg_overflow, _ = aerr.Reg("max_regs overflow")
)

type showRegFunc func() string

var (
	g_show_reg_funcs []showRegFunc
)

func verify_max_regs(max_regs uint32) (err error) {
	if max_regs < 1 || max_regs > g_max_real_id {
		err = aerr.New(errid_max_reg_overflow)
	}
	return
}

func reg_show(f showRegFunc) {
	g_show_reg_funcs = append(g_show_reg_funcs, f)
}
func ShowAllRegs() string {
	str_head := fmt.Sprintf("[regable] locked:%v", g_reg_locked)
	var buf bytes.Buffer
	buf.WriteString(str_head)
	for _, f := range g_show_reg_funcs {
		buf.WriteString(acst.NEW_LINE)
		buf.WriteString(f())
	}
	return buf.String()
}

func FinishReg() {
	g_reg_locked = true
}

func check_lock() (err error) {
	if g_reg_locked {
		s := "Reg when regs finished"
		err = fmt.Errorf(s)
		panic(s)
	}
	return
}

type regedId uint32

func MakeRegedId(reger_id uint32, real_id uint32) uint32 {
	return reger_id<<g_real_id_bits | (real_id & g_get_real_id)
}
func (cp regedId) ids() (reger_id uint32, id uint32) {
	reger_id = uint32(cp >> g_real_id_bits)
	id = uint32(cp & g_get_real_id)
	return
}
func (cp regedId) id() (id uint32) {
	_, id = cp.ids()
	return
}

//-----------------------------------------------
type RegList struct {
	list []*RegedObjMgr
}
type RegedObjMgr struct {
	id      int //id of this mgr
	idmap   map[RegedId]*RegedObj
	namemap map[string]RegedId
}
type RegedId uint //id of reged things
type RegedObj struct {
	name    string //name of this reg thing
	id      RegedId
	ref     uint   //ref count,lock free operation
	contend string //real reged thing
}
