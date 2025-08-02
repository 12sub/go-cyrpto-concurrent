package plugins

import (
	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

type GCMPlugin struct{}

func (p GCMPlugin) Encrypt(data []byte, key []byte) (string, error) {
	if err := utils.ValidateKeyLength(key, "gcm"); err != nil {
		return "", err
	}
	return crypto.EncryptAesGcm(data, string(key))
}

func (p GCMPlugin) Decrypt(data string, key []byte) ([]byte, error) {
	if err := utils.ValidateKeyLength(key, "gcm"); err != nil {
		return nil, err
	}
	return  crypto.DecryptAesGcm(data, string(key))
}

func (p GCMPlugin) Name() string {
	return "gcm"
}


func init() {
	utils.RegisterPlugin("gcm", GCMPlugin{})
}