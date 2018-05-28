package user

import (
	"testing"

	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/crypto"
)

const testName = "theJulyJasmine"

var keyLabel = []byte("test")

func TestIssuer_IssueCode(t *testing.T) {
	issuer := &Issuer{
		Encryptor: &crypto.RSACrypt{
			Logger:     logrus.New(),
			PubKeyPath: "../crypto/pub_test.pem",
			KeyLabel:   keyLabel,
		},
	}

	_, err := issuer.IssueCode(testName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssuer_Generate2DBarcode(t *testing.T) {
	issuer := &Issuer{
		Encryptor: &crypto.RSACrypt{
			Logger:     logrus.New(),
			PubKeyPath: "../crypto/pub_test.pem",
			KeyLabel:   keyLabel,
		},
	}

	_, err := issuer.Generate2DBarcode(testName)
	if err != nil {
		t.Fatal(err)
	}
}
