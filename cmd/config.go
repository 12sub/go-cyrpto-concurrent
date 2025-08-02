package cmd

import (
	"fmt"
	"log"
	"os"

	// "log"

	// "example.com/crypto-cli/utils"
	"example.com/crypto-cli/internal/config"
	"example.com/crypto-cli/utils"
	"github.com/spf13/cobra"
)

var configFile string
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Run encryption/decryption using a config file",
	Run: func(cmd *cobra.Command, args []string) {
		// if configFile == "" {
		// 	fmt.Println("No config file provided.")
		// 	return
		// }
		// cfg, err := config.LoadConfig(configFile)
		// if err != nil {
		// 	fmt.Printf("Failed to read %s: %v\n", configFile, err)
		// 	return
		// }
		// fmt.Println("Config file loaded:", cfg)
		cfgPath, _ := cmd.Flags().GetString("file")

		cfg, err := config.LoadConfig(cfgPath)
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		// Decode salt and derive key
		salt, err := utils.DecodeSalt(cfg.Salt)
		if err != nil {
			fmt.Println("Invalid salt in config:", err)
			return
		}
		key, err := utils.DeriveKeyWithScheme(cfg.DefaultPassword, salt, cfg.DefaultScheme)
		if err != nil {
			log.Fatalf("Key derivation failed: %v", err)
		}

		// ✅ Validate key length
		if err := utils.ValidateKeyLength(key, cfg.DefaultScheme); err != nil {
			fmt.Println("Key validation error:", err)
			return
		}

		switch cfg.FileTask.Mode {
		case "encrypt":
			cipher, err := utils.EncryptString(cfg.Input, key, cfg.DefaultScheme)
			if err != nil {
				log.Fatalf("Encryption failed: %v", err)
			}
			if err := os.WriteFile(cfg.Output, []byte(cipher), 0644); err != nil {
				log.Fatalf("Failed to write output file: %v", err)
			}
			log.Println("Encrypted data written to:", cfg.Output)
		case "decrypt":
			plain, err := utils.DecryptString(cfg.Input, key, cfg.DefaultScheme)
			if err != nil {
				log.Fatalf("Decryption failed: %v", err)
			}
			if err := os.WriteFile(cfg.Output, plain, 0644); err != nil {
				log.Fatalf("Failed to write output file: %v", err)
			}
			log.Println("Decrypted data written to:", cfg.Output)
		default:
			log.Fatalf("Invalid mode: %s", cfg.FileTask.Mode)
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.Flags().StringVar(&configFile, "file", "", "Path to YAML configuration file")
}
