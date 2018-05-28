package crypto

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCrypt_loadPublicKey(t *testing.T) {
	crypt := &RSACrypt{
		PubKeyPath: "pub_test.pem",
		Logger:     logrus.New(),
	}
	crypt.loadPublicKey()
	if crypt.publicKey == nil {
		t.Fatal("fail to load public key")
	}
}

func TestRSACrypt_Encrypt(t *testing.T) {
	crypt := &RSACrypt{
		PubKeyPath: "pub_test.pem",
		Logger:     logrus.New(),
		KeyLabel:   []byte("test"),
	}

	_, err := crypt.Encrypt([]byte("test"))
	if err != nil {
		t.Fatal(err)
	}
}
