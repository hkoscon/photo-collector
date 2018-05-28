package user

import (
	"bytes"
	"encoding/base64"
	"github.com/boombuler/barcode/qr"
	"hkoscon.org/photos/pkg/crypto"
	"hkoscon.org/photos/pkg/modals"
	"image/png"
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

	dst := make([]byte, base64.RawStdEncoding.EncodedLen(len(code)))
	base64.RawStdEncoding.Encode(dst, code)

	return dst, nil
}

func (i *Issuer) Generate2DBarcode(name string) (_ []byte, err error) {
	code, err := i.IssueCode(name)

	img, err := qr.Encode(string(code), qr.M, qr.Auto)
	if err != nil {
		return
	}

	var buffer bytes.Buffer

	if err := png.Encode(&buffer, img); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
