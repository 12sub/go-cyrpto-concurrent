package main

import (
	"example.com/crypto-cli/cmd"
	_ "example.com/crypto-cli/plugins"
	"example.com/crypto-cli/utils"
)

func main() {
	defer utils.Cleanup()
	cmd.Execute()
}
