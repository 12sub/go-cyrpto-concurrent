package plugins

// this is the chacha20-poly1305 plugin architecture implementation
import (
	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

type ChaChaPlugin struct{}

func (p ChaChaPlugin) Encrypt(data []byte, key []byte) (string, error) {
	if err := utils.ValidateKeyLength(key, "chacha"); err != nil {
		return "", err
	}
	return crypto.EncryptChaCha20(data, key)
}

func (p ChaChaPlugin) Decrypt(data string, key []byte) ([]byte, error) {
	if err := utils.ValidateKeyLength(key, "chacha"); err != nil {
		return nil, err
	}
	return crypto.DecryptChaCha20(data, key)
}

func (p ChaChaPlugin) Name() string {
	return "chacha"
}

func init() {
	utils.RegisterPlugin("chacha", ChaChaPlugin{})
}
