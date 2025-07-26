package plugins

import (
	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

type CBCPlugin struct{}

// making CBC and GCM into plugins
// first for encryption 
func (p CBCPlugin) Encrypt(data []byte, key []byte) (string, error) {
	return crypto.Encrypt(data, key)
}

func (p CBCPlugin) Decrypt(data string, key []byte) ([]byte, error) {
	plain, err := crypto.Decrypt(data, key)
	return []byte(plain), err
}

func init() {
	utils.RegisterPlugin("cbc", CBCPlugin{})
}