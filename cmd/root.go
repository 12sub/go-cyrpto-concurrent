package cmd

import (
	"fmt"
	"os"

	"example.com/crypto-cli/internal/config"
	"example.com/crypto-cli/utils"
	"github.com/spf13/cobra"
)

var (
	cfgPath   string
	AppConfig *config.Config
	logfile   bool
	LogLevel  string
)

var rootCmd = &cobra.Command{
	Use:   "Go Encrypter",
	Short: "A CLI Tool to encrypt/decrypt strings and files",
	Long:  "Encrypting and Decrypting strings and/or files using AES Encryption",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cfgPath != "" {
			cfg, err := config.LoadConfig(cfgPath)
			if err != nil {
				fmt.Println("Failed to load config:", err)
				os.Exit(1)
			}
			AppConfig = cfg
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "Path to YAML configuration file")
	rootCmd.PersistentFlags().StringVar(&LogLevel, "loglevel", "info", "Log level: debug, info, warn, error")
	rootCmd.PersistentFlags().BoolVar(&logfile, "logfile", false, "Enable Logging to file (crypto-cli.log)")
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(hashCmd)
	cobra.OnInitialize(initLogger)
}

func initLogger() {
	utils.InitLogger(LogLevel, logfile)
}
