package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"example.com/crypto-cli/crypto"
	"example.com/crypto-cli/utils"
)

var (
	mode 	string
	input	[]string
	key		string
	inputType	string
	concurrent	bool
)

var runCmd = &cobra.Command{
	Use:	"run",
	Short:	"Run encryption and decryption",
	Run:	func(cmd *cobra.Command, args []string) {
		// creating key through slice bytes
		k := []byte(key)
		if len(k) != 16 {
			fmt.Println("Key must be exactly 16 bytes.")
			return
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
	runCmd.Flags().StringSliceVar(&input, "input", []string{}, "Input strings or file paths")
	runCmd.Flags().StringVar(&key, "key", "1234567890abcdef", "16-byte key")
	runCmd.Flags().StringVar(&inputType, "type", "string", "Type: string or file")
	runCmd.Flags().BoolVar(&concurrent, "concurrent", false, "Enable concurrent file processing")

}

// function to encryption/decryption script logic
func handleString(in string, mode string, key []byte){
	if mode == "encrypt" {
		out, err := crypto.Encrypt([]byte(in), key)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Encrypted: ", out)
	} else {
		out, err := crypto.Decrypt(in, key)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Decrypted: ", out)
	}	
	
}

/// function to encryption/decryption file logic
func handleFile(path string, mode string, key []byte) {
	data, err := utils.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read %s: %v\n", path, err)
		return
	}
	if mode == "encrypt" {
		out, err := crypto.Encrypt(data, key)
		if err != nil {
			fmt.Println("File Encryption Error:", err)
			return
		}
		utils.WriteFile(path+".enc", []byte(out))
		fmt.Println("Encrypted: %s\n", path)
	} else {
		out, err := crypto.Decrypt(string(data), key)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		utils.WriteFile(path+".dec", []byte(out))
		fmt.Println("Decrypted: %s\n", path)
	}
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