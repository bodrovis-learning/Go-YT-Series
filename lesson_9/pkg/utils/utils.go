package utils

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"path/filepath"

	"brave_signer/internal/logger"
)

// ProcessFilePath converts a given path to an absolute path and verifies it points to a regular file.
func ProcessFilePath(path string) (string, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("converting to absolute path: %v", err)
	}

	fileInfo, err := os.Stat(absolutePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("path '%s' does not exist", path)
		}
		return "", fmt.Errorf("fetching path info: %v", err)
	}

	if !fileInfo.Mode().IsRegular() {
		return "", fmt.Errorf("path '%s' does not point to a file", path)
	}

	return absolutePath, nil
}

// GetPassphrase prompts the user for a passphrase and securely reads it.
func GetPassphrase() ([]byte, error) {
	fmt.Println("Enter passphrase:")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to set terminal to raw mode: %w", err)
	}
	defer safeRestore(int(os.Stdin.Fd()), oldState)

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to read passphrase: %w", err)
	}

	return passphrase, nil
}

// safeRestore attempts to restore the terminal to its original state and logs an error if it fails.
func safeRestore(fd int, state *term.State) {
	if err := term.Restore(fd, state); err != nil {
		logger.HaltOnErr(fmt.Errorf("failed to restore terminal state: %v", err), "Terminal restoration failed")
	}
}
