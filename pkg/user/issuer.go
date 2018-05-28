package user

import (
	"encoding/base64"
	"image/png"
	"os"

	"github.com/boombuler/barcode/qr"
	"hkoscon.org/photos/pkg/crypto"
	"hkoscon.org/photos/pkg/modals"
)

type Issuer struct {
	Encryptor crypto.Encrpytor `inject:"crypt"`
}

func (i *Issuer) IssueCode(name string) (_ []byte, err error) {
	photographer := &modals.Photographer{
		Name: name,
	}

	encoded, err := photographer.Marshal()
	if err != nil {
		return
	}

	code, err := i.Encryptor.Encrypt(encoded)
	if err != nil {
		return
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(code)))
	base64.StdEncoding.Encode(dst, code)
	return dst, nil
}

func (i *Issuer) Generate2DBarcode(name, filename string) (err error) {
	code, err := i.IssueCode(name)

	img, err := qr.Encode(string(code), qr.M, qr.Auto)
	if err != nil {
		return
	}

	output, err := os.Create(filename)
	if err != nil {
		return
	}
	defer output.Close()

	return png.Encode(output, img)
}
