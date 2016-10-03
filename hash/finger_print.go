package hash

import (
	"bytes"
	"fmt"
	"hash"
)

type Fingerprint struct {
	hs1, hs2, hs3 hash.Hash
	size          int64 //size wrote
}

func NewFingerprint() *Fingerprint {
	fp := &Fingerprint{}
	fp.hs1 = MD5.New()
	fp.hs2 = SHA1.New()
	fp.hs3 = CRC64.New()
	return fp
}
func (this *Fingerprint) Resetsize(size int64) {
	this.size = size
}
func (this *Fingerprint) String() string {
	p := fmt.Sprintf("%012x-%x-%x-%x", this.size&0xFFFFFFFFFFFF, this.hs1.Sum(nil)[4:10], this.hs2.Sum(nil)[7:13], this.hs3.Sum(nil)[1:7])
	return p
}

func (this *Fingerprint) Write(p []byte) (n int, err error) {
	if n, err = this.hs1.Write(p); err != nil {
		return
	}
	if n, err = this.hs2.Write(p); err != nil {
		return
	}
	if n, err = this.hs3.Write(p); err != nil {
		return
	}
	return
}

func (this *Fingerprint) Sum(b []byte) []byte {
	buf := bytes.NewBuffer(nil)
	size := uint64(this.size & 0xFFFFFFFFFFFF)
	for i := uint(6); i > 0; i-- {
		buf.WriteByte(byte(size >> ((i - 1) * 8) & 0xff))
	}
	buf.Write(this.hs1.Sum(nil)[4:10])
	buf.Write(this.hs2.Sum(nil)[7:13])
	buf.Write(this.hs3.Sum(nil)[1:7])
	return buf.Bytes()
}
func (this *Fingerprint) Reset() {
	this.hs1.Reset()
	this.hs2.Reset()
	this.hs3.Reset()
}
func (this *Fingerprint) Size() int {
	return int(this.size)
}
func (this *Fingerprint) BlockSize() int {
	return 24
}
