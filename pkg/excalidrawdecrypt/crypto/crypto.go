package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Decrypt(buffer, iv []byte, keyString string) ([]byte, error) {

	rawKey, err := base64.RawStdEncoding.DecodeString(keyString)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(rawKey)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGCM.Open(nil, iv, buffer, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
