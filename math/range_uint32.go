package math

import (
	"fmt"
	"math"
	//aerr "vipally.gmail.com/basic/errors"
)

var (
//errid_ru32_full, _ = aerr.Reg("[RangeUint32] out of range")
)

//ranged uint32 object
type RangeUint32 struct {
	val, min, max uint32
}

func NewRangeUint32(_val, _min, _max uint32) (r *RangeUint32, err error) {
	if err = verifyRangeUint32(_val, _min, _max); err == nil {
		r = &RangeUint32{val: _val, min: _min, max: _max}
	}
	return
}

func (me *RangeUint32) Reset() bool {
	me.val = me.min
	return true
}

func (me *RangeUint32) ResetR() bool {
	me.val = me.max
	return true
}

func (me *RangeUint32) Init(_val, _min, _max uint32) (err error) {
	if err = verifyRangeUint32(_val, _min, _max); err == nil {
		me.val, me.min, me.max = _val, _min, _max
	}
	return
}

func (me *RangeUint32) Add(other uint32) (r uint32, err error) {
	if math.MaxUint32-other < me.val {
		//return me.min, aerr.Newf(errid_ru32_full, "Add(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val + other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUint32) Sub(other uint32) (r uint32, err error) {
	if me.val < other {
		//return me.min, aerr.Newf(errid_ru32_full, "Sub(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val - other
	r, err = v, me.verify(v)
	return
}

func (me *RangeUint32) Mul(other uint32) (r uint32, err error) {
	if other > 0 && math.MaxUint32/other < me.val {
		//return me.min, aerr.Newf(errid_ru32_full, "Mul(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val * other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUint32) Div(other uint32) (r uint32, err error) {
	if 0 == other {
		//return me.min, aerr.Newf(errid_ru32_full, "Div(%d) error:%s", other, me.Info())
	}
	v := me.val / other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUint32) Mod(other uint32) (r uint32, err error) {
	if 0 == other {
		return me.min, fmt.Errorf("[RangeUint32] Mod(%d) error:%s", other, me.Info())
	}
	v := me.val % other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUint32) Inc() (r uint32, err error) {
	if math.MaxUint32 == me.val {
		//return me.min, aerr.Newf(errid_ru32_full, "Inc() will cause out of range :%s", me.Info())
	}
	r, err = me.Set(me.val + 1)
	return
}
func (me *RangeUint32) Dec(other uint32) (r uint32, err error) {
	if 0 == me.val {
		//return me.min, aerr.Newf(errid_ru32_full, "Dec() will cause out of range :%s", me.Info())
	}
	r, err = me.Set(me.val - 1)
	return
}
func (me *RangeUint32) Set(other uint32) (r uint32, err error) {
	r = me.val
	if err = me.verify(other); err == nil {
		me.val = other
	}
	return
}
func (me *RangeUint32) Get() (r_val, r_min, r_max uint32) {
	return me.val, me.min, me.max
}
func (me *RangeUint32) Val() uint32 {
	return me.val
}
func (me *RangeUint32) Min() uint32 {
	return me.min
}
func (me *RangeUint32) Max() uint32 {
	return me.max
}
func (me *RangeUint32) InRange(other uint32) (r bool) {
	r = (other >= me.min && other <= me.max)
	return
}
func (me *RangeUint32) InCurrentRange(other uint32) (r bool) {
	r = (other >= me.min && other <= me.val)
	return
}

func (me *RangeUint32) String() string {
	s := fmt.Sprintf("RangeUint32{val:%d min:%d max:%d}",
		me.val, me.min, me.max)
	return s
}

func (me *RangeUint32) Info() string {
	s := fmt.Sprintf("%#v", *me)
	return s
}

func (me *RangeUint32) verify(_val uint32) (err error) {
	if !me.InRange(_val) {
		//err = aerr.Newf(errid_ru32_full, "val(%d) out of range[%d,%d]", _val, me.min, me.max)
	}
	return
}

func verifyRangeUint32(_val, _min, _max uint32) (err error) {
	if _min > _max {
		//err = aerr.New(errid_min_max_error)
	} else if _val < _min || _val > _max {
		//err = aerr.Newf(errid_ru32_full, "val(%d) out of range[%d,%d]", _val, _min, _max)
	}
	return
}
