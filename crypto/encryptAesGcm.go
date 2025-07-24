package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func EncryptAesGcm(plaintext []byte, key []byte) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key must be 16, 24, or 32 bytes long")
	}

	// initiating cipher key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// creating a new GCM Block
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// Generating random nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// Implementing the AES Encryption Algorithm to create the ciphertext
	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)

	// prepared nonce 2 ciphertext
	final := append(nonce, ciphertext...)

	return hex.EncodeToString(final), nil
}