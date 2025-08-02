package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SaltSize   = 16
	Iterations = 10000
)

var KeyLen = map[string]int{
	"cbc":    16,
	"gcm":    16,
	"chacha": 32,
}

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
func DeriveKeyWithScheme(password string, salt []byte, scheme string) ([]byte, error) {
	length, ok := KeyLen[scheme]
	if !ok {
		return nil, fmt.Errorf("unsupported encryption scheme: %s", scheme)
	}
	return pbkdf2.Key([]byte(password), salt, Iterations, length, sha256.New), nil
}

// Creating a function to encode salt as hex string for CLI Friendly output
func EncodeSalt(salt []byte) string {
	return hex.EncodeToString(salt)
}

// Decoding hex string back to salt
func DecodeSalt(saltHex string) ([]byte, error) {
	return hex.DecodeString(saltHex)
}

func ValidateKeyLength(key []byte, scheme string) error {
	expected, ok := KeyLen[scheme]
	if !ok {
		return fmt.Errorf("unsupported encryption scheme: %s", scheme)
	}
	if len(key) != expected {
		return fmt.Errorf("invalid key length for %s: expected %d bytes, got %d bytes", scheme, expected, len(key))
	}
	return nil
}
