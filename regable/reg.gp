package regable

import (
	"bytes"
	"fmt"
	<LOCK_COMMENT>"sync"
	acst "vipally.gmail.com/basic/consts"
	aerr "vipally.gmail.com/basic/errors"
	amath "vipally.gmail.com/basic/math"
)

const (
	default_<TYPENAME_L_SHORT>_reg_cnt = <DEFAULT_REG_CNT>
)

var (
	g_<TYPENAME_L_SHORT>_rgr_id_gen, _         = amath.NewRangeUint32(g_invalid_id+1, g_invalid_id, g_max_reger_id)
	errid_<TYPENAME_L_SHORT>_id, _  = aerr.Reg("<TYPENAME_U>Id error")
	errid_<TYPENAME_L_SHORT>_obj, _ = aerr.Reg("<TYPENAME_U> object error")
)

var (
	g_<TYPENAME_L_SHORT>_reger_list []*<TYPENAME_U>Reger
)

func init() {
	reg_show(ShowAll<TYPENAME_U>Regers)
}

//new reger
func New<TYPENAME_U>Reger(name string) (r *<TYPENAME_U>Reger, err error) {
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	if err = check_lock(); err != nil {
		return
	}
	id := g_invalid_id
	if id, err = g_<TYPENAME_L_SHORT>_rgr_id_gen.Inc(); err != nil {
		return
	}
	p := new(<TYPENAME_U>Reger)
	if err = p.init(name); err == nil {
		p.reger_id = uint8(id)
		r = p
		g_<TYPENAME_L_SHORT>_reger_list = append(g_<TYPENAME_L_SHORT>_reger_list, p)
	}
	return
}

func MustNew<TYPENAME_U>Reger(name string) (r *<TYPENAME_U>Reger) {
	if reg, err := New<TYPENAME_U>Reger(name); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}

//show all regers
func ShowAll<TYPENAME_U>Regers() string {
	var buf bytes.Buffer
	s := fmt.Sprintf("[<TYPENAME_U>Regers] count:%d", len(g_<TYPENAME_L_SHORT>_reger_list))
	buf.WriteString(s)
	for _, v := range g_<TYPENAME_L_SHORT>_reger_list {
		buf.WriteString(acst.NEW_LINE)
		buf.WriteString(v.String())
	}
	return buf.String()
}

//reger object
type <TYPENAME_U>Reger struct {
	reger_id uint8
	name     string
	id_gen   amath.RangeUint32
	reg_list []*_<TYPENAME_L>Record
}

func (me *<TYPENAME_U>Reger) init(name string) (err error) {
	me.name = name
	if err = me.id_gen.Init(g_invalid_id+1, g_invalid_id,
		g_invalid_id+default_<TYPENAME_L_SHORT>_reg_cnt); err != nil {
		return
	}
	me.reg_list = make([]*_<TYPENAME_L>Record, 0, 0)
	return
}

//set max reg count at a reger
func (me *<TYPENAME_U>Reger) MaxReg(max_regs uint32) (rmax uint32, err error) {
	if err = verify_max_regs(max_regs); err != nil {
		return
	}
	cur, min, _ := me.id_gen.Get()
	if err = me.id_gen.Init(cur, min, g_invalid_id+max_regs); err != nil {
		return
	}
	rmax = me.id_gen.Max()
	return
}

//reg a value
func (me *<TYPENAME_U>Reger) Reg(<NAME_COMMENT_B>name string,<NAME_COMMENT_E> val <VALUE_TYPE>) (r <TYPENAME_U>Id, err error) {
	r = <TYPENAME_U>Id(g_invalid_id)
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	id := g_invalid_id
	if err = check_lock(); err != nil {
		return
	}
	if id, err = me.id_gen.Inc(); err != nil {
		return
	}
	p := me.new_rec(<NAME_COMMENT_B>name,<NAME_COMMENT_E> val)
	p.id = id
	me.reg_list = append(me.reg_list, p)
	r = <TYPENAME_U>Id(MakeRegedId(uint32(me.reger_id), id))
	return
}

func (me *<TYPENAME_U>Reger) MustReg(<NAME_COMMENT_B>name string,<NAME_COMMENT_E> val <VALUE_TYPE>) (r <TYPENAME_U>Id) {
	if reg, err := me.Reg(<NAME_COMMENT_B>name,<NAME_COMMENT_E> val); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}

//show string
func (me *<TYPENAME_U>Reger) String() string {
	var buf bytes.Buffer
	s := fmt.Sprintf("[<TYPENAME_U>Reger#%d: %s] ids:%s",
		me.reger_id, me.name, me.id_gen.String())
	buf.WriteString(s)
	for i, v := range me.reg_list {
		<LOCK_COMMENT>v.lock.RLock()
		s = fmt.Sprintf("\n#%d [%s]: %v",
			uint32(i)+g_invalid_id+1, <NAME_INSTEAD><NAME_COMMENT_B>v.name<NAME_COMMENT_E>,
			v.val)
		<LOCK_COMMENT>v.lock.RUnlock()
		buf.WriteString(s)
	}
	return buf.String()
}

