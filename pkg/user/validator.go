package user

import (
	"encoding/base64"
	"hkoscon.org/photos/pkg/crypto"
	"hkoscon.org/photos/pkg/modals"
)

type Validator struct {
	Decryptor crypto.Decryptor `inject:"decryptor"`
}

func (v *Validator) GetName(input string) (name string, err error) {
	code, err := base64.RawStdEncoding.DecodeString(input)
	if err != nil {
		return
	}

	content, err := v.Decryptor.Decrypt(code)
	if err != nil {
		return
	}

	photographer := new(modals.Photographer)
	if err := photographer.Unmarshal(content); err != nil {
		return "", err
	}

	return photographer.GetName(), nil
}
