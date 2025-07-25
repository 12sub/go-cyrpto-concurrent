package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"example.com/crypto-cli/utils"
)

func EncryptAesGcm(plaintext []byte, password string) (string, error) {
	// creating salt for hashing
	salt, err := utils.GenerateSalt()
	if err != nil {
		return "", err
	}
	// implementing salt for key generation
	derivedKey := utils.DeriveKey(password, salt)

	// initiating cipher key
	block, err := aes.NewCipher(derivedKey)
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
	// combining salt + nonce + ciphertext and return base64
	final := append(salt, nonce...)
	final = append(final, ciphertext...)

	return base64.StdEncoding.EncodeToString(final), nil
}
