package stl

//s set of bit for vistit
type BitSet struct {
	d []uint
}

//func (this *BitSet) Bits(start, end uint) *Bits {}

//operator of bits
type Bits struct {
	start, end uint8
	d          *BitSet
}

//func (this *Bits) Get() uint       {}
//func (this *Bits) Set(v uint) uint {}
