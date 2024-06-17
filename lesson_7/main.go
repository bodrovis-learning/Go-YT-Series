package main

import (
	"lesson7/cmd"
	"lesson7/cmd/keys"
	"lesson7/cmd/signatures"
	"lesson7/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(err, "Initial setup failed")
	}
}
