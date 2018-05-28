package user

import (
	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/crypto"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	crypt := &crypto.RSACrypt{
		Logger:     logrus.New(),
		PubKeyPath: "../crypto/pub_test.pem",
		PriKeyPath: "../crypto/pri_test.pem",
		KeyLabel:   keyLabel,
	}

	crypt.LoadKey([]byte("1234"))

	issuer := &Issuer{
		Encryptor: crypt,
	}

	code, err := issuer.IssueCode(testName)
	if err != nil {
		t.Fatal(err)
	}

	validator := &Validator{
		Decryptor: crypt,
	}

	name, err := validator.GetName(string(code))
	if err != nil {
		t.Fatal(err)
	}

	if name != testName {
		t.Fatalf("incorrect name %s", name)
	}
}
