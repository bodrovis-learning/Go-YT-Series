package main

import (
	"errors"

	"brave_signer/cmd"
	"brave_signer/cmd/keys"
	"brave_signer/cmd/signatures"
	"brave_signer/logger"
)

func main() {
	rootCmd := cmd.RootCmd()
	keys.Init(rootCmd)
	signatures.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(errors.New("cannot proceed, exiting now"), "Initial setup failed")
	}
}
