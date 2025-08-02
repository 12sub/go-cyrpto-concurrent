package plugins

import (
	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

type CBCPlugin struct{}

// making CBC and GCM into plugins
// first for encryption 
func (p CBCPlugin) Encrypt(data []byte, key []byte) (string, error) {
	if err := utils.ValidateKeyLength(key, "cbc"); err != nil {
		return "", err
	}
	return crypto.Encrypt(data, key)
}

func (p CBCPlugin) Decrypt(data string, key []byte) ([]byte, error) {
	if err := utils.ValidateKeyLength(key, "cbc"); err != nil {
		return nil, err
	}
	plain, err := crypto.Decrypt(data, key)
	return []byte(plain), err
}

func (p CBCPlugin) Name() string {
	return "cbc"
}

func init() {
	utils.RegisterPlugin("cbc", CBCPlugin{})
}