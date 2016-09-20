//Package rand implements some useful rand object
package rand

import (
	"sync/atomic"
	"time"
	ctime "vipally.gmail.com/basic/c/time"
)

const (
	g_prime_a32 = 7368787
	g_prime_c32 = 2750159

	g_prime1_a32 = 8615693
	g_prime1_c32 = 3894211
)

var (
	g_rand32 = NewRand32S(uint32(RandSeed()))
)

func RandSeed() uint64 {
	r := ctime.RDTSC() + uint64(time.Now().Unix())
	return r
}

func Rand32() uint32 {
	return g_rand32.Rand()
}
func RandMax32(_max uint32) uint32 {
	return g_rand32.RandMax(_max)
}
func RandRange32(_min, _max uint32) uint32 {
	return g_rand32.RandRange(_min, _max)
}

//32-bits rand number generator
//It is thread safe
type Rand32T struct {
	seed uint32
	//lock sync.Mutex
}

//new a initialized rand32 object
func NewRand32S(_seed uint32) *Rand32T {
	return &Rand32T{seed: _seed}
}

//new a rand32 object initialized by auto-generated seed
func NewRand32() *Rand32T {
	return NewRand32S(g_rand32.randBase())
}

//next rand number
func (me *Rand32T) Rand() uint32 {
	var old, new uint32
	for { //mutithread lock-free operation
		old = atomic.LoadUint32(&me.seed)
		new = old*g_prime_a32 + g_prime_c32
		if atomic.CompareAndSwapUint32(&me.seed, old, new) {
			break
		}
	}
	return new

	//me.seed = me.seed*g_prime_a32 + g_prime_c32
	//return me.seed
}

//new rand seed list
func (me *Rand32T) randBase() uint32 {
	var old, new uint32
	for { //mutithread lock-free operation
		old = atomic.LoadUint32(&me.seed)
		new = old*g_prime1_a32 + g_prime1_c32
		if atomic.CompareAndSwapUint32(&me.seed, old, new) {
			break
		}
	}
	return new
}

//generate rand number in range
func (me *Rand32T) RandRange(_min, _max uint32) uint32 {
	if _max < _min {
		_max, _min = _min, _max
	}
	d := _max - _min + 1
	r := me.Rand()
	ret := r%d + _min

	return ret
}

//generate rand number with max
func (me *Rand32T) RandMax(_max uint32) uint32 {
	return me.RandRange(0, _max-1)
}

//get seed
func (me *Rand32T) Seed() uint32 {
	return atomic.LoadUint32(&me.seed)
}

//set seed
func (me *Rand32T) Srand(_seed uint32) uint32 {
	ret := atomic.SwapUint32(&me.seed, _seed) //mutithread lock-free operation
	return ret
}
