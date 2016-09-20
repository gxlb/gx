//Package seq implements some useful sequence generating objects
package seq

import (
	"fmt"
	aerr "vipally.gmail.com/basic/errors"
	amath "vipally.gmail.com/basic/math"
	async "vipally.gmail.com/basic/sync"
)

var (
	errid_Seq64_err, _ = aerr.Reg("Seq64 error")
)

func SeqNext64(seq *async.AutoLock) (r uint64) {
	s := seq.LockGet().(*Seq64)
	defer seq.Unlock()
	r, _ = s.Next()
	return
}

type Seq64 struct {
	id        amath.RangeUInt64
	id_name   string
	id_option SeqOption
}

func NewSeq64(_min, _max uint64,
	_name string, _option SeqOption) *Seq64 {
	p := &Seq64{
		id_name:   _name,
		id_option: _option}
	err := p.id.Init(_min, _min, _max)
	if nil != err {
		//need log
		return nil
	}
	p.id.Reset()
	if (_option & _get_usage_option) == NOUSE_FST {
		p.id.Inc()
	}
	return p
}

func (me *Seq64) Current() (r uint64) {
	r, _, _ = me.id.Get()
	return
}
func (me *Seq64) Min() (r uint64) {
	return me.id.Min()
}
func (me *Seq64) Max() (r uint64) {
	return me.id.Max()
}
func (me *Seq64) Next() (r uint64, err error) {
	defer func() { //error handler
		if err != nil && me.id_option&PANIC_ERR == PANIC_ERR {
			panic(err)
		}
	}()

	if r, err = me.pop_recycle(REC_FST); err != nil { //use recycled first
		return
	}

	r, err = me.id.Inc()
	if err != nil {
		if me.id_option&_get_cecycle == NO_REC {
			s := fmt.Sprintf("Seq used up:%#v", *me)
			panic(s)
		}
		if r, err = me.pop_recycle(REC_LST); err != nil {
			return
		}
	}

	return
}

func (me *Seq64) NextNoErr() (r uint64) {
	var err error
	if r, err = me.Next(); err != nil {
		//log
		fmt.Println(err)
	}
	return
}

func (me *Seq64) Recycle(r uint64) (err error) {
	if me.id.InCurrentRange(r) {
		err = me.push_recycle(r)
	} else {
		err = aerr.Newf(errid_Seq64_err, "Recycle(%d) not range : %s", r, me.Infor())
	}
	return
}

func (me *Seq64) Infor() string {
	s := fmt.Sprintf("%#v", *me)
	return s
}

//recycled id manager
//storage depend
func (me *Seq64) push_recycle(r uint64) (err error) {
	return
}
func (me *Seq64) pop_recycle(_recycle_type SeqOption) (r uint64, err error) {
	if (me.id_option & _get_cecycle) == _recycle_type {

	}
	return
}
