package bitset

const (
	bitsOneElem        = 64
	one         uint64 = 1
)

// NewBitSet create a bitset with at least capacity
func NewBitSet(capacity int) *BitSet {
	n := (capacity + bitsOneElem - 1) / bitsOneElem
	return &BitSet{
		b: make([]uint64, n),
	}
}

// BitSet is a set of bits
type BitSet struct {
	b []uint64
}

// Add try to set a bit with 1
// it return false if this bit has been set.
func (bs *BitSet) Add(bit uint64) bool {
	n, b := bit/bitsOneElem, bit%bitsOneElem
	mask := one << b
	if bs.b[n]&mask != 0 {
		return false
	}
	bs.b[n] |= mask
	return true
}

// Set set a bit as 1
func (bs *BitSet) Set(bit uint64) {
	n, b := bit/bitsOneElem, bit%bitsOneElem
	mask := one << b
	bs.b[n] |= mask
}

// Unset set a bit as 0
func (bs *BitSet) Unset(bit uint64) {
	n, b := bit/bitsOneElem, bit%bitsOneElem
	mask := one << b
	bs.b[n] &= ^mask
}

// Has check if a bit has been set as 1
func (bs BitSet) Has(bit uint64) bool {
	n, b := bit/bitsOneElem, bit%bitsOneElem
	mask := one << b
	return bs.b[n]&mask != 0
}
