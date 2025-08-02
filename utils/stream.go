package utils

// implementing streaming mode for large files
// streaming mode improves memory efficiently by processing large files in chunks rather
// than reading them entirely into file

// in this code: we are going to
// 1. open file for reading
// 2. Creating output file for writing
// encrypt or decrypt data chunk by chunk
// support both AES-CBC and AES-GCM modes

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"io"
	"os"
)

func EncryptStreamToWriter(filePath string, mode cipher.BlockMode, out io.Writer) error {
	in, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer in.Close()
	return EncryptStream(in, out, mode)
}

// function to encrypt a file stream using the provided cipher stream
func EncryptStream(in io.Reader, out io.Writer, mode cipher.BlockMode) error {
	
	// making a buffer size in 16 bytes and reading through lines of a file
	buf := make([]byte, 1024)
	for {
		n, err := in.Read(buf)
		if err != nil &&  err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		block := buf[:n]
		if n < 1024 {
			block = PKCS7Pad(block, mode.BlockSize())
		}
		enc := make([]byte, len(block))
		mode.CryptBlocks(enc, block)
		if _, err := out.Write(enc); err != nil {
			return err
		}
		if err == io.EOF {
			break
		}
	}
	return nil

}

// function to decrypt a file stream using the provided cipher stream
func DecryptStream(in io.Reader, out io.Writer, mode cipher.BlockMode) error {
	blockSize := mode.BlockSize()
	buffer := make([]byte, blockSize)

	var prevBlock []byte
	for {
		_, err := io.ReadFull(in, buffer)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		block := make([]byte, blockSize)
		mode.CryptBlocks(block, buffer)

		// Don't write the last block immediately – it might contain padding
		if prevBlock != nil {
			if _, err := out.Write(prevBlock); err != nil {
				return err
			}
		}
		prevBlock = block
	}

	if prevBlock != nil {
		unpadded, err := pkcs7Unpad(prevBlock, blockSize)
		if err != nil {
			return err
		}
		if _, err := out.Write(unpadded); err != nil {
			return err
		}
	}

	return nil
}


// function decryptstreamfrom reader decrypts from reader and writes to writer
func DecryptStreamFromReader(in io.Reader,  out io.Writer, mode cipher.BlockMode) error {
	return DecryptStream(in, out, mode)
}

// adding pkcs7 padding
// this is a padding scheme used in block cipher encryption algorithm like AES in CBC mode
func PKCS7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid padding size")
	}
	padLen := int(data[len(data) - 1])
	if padLen > blockSize || padLen == 0 {
		return nil, errors.New("invalid padding")
	}
	for _, p := range data[len(data)-padLen:] {
		if int(p) != padLen {
			return nil, errors.New("invalid padding pattern")
		}
	}
	return data[:len(data)-padLen], nil
}