package utils

import (
	"fmt"
)

func EncryptString(plainText string, key []byte, scheme string) (string, error) {
	plugin, ok := GetPlugin(scheme)
	if !ok {
		return "", fmt.Errorf("encryption scheme '%s' not supported", scheme)
	}

	cipherText, err := plugin.Encrypt([]byte(plainText), key)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %w", err)
	}
	return cipherText, nil
}

func DecryptString(cipherText string, key []byte, scheme string) ([]byte, error) {
	plugin, ok := GetPlugin(scheme)
	if !ok {
		return nil, fmt.Errorf("decryption scheme '%s' not supported", scheme)
	}

	plain, err := plugin.Decrypt(cipherText, key)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return plain, nil
}
