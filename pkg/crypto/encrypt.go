package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type Encrpytor interface {
	Encrypt(data []byte) ([]byte, error)
}

func (c *RSACrypt) loadPublicKey() {
	c.Logger.Debugln("Load Public Key")
	keyBlocks, err := ioutil.ReadFile(c.PubKeyPath)
	if err != nil {
		c.Logger.Error(err)
		panic(err)
	}

	block, _ := pem.Decode(keyBlocks)

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		c.Logger.Error(err)
		panic(err)
	}

	switch pub := pubKey.(type) {
	case *rsa.PublicKey:
		c.publicKey = pub
	default:
		panic("unknown type of public key")
	}
}

func (c *RSACrypt) Encrypt(data []byte) ([]byte, error) {
	if c.publicKey == nil {
		c.loadPublicKey()
	}
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, c.publicKey, data, c.KeyLabel)
}
