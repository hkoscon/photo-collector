package crypto

import (
	"crypto/rsa"

	"github.com/sirupsen/logrus"
)

type RSACrypt struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey

	KeyLabel   []byte `inject:"key label"`
	PubKeyPath string `inject:"public key path"`
	PriKeyPath string `inject:"private key path"`

	Logger logrus.FieldLogger `inject:"crypt logger"`
}
