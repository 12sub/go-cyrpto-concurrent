package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize   = 16
	Iterations = 10000
	KeyLen     = 16
)

// Creating random salt
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, SaltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// creating function to derive the key using pbkdf2 2 create a key from password + salt
func DeriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, Iterations, KeyLen, sha256.New)
}

// Creating a function to encode salt as hex string for CLI Friendly output
func EncodeSalt(salt []byte) string {
	return hex.EncodeToString(salt)
}

// Decoding hex string back to salt
func DecodeSalt(saltHex string) ([]byte, error) {
	return hex.DecodeString(saltHex)
}
