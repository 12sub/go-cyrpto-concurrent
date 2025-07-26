package plugins

import (
	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

type GCMPlugin struct{}

func (p GCMPlugin) Encrypt(data []byte, key []byte) (string, error) {
	return crypto.EncryptAesGcm(data, string(key))
}

func (p GCMPlugin) Decrypt(data string, key []byte) ([]byte, error) {
	return  crypto.DecryptAesGcm(data, string(key))
}

func init() {
	utils.RegisterPlugin("gcm", GCMPlugin{})
}