type _<TYPENAME_L>Record struct {
	<NAME_COMMENT_B>name string<NAME_COMMENT_E>
	val  <VALUE_TYPE>
	id   uint32
	<LOCK_COMMENT>lock sync.RWMutex
}

func (me *<TYPENAME_U>Reger) new_rec(<NAME_COMMENT_B>name string,<NAME_COMMENT_E> val <VALUE_TYPE>) (r *_<TYPENAME_L>Record) {
	r = new(_<TYPENAME_L>Record)
	<NAME_COMMENT_B>r.name = name<NAME_COMMENT_E>
	r.val = val
	return
}

type <TYPENAME_U>Id regedId

func (cp <TYPENAME_U>Id) get() (rg *<TYPENAME_U>Reger, r *_<TYPENAME_L>Record, err error) {
	idrgr, id := regedId(cp).ids()
	idregidx, idx := idrgr-g_invalid_id-1, id-g_invalid_id-1
	if idrgr == g_invalid_id || !g_<TYPENAME_L_SHORT>_rgr_id_gen.InCurrentRange(idrgr) {
		err = aerr.New(errid_<TYPENAME_L_SHORT>_id)
	}
	rg = g_<TYPENAME_L_SHORT>_reger_list[idregidx]
	if id == g_invalid_id || !rg.id_gen.InCurrentRange(id) {
		err = aerr.New(errid_<TYPENAME_L_SHORT>_id)
	}
	r = rg.reg_list[idx]
	return
}

//check if valid
func (cp <TYPENAME_U>Id) Valid() (rvalid bool) {
	if _, _, e := cp.get(); e == nil {
		rvalid = true
	}
	return
}

//get value
func (cp <TYPENAME_U>Id) Get() (r <VALUE_TYPE>, err error) {
	_, rc, e := cp.get()
	if e != nil {
		return r, e
	}
	return rc.Get()
}

//get value with out error, if has error will cause panic
func (cp <TYPENAME_U>Id) GetNoErr() (r <VALUE_TYPE>) {
	_, rc, e := cp.get()
	if e != nil {
		panic(e.Error())
	}
	return rc.GetNoErr()
}

//set value
<LOCK_COMMENT>func (cp <TYPENAME_U>Id) Set(val <VALUE_TYPE>) (r <VALUE_TYPE>, err error) {
<LOCK_COMMENT>	_, rc, e := cp.get()
<LOCK_COMMENT>	if e != nil {
<LOCK_COMMENT>		return r, e
<LOCK_COMMENT>	}
<LOCK_COMMENT>	return rc.Set(val)
<LOCK_COMMENT>}

//reverse bool value(as a switch)
<REVERSE_COMMENT>func (cp <TYPENAME_U>Id) Reverse() (r <VALUE_TYPE>, err error) {
<REVERSE_COMMENT>	_, rc, e := cp.get()
<REVERSE_COMMENT>	if e != nil {
<REVERSE_COMMENT>		return r, e
<REVERSE_COMMENT>	}
<REVERSE_COMMENT>	return rc.Reverse()
<REVERSE_COMMENT>}

//get reger_id and real_id
func (cp <TYPENAME_U>Id) Ids() (reger_id, real_id uint32) {
	return regedId(cp).ids()
}

//show string
func (cp <TYPENAME_U>Id) String() (r string) {
	idrgr, id := regedId(cp).ids()
	_, rc, err := cp.get()
	if err != nil {
		r = fmt.Sprintf("invalid <TYPENAME_U>Id#(%d|%d)", idrgr, id)
	} else {
		r = rc.String()
	}
	return
}

//get name
<NAME_COMMENT_B>
func (cp <TYPENAME_U>Id) Name() (r string, err error) {
	_, rc, e := cp.get()
	if e == nil {
		r, err = rc.Name()
	} else {
		err = e
	}
	return
}
<NAME_COMMENT_E>

//get as object for fast access
func (cp <TYPENAME_U>Id) Oject() (r <TYPENAME_U>Obj) {
	_, rc, e := cp.get()
	if e == nil {
		r.obj = rc
	}
	return
}

//get name
<NAME_COMMENT_B>
func (me *_<TYPENAME_L>Record) Name() (r string, err error) {
	if me != nil {
<LOCK_COMMENT>		me.lock.RLock()
<LOCK_COMMENT>		defer me.lock.RUnlock()
		r = me.name
	} else {
		err = aerr.New(errid_<TYPENAME_L_SHORT>_obj)
	}
	return
}
<NAME_COMMENT_E>

//get value
func (me *_<TYPENAME_L>Record) Get() (r <VALUE_TYPE>, err error) {
	if me != nil {
<LOCK_COMMENT>		me.lock.RLock()
<LOCK_COMMENT>		defer me.lock.RUnlock()
		r = me.val
	} else {
		err = aerr.New(errid_<TYPENAME_L_SHORT>_obj)
	}
	return
}

