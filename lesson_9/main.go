package main

import (
	"errors"

	"brave_signer/cmd"
	"brave_signer/internal/logger"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(errors.New("cannot proceed, exiting now"), "Initial setup failed")
	}
}
