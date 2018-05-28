package crypto

import (
	"crypto/rsa"

	"github.com/sirupsen/logrus"
)

type RSACrypt struct {
	privateKey *rsa.PrivateKey
	KeyPath    string

	publicKey  *rsa.PublicKey
	PubKeyPath string

	keyLabel []byte

	Logger logrus.FieldLogger `inject:"crypt logger"`
}
