//Package rand implements some useful rand object
package rand

import (
	"sync/atomic"
)

var (
	gRand<BITS> = NewRand<BITS>S(<SEED_TYPE>(RandSeed(0)))
)

//generate a rand number
func Rand<BITS>() <SEED_TYPE> {
	return gRand<BITS>.Rand()
}

//generate a rand number less than max
func RandMax<BITS>(max <SEED_TYPE>) <SEED_TYPE> {
	return gRand<BITS>.RandMax(max)
}
func RandRange<BITS>(min, max <SEED_TYPE>) <SEED_TYPE> {
	return gRand<BITS>.RandRange(min, max)
}

//rand number generator
//It is thread safe
type Rand<BITS>T struct {
	seed <SEED_TYPE>
	//lock sync.Mutex
}

//new a initialized rand<BITS> object
func NewRand<BITS>S(seed <SEED_TYPE>) *Rand<BITS>T {
	return &Rand<BITS>T{seed: seed}
}

//new a rand<BITS> object initialized by auto-generated seed
func NewRand<BITS>() *Rand<BITS>T {
	return NewRand<BITS>S(gRand<BITS>.randBase())
}

//next rand number
func (me *Rand<BITS>T) Rand() <SEED_TYPE> {
	var o, n <SEED_TYPE>
	for { //mutithread lock-free operation
		o = atomic.LoadUint<BITS>(&me.seed)
		n = o*<PRIME_A> + <PRIME_C>
		if atomic.CompareAndSwapUint<BITS>(&me.seed, o, n) {
			break
		}
	}
	return n

	//me.seed = me.seed*g_prime_a<BITS> + g_prime_c<BITS>
	//return me.seed
}

//new rand seed list
func (me *Rand<BITS>T) randBase() <SEED_TYPE> {
	return <SEED_TYPE>(RandSeed(uint64(me.Rand())))
}

//generate rand number in range
func (me *Rand<BITS>T) RandRange(min, max <SEED_TYPE>) <SEED_TYPE> {
	if max < min {
		max, min = min, max
	}
	d := max - min + 1
	r := me.Rand()
	ret := r%d + min

	return ret
}

//generate rand number with max value
func (me *Rand<BITS>T) RandMax(max <SEED_TYPE>) <SEED_TYPE> {
	return me.RandRange(0, max)
}

//get seed
func (me *Rand<BITS>T) Seed() <SEED_TYPE> {
	return atomic.LoadUint<BITS>(&me.seed)
}

//set seed
func (me *Rand<BITS>T) Srand(_seed <SEED_TYPE>) <SEED_TYPE> {
	ret := atomic.SwapUint<BITS>(&me.seed, _seed) //mutithread lock-free operation
	return ret
}
