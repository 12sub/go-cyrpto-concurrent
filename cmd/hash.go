package cmd

import (
	"fmt"

	"example.com/crypto-cli/crypto"
	"github.com/spf13/cobra"
)

// creating variables
var hashInput string
var hashAlgo string
var hashFile string
var hashCompare string

// creating cobra logic
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Generate a hash for both strings and files",
	Run: func(cmd *cobra.Command, args []string) {
		var result string

		if hashFile != "" {
			result, err := crypto.HashFile(hashFile, hashAlgo)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("%s file hash: %s\n", hashAlgo, result)
			return
		}
		if hashInput != "" {
			result, err := crypto.HashString(hashInput, hashAlgo)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("%s hash: %s\n", hashAlgo, result)
			return
		}
		if hashCompare != "" {
		if result == hashCompare {
			fmt.Println("✅ Hash matches!")
		} else {
			fmt.Println("❌ Hash mismatch.")
			fmt.Printf("Expected: %s\nGot:      %s\n", hashCompare, result)
		}
		return
	}
		fmt.Println("Error: You must provide either --input or --file")
	},
}

func init() {
	hashCmd.Flags().StringVar(&hashInput, "input", "", "Input string to hash")
	hashCmd.Flags().StringVar(&hashAlgo, "algo", "sha256", "HAshing Algorithm: sha256, sha512, md5")
	hashCmd.Flags().StringVar(&hashFile, "file", "", "File string to hash")
	hashCmd.Flags().StringVar(&hashCompare, "compare", "", "Compare computed hash with given hash")
}
