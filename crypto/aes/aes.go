package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7Padding(b []byte, blockSize int) []byte {
	size := blockSize - len(b)%blockSize
	buf := make([]byte, len(b)+size)
	copy(buf[:len(b)], b)
	for i := len(b); i < len(buf); i++ {
		buf[i] = byte(size)
	}
	return buf
}

func PKCS7PaddingString(str string, blockSize int) []byte {
	size := blockSize - len(str)%blockSize
	buf := make([]byte, len(str)+size)
	copy(buf[:len(str)], str)
	for i := len(str); i < len(buf); i++ {
		buf[i] = byte(size)
	}
	return buf
}

func PKCS7Unpadding(b []byte) []byte {
	size := int(b[len(b)-1])
	return b[:len(b)-size]
}

type AESEncrypter struct {
	blockMode cipher.BlockMode
	blockSize int
}

// CBC模式/PKCS7Padding
func NewAESEncrypter(key []byte) (*AESEncrypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	return &AESEncrypter{
		blockMode: blockMode,
		blockSize: blockSize,
	}, nil
}

func (this *AESEncrypter) Encrypt(b []byte) []byte {
	padded := PKCS7Padding(b, this.blockSize)
	encrypted := make([]byte, len(padded))
	this.blockMode.CryptBlocks(encrypted, padded)
	return encrypted
}

func (this *AESEncrypter) EncryptString(str string) []byte {
	padded := PKCS7PaddingString(str, this.blockSize)
	encrypted := make([]byte, len(padded))
	this.blockMode.CryptBlocks(encrypted, padded)
	return encrypted
}

type AESDecrypter struct {
	blockMode cipher.BlockMode
	blockSize int
}

// CBC模式/PKCS7Padding
func NewAESDecrypter(key []byte) (*AESDecrypter, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	return &AESDecrypter{
		blockMode: blockMode,
		blockSize: blockSize,
	}, nil
}

func (this *AESDecrypter) Decrypt(b []byte) []byte {
	decrypted := make([]byte, len(b))
	this.blockMode.CryptBlocks(decrypted, b)
	return PKCS7Unpadding(decrypted)
}

func (this *AESDecrypter) DecryptToString(b []byte) string {
	return string(this.Decrypt(b))
}
