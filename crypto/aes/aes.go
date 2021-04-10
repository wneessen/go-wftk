package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// AES256 encrypt a plaintext string and return IV and CipherText
// as Base64
func EncryptBase64(plainText string, encKey string) (string, error) {
	return EncryptAuthBase64(plainText, encKey, "")
}

// AES256 decrypt a ciphertext string as base64 and return the plaintext
func DecryptBase64(plainText string, encKey string) (string, error) {
	return DecryptAuthBase64(plainText, encKey, "")
}

// AES256 encrypt a plaintext string with authentication data
// and return IV and CipherText as Base64
func EncryptAuthBase64(plainText string, encKey string, authData string) (string, error) {
	data := []byte(plainText)
	key := []byte(encKey)
	auth := []byte(authData)

	cipherText, cryptoIv, err := encryptBytesWithAuthData(data, key, auth)
	if err != nil {
		return "", err
	}

	iv64 := base64.StdEncoding.EncodeToString(cryptoIv)
	ct64 := base64.StdEncoding.EncodeToString(cipherText)
	returnBytes := iv64 + ct64

	return returnBytes, nil
}

// AES256 decrypt a ciphertext string with authentication data as base64
// and return the plaintext
func DecryptAuthBase64(cipherText string, decKey, authData string) (string, error) {
	auth := []byte(authData)
	key := []byte(decKey)
	cb64, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	civ := cb64[:12]
	ct := cb64[12:]

	plainText, err := decryptBytesWithAuthData(ct, key, civ, auth)
	if err != nil {
		return "", err
	}
	fmt.Println(string(plainText))

	return string(plainText), nil
}

// Internal function to handle encryption
func encryptBytesWithAuthData(d []byte, k []byte, a []byte) ([]byte, []byte, error) {
	// Check key length
	if len(k) != 32 {
		return []byte{}, []byte{}, fmt.Errorf("invalid key size, required key size is 256 bits")
	}

	// Create a new cipher object
	cphr, err := aes.NewCipher(k)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	// Read random data for nonce
	cryptoIv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, cryptoIv); err != nil {
		return []byte{}, []byte{}, err
	}

	// Create a GCM block cipher
	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		return []byte{}, []byte{}, err
	}

	// Encrypt the data
	cipherText := gcm.Seal(nil, cryptoIv, d, a)

	return cipherText, cryptoIv, nil
}

// Internal function to handle decryption
func decryptBytesWithAuthData(d []byte, k []byte, i []byte, a []byte) ([]byte, error) {
	// Create a new cipher object
	cphr, err := aes.NewCipher(k)
	if err != nil {
		return []byte{}, err
	}

	// Create a GCM block cipher
	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		return []byte{}, err
	}

	// Encrypt the data
	plainText, err := gcm.Open(nil, i, d, a)
	if err != nil {
		return []byte{}, err
	}

	return plainText, nil
}
