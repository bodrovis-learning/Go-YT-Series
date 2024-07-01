package signatures

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"brave_signer/internal/config"
	"brave_signer/internal/logger"
	"brave_signer/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	signaturesCmd.AddCommand(signaturesVerifyFileCmd)

	signaturesVerifyFileCmd.Flags().String("pub-key-path", "pub_key.pem", "Path to the public key")
}

var signaturesVerifyFileCmd = &cobra.Command{
	Use:   "verifyfile",
	Short: "Verify the signature of a file.",
	Long:  `Verify the digital signature of a specified file using an RSA public key. The command expects a signature file named "<original_filename>.sig" located in the same directory as the file being verified. The public key should be in PEM format.`,
	Run: func(cmd *cobra.Command, args []string) {
		localViper := cmd.Context().Value(config.ViperKey).(*viper.Viper)

		fullPubKeyPath, err := utils.ProcessFilePath(localViper.GetString("pub-key-path"))
		logger.HaltOnErr(err, "failed to process pub key path")

		fullFilePath, err := utils.ProcessFilePath(localViper.GetString("file-path"))
		logger.HaltOnErr(err, "failed to process file path")

		signatureRaw, err := readSignature(fullFilePath)
		logger.HaltOnErr(err, "cannot read signature")

		digest, err := hashFile(fullFilePath)
		logger.HaltOnErr(err, "cannot hash file")

		publicKey, err := loadPublicKey(fullPubKeyPath)
		logger.HaltOnErr(err, "cannot load pub key from file")

		signerInfo, err := verifyFileSignature(publicKey, digest, signatureRaw)
		logger.HaltOnErr(err, "cannot verify signature")

		logger.Info(fmt.Sprintf("Verification successful for file: %s\n", filepath.Base(fullFilePath)))
		logger.Info(fmt.Sprintf("Verified using public key: %s\n", filepath.Base(fullPubKeyPath)))
		logger.Info(fmt.Sprintf("Signer info:\n%s\n", signerInfo))
	},
}

func readSignature(initialFilePath string) ([]byte, error) {
	dir := filepath.Dir(initialFilePath)
	baseName := filepath.Base(initialFilePath)

	sigFilePath := filepath.Join(dir, baseName+".sig")

	return os.ReadFile(sigFilePath)
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %v", err)
	}

	block, rest := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("public key file does not contain a valid PEM block")
	}
	if len(rest) > 0 {
		return nil, fmt.Errorf("additional data found after the first PEM block, which could indicate multiple PEM blocks or corrupted data")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, fmt.Errorf("public key file does not contain an RSA public key")
	}
}

func verifyFileSignature(publicKey *rsa.PublicKey, digest []byte, signatureRaw []byte) ([]byte, error) {
	buf := bytes.NewReader(signatureRaw)

	var nameLength uint32
	if err := binary.Read(buf, binary.BigEndian, &nameLength); err != nil {
		return nil, fmt.Errorf("failed to read signer info length: %v", err)
	}

	signerInfo := make([]byte, nameLength)
	if _, err := buf.Read(signerInfo); err != nil {
		return nil, fmt.Errorf("failed to read signer info: %v", err)
	}

	signature := make([]byte, buf.Len())
	if _, err := buf.Read(signature); err != nil {
		return nil, fmt.Errorf("failed to read signature: %v", err)
	}

	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto}
	if err := rsa.VerifyPSS(publicKey, crypto.SHA3_256, digest, signature, opts); err != nil {
		return nil, fmt.Errorf("signature verification failed: %v", err)
	}

	return signerInfo, nil
}
