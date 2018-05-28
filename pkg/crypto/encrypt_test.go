package crypto

import "testing"

func TestCrypt_loadPublicKey(t *testing.T) {
	crypt := &RSACrypt{
		PubKeyPath: "pub_test.pem",
	}
	crypt.loadPublicKey()
	if crypt.publicKey == nil {
		t.Fatal("fail to load public key")
	}
}
