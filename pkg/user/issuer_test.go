package user

import (
	"testing"

	"github.com/sirupsen/logrus"
	"hkoscon.org/photos/pkg/crypto"
)

const testName = "theJulyJasmine"

func TestIssuer_Generate2DBarcode(t *testing.T) {
	issuer := &Issuer{
		Encryptor: &crypto.RSACrypt{
			Logger:     logrus.New(),
			PubKeyPath: "../crypto/pub_test.pem",
		},
	}

	if err := issuer.Generate2DBarcode(testName, "qrcode.png"); err != nil {
		t.Fatal(err)
	}
}