//get value without error,if has error will cause panic
func (me *_<TYPENAME_L>Record) GetNoErr() (r <VALUE_TYPE>) {
	r0, err := me.Get()
	if err != nil {
		panic(err.Error())
	}
	r = r0
	return
}

//set value
<LOCK_COMMENT>func (me *_<TYPENAME_L>Record) Set(val <VALUE_TYPE>) (r <VALUE_TYPE>, err error) {
<LOCK_COMMENT>	if nil != me {
<LOCK_COMMENT>		me.lock.Lock()
<LOCK_COMMENT>		defer me.lock.Unlock()
<LOCK_COMMENT>		me.val = val
<LOCK_COMMENT>		r = val
<LOCK_COMMENT>	} else {
<LOCK_COMMENT>		err = aerr.New(errid_<TYPENAME_L_SHORT>_obj)
<LOCK_COMMENT>	}
<LOCK_COMMENT>	return
<LOCK_COMMENT>}

//reverse on bool value
<REVERSE_COMMENT>func (me *_<TYPENAME_L>Record) Reverse() (r <VALUE_TYPE>, err error) {
<REVERSE_COMMENT>	if nil != me {
<REVERSE_COMMENT><LOCK_COMMENT>		me.lock.Lock()
<REVERSE_COMMENT><LOCK_COMMENT>		defer me.lock.Unlock()
<REVERSE_COMMENT>		me.val = !me.val
<REVERSE_COMMENT>		r = me.val
<REVERSE_COMMENT>	} else {
<REVERSE_COMMENT>		err = aerr.New(errid_<TYPENAME_L_SHORT>_obj)
<REVERSE_COMMENT>	}
<REVERSE_COMMENT>	return
<REVERSE_COMMENT>}

//get as Id
func (me *_<TYPENAME_L>Record) Id() (r <TYPENAME_U>Id) {
	if me != nil {
		r = <TYPENAME_U>Id(me.id)
	}
	return
}

//show string
func (me *_<TYPENAME_L>Record) String() (r string) {
	if me != nil {
		idrgr, id := regedId(me.id).ids()
<LOCK_COMMENT>		me.lock.RLock()
<LOCK_COMMENT>		defer me.lock.RUnlock()
		r = fmt.Sprintf("<TYPENAME_U>#(%d|%d|%s)%v", idrgr, id, <NAME_INSTEAD><NAME_COMMENT_B>me.name<NAME_COMMENT_E>, me.val)
	} else {
		r = fmt.Sprintf("invalid <VALUE_TYPE> object")
	}
	return
}

//object of reged value,it is more efficient to access than Id object
type <TYPENAME_U>Obj struct {
	obj *_<TYPENAME_L>Record
}

//check if valid
func (cp <TYPENAME_U>Obj) Valid() (rvalid bool) {
	return cp.obj != nil
}

//get value
func (cp <TYPENAME_U>Obj) Get() (r <VALUE_TYPE>, err error) {
	return cp.obj.Get()
}

//get value against error,if has error will cause panic
func (cp <TYPENAME_U>Obj) GetNoErr() (r <VALUE_TYPE>) {
	return cp.obj.GetNoErr()
}

//set value
<LOCK_COMMENT>func (cp <TYPENAME_U>Obj) Set(val <VALUE_TYPE>) (r <VALUE_TYPE>, err error) {
<LOCK_COMMENT>	return cp.obj.Set(val)
<LOCK_COMMENT>}

//reverse bool object
<REVERSE_COMMENT>func (cp <TYPENAME_U>Obj) Reverse() (r <VALUE_TYPE>, err error) {
<REVERSE_COMMENT>	return cp.obj.Reverse()
<REVERSE_COMMENT>}

//show string
func (cp <TYPENAME_U>Obj) String() (r string) {
	return cp.obj.String()
}

//get name
<NAME_COMMENT_B>
func (cp <TYPENAME_U>Obj) Name() (r string, err error) {
	return cp.obj.Name()
}
<NAME_COMMENT_E>

//get as Id
func (cp <TYPENAME_U>Obj) Id() (r <TYPENAME_U>Id) {
	return cp.obj.Id()
}

//reg and return an object agent
func (me *<TYPENAME_U>Reger) RegO(<NAME_COMMENT_B>name string,<NAME_COMMENT_E> val <VALUE_TYPE>) (r <TYPENAME_U>Obj, err error) {
	id, e := me.Reg(<NAME_COMMENT_B>name,<NAME_COMMENT_E> val)
	if e == nil {
		r = id.Oject()
	} else {
		err = e
	}
	return
}

func (me *<TYPENAME_U>Reger) MustRegO(<NAME_COMMENT_B>name string,<NAME_COMMENT_E> val <VALUE_TYPE>) (r <TYPENAME_U>Obj) {
	if reg, err := me.RegO(<NAME_COMMENT_B>name,<NAME_COMMENT_E> val); err != nil {
		panic(err)
	} else {
		r = reg
	}
	return
}
