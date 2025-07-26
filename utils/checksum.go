package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// creating a function to compute the checksum of original input using sha256
func ComputeSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// creating func to write checksum files
// return error value
// pass in path and string of checksum
func WriteChecksumFile(path, checksum string) error {
	return os.WriteFile(path+".sha256", []byte(checksum), 0644)
}

// creating func to read checksum files
// return string value and error value
// pass in string of path
func ReadChecksumFile(path string) (string, error) {
	data, err := os.ReadFile(path + ".sha256")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
