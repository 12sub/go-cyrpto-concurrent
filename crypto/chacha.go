package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

// encryption with chacha20: encrypts using ChaCha20-Poly1305
func EncryptChaCha20(plain []byte, key []byte) (string, error) {
	// implementing chacha20poly1305 encryption algorithm
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return "", err
	}
	// creating nonce of chacha20poly1305
	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// encrypting the text to generate the ciphertext
	// and return the encoding of the ciphertext
	ciphertext := aead.Seal(nonce, nonce, plain, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// decryption with chacha20: decrypt using ChaCha20-Poly1305
func DecryptChaCha20(enc string, key []byte) ([]byte, error) {
	// decoding the ciphertext
	data, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return nil, err
	}

	// creating a new chacha20poly1305 key
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	// removing nonce size
	if len(data) < chacha20poly1305.NonceSize {
		return nil, errors.New("invalid data size")
	}

	nonce := data[:chacha20poly1305.NonceSize]
	plaintext := data[chacha20poly1305.NonceSize:]

	// return the decrypted file/string
	return aead.Open(nil, nonce, plaintext, nil)
}
