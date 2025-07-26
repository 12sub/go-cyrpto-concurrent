package cmd

import (
	"fmt"
	"sync"

	"example.com/crypto-cli/utils"
	"github.com/spf13/cobra"
)

var (
	mode       string
	input      []string
	key        string
	inputType  string
	concurrent bool
	scheme     string
	password   string
	salt       string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run encryption and decryption",
	Run: func(cmd *cobra.Command, args []string) {
		// creating key through slice bytes
		var k []byte
		if password != "" {
			var s []byte
			var err error
			if salt == "" {
				s, err := utils.GenerateSalt()
				if err != nil {
					utils.Error("Error generating salt: %v", err)
					return
				}
				utils.Info("Generated Salt (Save this for decryption): %s", utils.EncodeSalt(s))
			} else {
				s, err = utils.DecodeSalt(salt)
				if err != nil {
					utils.Error("Invalid salt: %v", err)
					return
				}
			}
			k = utils.DeriveKey(password, s)
		} else {
			k = []byte(key)
			if len(k) != 16 {
				fmt.Println("Key must be exactly 16 bytes.")
				return
			}
		}

		if inputType == "string" {
			for _, in := range input {
				handleString(in, mode, k)
			}
		} else {
			if concurrent {
				handleFilesConcurrently(input, mode, k)
			} else {
				for _, file := range input {
					handleFile(file, mode, k)
				}
			}
		}
	},
}

// flags for encryption / decryption of files, strings and a single file
func init() {
	runCmd.Flags().StringVar(&mode, "mode", "encrypt", "Mode: encrypt or decrypt")
	runCmd.Flags().StringVar(&scheme, "scheme", "cbc", "Encryption scheme: cbc or gcm")
	runCmd.Flags().StringSliceVar(&input, "input", []string{}, "Input strings or file paths")
	runCmd.Flags().StringVar(&key, "key", "1234567890abcdef", "16-byte key")
	runCmd.Flags().StringVar(&inputType, "type", "string", "Type: string or file")
	runCmd.Flags().StringVar(&password, "password", "", "Password to derive key using PBKDF2")
	runCmd.Flags().StringVar(&salt, "salt", "", "Hex-encoded salt for PBKDF2 (optional for decryption)")
	runCmd.Flags().BoolVar(&concurrent, "concurrent", false, "Enable concurrent file processing")

}

// function to encryption/decryption script logic
func handleString(in string, mode string, key []byte) {
	var out string
	var err error

	// switch scheme {
	// case "cbc":
	// 	if mode == "encrypt" {
	// 		out, err = crypto.Encrypt([]byte(in), key)
	// 	} else {
	// 		out, err = crypto.Decrypt(in, key)
	// 	}
	// case "gcm":
	// 	if mode == "encrypt" {
	// 		out, err = crypto.EncryptAesGcm([]byte(in), string(key))
	// 	} else {
	// 		plain, err := crypto.DecryptAesGcm(in, string(key))
	// 		if err == nil {
	// 			out = string(plain)
	// 		}
	// 	}
	// default:
	// 	fmt.Println("Unsupported scheme:", scheme)
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	plugin, ok := utils.GetPlugin(scheme)
	if !ok {
		fmt.Println("Unsupported scheme:", scheme)
		return
	}

	if mode == "encrypt" {
		out, err = plugin.Encrypt([]byte(in), key)
		if err != nil {
			fmt.Println("Error encrypting:", err)
			return
		}
	} else {
		var plain []byte
		plain, err = plugin.Decrypt(in, key)
		if err == nil {
			out = string(plain)
		}
	}
	if mode == "encrypt" {
		fmt.Println("Encrypted: ", out)
	} else {
		fmt.Println("Decrypted: ", out)
	}

}

// / function to encryption/decryption file logic
func handleFile(path string, mode string, key []byte) {

	data, err := utils.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", path, err)
		return
	}
	var out string
	plugin, ok := utils.GetPlugin(scheme)
	if !ok {
		fmt.Println("Unsupported scheme:", scheme)
		return
	}
	if mode == "encrypt" {
		out, err = plugin.Encrypt(data, key)
		if err == nil {
			checksum := utils.ComputeSHA256(data)
			utils.WriteChecksumFile(path, checksum)
			fmt.Println("SHA256", checksum)
			return
		}
	} else {
		var plain []byte
		plain, err := plugin.Decrypt(string(data), key)
		if err == nil {
			original := []byte(plain)
			newChecksum := utils.ComputeSHA256(original)
			oldChecksum, err := utils.ReadChecksumFile(path)
			// checking to see if the new checksum is the same as the old one
			if err == nil && newChecksum != oldChecksum {
				fmt.Println("WARNING: Decrypted output checksum mismatch! file match not found")
			} else {
				fmt.Println("DECRYPTION SUCCESSFUL: Decrypted output matches the original checksum")
			}
		}
	}
	extension := ".enc"
	if mode == "decrypt" {
		extension = ".dec"
	}
	utils.WriteFile(path+extension, []byte(out))
	fmt.Printf("%s: %s\n", mode, path)
}

// func for encryption/ decryption of multiple files concurrently
// using waitGroup for concurrency
func handleFilesConcurrently(paths []string, mode string, key []byte) {
	var wg sync.WaitGroup
	for _, file := range paths {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			handleFile(f, mode, key)
		}(file)
	}
	wg.Wait()
}
