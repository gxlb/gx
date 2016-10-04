package hash

import (
	"bytes"
	"hash"
	"io"
)

var (
	gId uint32
)

type NoneHash struct {
	id uint32
}

func NewNoneHash() hash.Hash {
	gId++
	return &NoneHash{id: gId}
}

func (this *NoneHash) ReadFrom(r io.Reader) (n int64, err error) {
	return
}

func (this *NoneHash) Write(p []byte) (n int, err error) {
	return
}

func (this *NoneHash) Sum(b []byte) []byte {
	buf := bytes.NewBuffer(nil)
	id := this.id
	for i := uint(4); i > 0; i-- {
		buf.WriteByte(byte(id >> ((i - 1) * 8) & 0xff))
	}
	return buf.Bytes()
}

func (this *NoneHash) Reset() {
}

func (this *NoneHash) Size() int {
	return 0
}

func (this *NoneHash) BlockSize() int {
	return 0
}
