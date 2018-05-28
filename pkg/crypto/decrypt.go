package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type Decryptor interface {
	Decrypt()
}

func (c *RSACrypt) loadKey(password []byte) {
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
