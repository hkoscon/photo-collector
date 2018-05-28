package crypto

import (
	"crypto/x509"
	"io/ioutil"
	"testing"
)

func TestCrypt_loadKey(t *testing.T) {
	crypt := &RSACrypt{
		KeyPath: "pri_test.pem",
	}

	crypt.loadKey([]byte("1234"))
	if crypt.privateKey == nil {
		t.Fatal("fail to load private key")
	}

	keyData, _ := ioutil.ReadFile("pub_test.pem")

	pubKey, err := x509.ParsePKCS1PublicKey(keyData)
	if err != nil {
		t.Fatal(err)
	}

	if pubKey.E != crypt.privateKey.PublicKey.E {
		t.Fatal("different E")
	}
}
