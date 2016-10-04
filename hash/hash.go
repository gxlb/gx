//package hash uniform the interface to create a hash objet
package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/crc32"
	"hash/crc64"
)

type HashMethod uint

const (
	NONE HashMethod = iota
	MD5
	SHA1
	SHA256
	SHA512
	CRC32
	CRC64
	FINGERPRINT //finger print, use md5+sha1+crc64
)

var (
	gCrc6ISO4Table = crc64.MakeTable(crc64.ISO)
)

//create a hash object
func (me HashMethod) New() (h hash.Hash) {
	switch me {
	case MD5:
		h = md5.New()
	case SHA1:
		h = sha1.New()
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	case CRC32:
		h = crc32.NewIEEE()
	case CRC64:
		h = crc64.New(gCrc6ISO4Table)
	case FINGERPRINT:
		h = NewFingerprint()
	case NONE:
		h = NewNoneHash()
	default:
		panic(me)
	}
	return
}

//create a hash object
func New(t HashMethod) hash.Hash {
	return t.New()
}
