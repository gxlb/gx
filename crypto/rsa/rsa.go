package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const (
	defaultHash = crypto.SHA512
)

var (
	errPublicKey  = errors.New("public key error")
	errPrivateKey = errors.New("private key error")
)

func MustNewClient(publicKey []byte) *RsaClient {
	c, err := NewClient(publicKey)
	if err != nil {
		panic(err)
	}
	return c
}

func NewClient(publicKey []byte) (*RsaClient, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errPublicKey
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	c := &RsaClient{publicKey: pub}
	return c, nil
}

type RsaClient struct {
	publicKey *rsa.PublicKey
}

func (this *RsaClient) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, plaintext)
}

func (this *RsaClient) VerifyDefault(src []byte, sign []byte) error {
	return this.Verify(src, sign, defaultHash)
}

func (this *RsaClient) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
}

func MustNewServer(privateKey []byte) *RsaServer {
	s, err := NewServer(privateKey)
	if err != nil {
		panic(err)
	}
	return s
}

func NewServer(privateKey []byte) (*RsaServer, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errPrivateKey
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	s := &RsaServer{privateKey: priv}
	return s, nil
}

type RsaServer struct {
	privateKey *rsa.PrivateKey
}

func (this *RsaServer) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, ciphertext)
}

func (this *RsaServer) SignDefault(src []byte) ([]byte, error) {
	return this.Sign(src, defaultHash)
}

func (this *RsaServer) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
}

func GenRsaKeys(bits int) (privateKey, publicKey string, err error) {
	//generate private key
	priKey, _err := rsa.GenerateKey(rand.Reader, bits)
	if _err != nil {
		return "", "", _err
	}
	derStream := x509.MarshalPKCS1PrivateKey(priKey)
	block := &pem.Block{
		Type:  "private key",
		Bytes: derStream,
	}

	var bufPri bytes.Buffer
	err = pem.Encode(&bufPri, block)
	if err != nil {
		return "", "", err
	}
	privateKey = bufPri.String()

	// generate public key from private key
	pubKey := &priKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "public key",
		Bytes: derPkix,
	}
	var bufPub bytes.Buffer
	err = pem.Encode(&bufPub, block)
	if err != nil {
		return "", "", err
	}
	publicKey = bufPub.String()
	return
}
