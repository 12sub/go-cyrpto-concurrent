package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"example.com/crypto-cli/utils"
)

func DecryptAesGcm(cipherHex string, password string) ([]byte, error) {
	// decoding the strings
	data, err := base64.StdEncoding.DecodeString(cipherHex)
	if err != nil {
		return nil, err
	}
	// Extract the salt, nonce and get ciphertext
	// note: salt (first 16 bytes), nonce (next 12 bytes), rest is ciphertext

	salt := data[:utils.SaltSize]
	nonce := data[utils.SaltSize:utils.SaltSize+12]
	ciphertext := data[utils.SaltSize+12:]

	// getting your derived key
	derivedKey := utils.DeriveKey(password, salt)

	// initiating cipher key
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return nil, err
	}
	// creating a new GCM Block
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// creating nonce size and calculating its length from the original nonce size
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext is too short")
	}

	// decrypting the ciphertest to its original file
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
