// Copyright @Ally 2014. All rights reserved.
// license that can be found in the LICENSE file.
// contact the author by Email: vipally@gmail.com
// blog site: http://blog.sina.com.cn/ally2014

package math

import (
	"fmt"
	"math"
	//aerr "vipally.gmail.com/basic/errors"
)

var (
//errid_ru64_full, _     = aerr.Reg("[RangeUInt64] out of range")
//errid_min_max_error, _ = aerr.Reg("[min > max error]")
)

//ranged uint64 object
type RangeUInt64 struct {
	val, min, max uint64
}

func NewRangeUInt64(_val, _min, _max uint64) (r *RangeUInt64, err error) {
	if err = verifyRangeUInt64(_val, _min, _max); err == nil {
		r = &RangeUInt64{val: _val, min: _min, max: _max}
	}
	return
}

func (me *RangeUInt64) Reset() bool {
	me.val = me.min
	return true
}

func (me *RangeUInt64) ResetR() bool {
	me.val = me.max
	return true
}

func (me *RangeUInt64) Init(_val, _min, _max uint64) (err error) {
	if err = verifyRangeUInt64(_val, _min, _max); err == nil {
		me.val, me.min, me.max = _val, _min, _max
	}
	return
}

func (me *RangeUInt64) Add(other uint64) (r uint64, err error) {
	if math.MaxUint64-other < me.val {
		//return me.min, aerr.Newf(errid_ru64_full, "Add(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val + other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUInt64) Sub(other uint64) (r uint64, err error) {
	if me.val < other {
		//return me.min, aerr.Newf(errid_ru64_full, "Sub(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val - other
	r, err = v, me.verify(v)
	return
}

func (me *RangeUInt64) Mul(other uint64) (r uint64, err error) {
	if other > 0 && math.MaxUint64/other < me.val {
		//return me.min, aerr.Newf(errid_ru64_full, "Mul(%d) will cause out of range :%s", other, me.Info())
	}
	v := me.val * other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUInt64) Div(other uint64) (r uint64, err error) {
	if 0 == other {
		//return me.min, aerr.Newf(errid_ru64_full, "Div(%d) error:%s", other, me.Info())
	}
	v := me.val / other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUInt64) Mod(other uint64) (r uint64, err error) {
	if 0 == other {
		//return me.min, fmt.Errorf("[RangeUInt64] Mod(%d) error:%s", other, me.Info())
	}
	v := me.val % other
	r, err = v, me.verify(v)
	return
}
func (me *RangeUInt64) Inc() (r uint64, err error) {
	if math.MaxUint64 == me.val {
		//return me.min, aerr.Newf(errid_ru64_full, "Inc() will cause out of range :%s", me.Info())
	}
	r, err = me.Set(me.val + 1)
	return
}
func (me *RangeUInt64) Dec(other uint64) (r uint64, err error) {
	if 0 == me.val {
		//return me.min, aerr.Newf(errid_ru64_full, "Dec() will cause out of range :%s", me.Info())
	}
	r, err = me.Set(me.val - 1)
	return
}
func (me *RangeUInt64) Set(other uint64) (r uint64, err error) {
	r = me.val
	if err = me.verify(other); err == nil {
		me.val = other
	}
	return
}
func (me *RangeUInt64) Get() (r_val, r_min, r_max uint64) {
	return me.val, me.min, me.max
}
func (me *RangeUInt64) Val() uint64 {
	return me.val
}
func (me *RangeUInt64) Min() uint64 {
	return me.min
}
func (me *RangeUInt64) Max() uint64 {
	return me.max
}
func (me *RangeUInt64) InRange(other uint64) (r bool) {
	r = (other >= me.min && other <= me.max)
	return
}

func (me *RangeUInt64) InCurrentRange(other uint64) (r bool) {
	r = (other >= me.min && other <= me.val)
	return
}

func (me *RangeUInt64) Info() string {
	s := fmt.Sprintf("%#v", *me)
	return s
}

func (me *RangeUInt64) verify(_val uint64) (err error) {
	if !me.InRange(_val) {
		//err = aerr.Newf(errid_ru64_full, "val(%d) out of range[%d,%d]", _val, me.min, me.max)
	}
	return
}

func verifyRangeUInt64(_val, _min, _max uint64) (err error) {
	if _min > _max {
		//err = aerr.New(errid_min_max_error)
	} else if _val < _min || _val > _max {
		//err = aerr.Newf(errid_ru64_full, "val(%d) out of range[%d,%d]", _val, _min, _max)
	}
	return
}
