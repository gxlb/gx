package rand

import (
	async "vipally.gmail.com/basic/sync"
)

var (
	g_rand64 = async.NewAutoLock(NewRand64S(uint64(RandSeed())))
)

//64-bits rand number generator
type Rand64 struct {
	seed uint64
}

//new a initialized rand64 object
func NewRand64S(_seed uint64) *Rand64 {
	return &Rand64{seed: _seed}
}

//new a rand64 object initialized by auto-generated seed
func NewRand64() *Rand64 {
	r := g_rand64.LockGet().(*Rand64)
	defer g_rand64.Unlock()
	return NewRand64S(r.randBase())
}

//next rand number
func (me *Rand64) Rand() uint64 {
	me.seed = me.seed*4294955897 + 4094975977 //4294967291 4294967279
	return me.seed
}

//new rand seed list
func (me *Rand64) randBase() uint64 {
	me.seed = me.seed*4165989847 + 4098167467
	return me.seed
}

//generate rand number in range
func (me *Rand64) RandRange(_min, _max uint64) uint64 {
	if _max < _min {
		_max, _min = _min, _max
	}
	d := _max - _min + 1
	r := me.Rand()
	ret := r%d + _min

	return ret
}

//generate rand number with max
func (me *Rand64) RandMax(_max uint64) uint64 {
	return me.RandRange(0, _max-1)
}

//get seed
func (me *Rand64) Seed() uint64 {
	return me.seed
}

//set seed
func (me *Rand64) Srand(_seed uint64) uint64 {
	me.seed = _seed
	return me.seed
}
