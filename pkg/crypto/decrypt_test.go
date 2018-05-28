package crypto

import (
	"testing"
)

func TestCrypt_loadKey(t *testing.T) {
	crypt := &RSACrypt{
		PriKeyPath: "pri_test.pem",
	}

	crypt.LoadKey([]byte("1234"))
	if crypt.privateKey == nil {
		t.Fatal("fail to load private key")
	}
}
