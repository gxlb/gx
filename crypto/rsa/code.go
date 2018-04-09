package rsa

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"gitee.com/vipally/gx/crypto/ecb"
)

var key = []byte("ac643718-9730-e711-59210")

func Encode(src []byte) string {
	bin, err := ecb.TripleDesEncrypt(src, key)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bin)
}

func Decode(src string) []byte {
	c, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil
	}
	content, err := ecb.TripleDesDecrypt(c, key)
	if err != nil {
		return nil
	}
	return content
}

func Hash(src []byte) []byte {
	h := sha512.New()
	for i := 0; i < 5; i++ {
		h.Write(key)
		h.Write(src)
		h.Write(key)
	}
	return h.Sum(nil)
}

func EncodeString(src string) string {
	bin, err := ecb.TripleDesEncrypt([]byte(src), key)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bin)
}

func DecodeString(src string) string {
	c, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	content, err := ecb.TripleDesDecrypt(c, key)
	if err != nil {
		return ""
	}
	return string(content)
}

func HashString(str string) string {
	h := sha512.New()
	for i := 0; i < 5; i++ {
		h.Write(key)
		h.Write([]byte(str))
		h.Write(key)
	}
	return BytesToHexString(h.Sum(nil))
}

func HexStringToBytes(s string) []byte {
	if b, err := hex.DecodeString(s); err == nil {
		return b
	}
	return nil
}
func BytesToHexString(b []byte) string {
	return hex.EncodeToString(b)
}
