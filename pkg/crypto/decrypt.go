package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

var ErrNoKey = errors.New("no private key")

type Decryptor interface {
	Decrypt(chipper []byte) (_ []byte, err error)
}

func (c *RSACrypt) LoadKey(password []byte) {
	keyBlocks, err := ioutil.ReadFile(c.PriKeyPath)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(keyBlocks)

	rawKey, err := x509.DecryptPEMBlock(block, password)
	if err != nil {
		panic(err)
	}

	c.privateKey, err = x509.ParsePKCS1PrivateKey(rawKey)
	if err != nil {
		panic(err)
	}
}

func (c *RSACrypt) Decrypt(chipper []byte) (_ []byte, err error) {
	if c.privateKey == nil || c.KeyLabel == nil {
		return nil, ErrNoKey
	}

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, c.privateKey, chipper, c.KeyLabel)
}
