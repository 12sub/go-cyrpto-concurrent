package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"Go Encrypter",
	Short:	"A CLI Tool to encrypt/decrypt strings and files",
	Long: 	"Encrypting and Decrypting strings and/or files using AES Encryption",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(hashCmd)
}