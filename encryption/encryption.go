// Package encryption contains all the ECB and CBC encryption routines
package encryption

import (
	"github.com/mmatww/kmsencrypt/errorhandle"
	"github.com/mmatww/kmsencrypt/padding"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// BlockSize Export this value (which is always 16 lol) to other packages so they don't need
// to import crypto/aes
var BlockSize = aes.BlockSize

// DecryptFile This function uses the decrypted data encryption key and the
// retrieved IV to decrypt the data file
func DecryptFile(data []byte, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		errorhandle.GenError(errors.New("DecryptFile - There was a cipher initialization error"))
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return padding.Unpad(data)
}

// EncryptFile This function uses the provided data encryption key and generates
// an IV to encrypt the data file
func EncryptFile(data []byte, key []byte) ([]byte, []byte) {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		errorhandle.GenError(errors.New("Encryptfile - There was an IV generation error"))
	}
	pmessage := padding.Pad(data)
	ciphertext := make([]byte, len(pmessage))
	c, kerr := aes.NewCipher(key)
	if kerr != nil {
		errorhandle.GenError(errors.New("EncryptFile - There was a cipher initialization error"))
	}
	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(ciphertext, pmessage)
	return ciphertext, iv
}

// GenerateDataKey Does what's on the tin, generates the data encryption key
func GenerateDataKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		errorhandle.GenError(errors.New("GenerateDataKey - There was a key generation error"))
	}
	return key
}
