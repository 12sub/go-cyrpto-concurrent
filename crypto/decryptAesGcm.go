package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

func DecryptAesGcm(cipherHex string, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key must be 16, 24, or 32 bytes long")
	}

	data, err := hex.DecodeString(cipherHex)
	if err != nil {
		return nil, err
	}

	if len(data) < 12 {
		return nil, errors.New("ciphertext too short")
	}

	nonce := data[:12]
	ciphertext := data[12:]

	// initiating cipher key
	block, err := aes.NewCipher(key)
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
