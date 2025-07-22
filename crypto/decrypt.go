package crypto


import (
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
)

// func to unpad padded algorithm
func unpad(src []byte) []byte {
	length := len(src)
	unpad1 := int(src[length-1])
	return src[:(length - unpad1)]  
}

// func for decryption algorithm
func Decrypt(cryptoText string, key []byte) (string, error) {
	// Decoding encrypted string
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	// creating private cipher key 
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// creating initialization vector
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// converting ciphertext to plaintext
	// Decrypting ciphertext using CBCDecrypter
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// return unpadded plaintext
	return string(unpad(ciphertext)), nil

}