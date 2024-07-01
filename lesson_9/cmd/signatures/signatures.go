package signatures

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

var signaturesCmd = &cobra.Command{
	Use:   "signatures",
	Short: "Create and verify signatures.",
	Long:  `Use subcommands to create signature (.sig) with private key and verify signature with public key.`,
}

// Init initializes signatures commands
func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(signaturesCmd)

	signaturesCmd.PersistentFlags().String("file-path", "", "Path to the file that should be signed")
}

func hashFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %v", filePath, err)
	}
	defer file.Close()

	hasher := sha3.New256()

	if _, err := io.Copy(hasher, file); err != nil {
		return nil, fmt.Errorf("error while hashing file %s: %v", filePath, err)
	}

	return hasher.Sum(nil), nil
}
