package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"os"
)

// Creating a func to hash strings hash logic
func HashString(input string, algo string) (string, error) {
	// creating variable hash with a type []byte
	var hash []byte

	// implementing switch to select between hashing algorithms
	switch algo {
	case "sha256":
		sum := sha256.Sum256([]byte(input))
		hash = sum[:]
	case "sha512":
		sum := sha512.Sum512([]byte(input))
		hash = sum[:]
	case "md5":
		sum := md5.Sum([]byte(input))
		hash = sum[:]
	default:
		return "", errors.New("unsupported algorithm: choose sha256, sha512, or md5")
	}
	return hex.EncodeToString(hash), nil
}

func HashFile(filepath string, algo string) (string, error) {
	// creating variable hash with a type hash.hash
	var h hash.Hash

	// implementing switch to select between hashing algorithms
	switch algo {
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	case "md5":
		h = md5.New()
	default:
		return "", errors.New("unsupported algorithm: choose sha256, sha512, or md5")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
