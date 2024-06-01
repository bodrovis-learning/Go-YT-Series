package processor

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"os"
)

func Encrypt(data string) error {
	passphrase, err := GetPassphrase()
	if err != nil {
		return err
	}

	salt, err := makeSalt()
	if err != nil {
		return err
	}

	key, err := DeriveKeyFrom(passphrase, salt)
	if err != nil {
		return err
	}

	crypter, err := MakeCrypterFrom(key)
	if err != nil {
		return err
	}

	nonce, err := makeNonceFor(crypter)
	if err != nil {
		return err
	}

	encryptedData := crypter.Seal(nil, nonce, []byte(data), nil)

	err = saveToFile(encryptedData, nonce, salt)
	if err != nil {
		return err
	}

	return nil
}

func saveToFile(encryptedData, nonce, salt []byte) error {
	encPackage := EncryptedPackage{
		Nonce:         nonce,
		Salt:          salt,
		EncryptedData: encryptedData,
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(encPackage); err != nil {
		return err
	}

	if err := os.WriteFile("encrypted_data.bin", buffer.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func makeSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

func makeNonceFor(crypter cipher.AEAD) ([]byte, error) {
	nonce := make([]byte, crypter.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return nonce, nil
}
