package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// very important in block cipher encryption like AES and CBC
// it encrypts data in fixed size blocks. namely: 16 bytes
func pad(src []byte) []byte {
	// aes.Blocksize is 16 bytes
	// len(src)%aes.BlockSize gives you the remainder
	// subtract remainder from aes.BlockSize to tell you how many
	// bytes you need to add to make full block
	// padding :which utilizes cipher key and
	// initialization vector= aes.BlockSize - len(src)%aes.BlockSize

	// creates slice of padding bytes.
	// this is PKCS#7 padding scheme
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// function for encryption algorithm
func Encrypt(text []byte, key []byte) (string, error) {
	//initiating public cipher key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// including pad func into plaintext
	plaintext := pad(text)

	// creating initialization vector
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// converting plaintext to ciphertext using
	// CBC AES Block encryption algorithm which utilizes cipher key and
	// initialization vector
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// return encoded ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}


