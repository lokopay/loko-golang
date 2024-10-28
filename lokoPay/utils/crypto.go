package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/wumansgy/goEncrypt/aes"
)

func GenerateHMAC(message, secretKey string) string {
	// Convert the secret key to a byte slice
	key := []byte(secretKey)

	// Create a new HMAC by defining the hash type and the key
	h := hmac.New(sha256.New, key)

	// Write the message to the HMAC object
	h.Write([]byte(message))

	// Get the final HMAC result
	hmac := h.Sum(nil)

	// Return the HMAC as a hexadecimal string
	return base64.StdEncoding.EncodeToString(hmac)
}

func GenerateHashKey(key string, length int) ([]byte, error) {
	hasher := sha256.New()
	if _, err := hasher.Write([]byte(key)); err != nil {
		return nil, err
	}
	derivedKey := hasher.Sum(nil)
	return derivedKey[:length], nil // Take first 16 bytes for AES key
}

// AesEncrypt AES加密
func AesEncrypt(message, key string) string {
	hash, err := GenerateHashKey(key, 16)
	if err != nil {
		return ""
	}
	base64Text, err := aes.AesEcbEncrypt([]byte(message), hash)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(base64Text)
}

// AesDecrypt AES解密
func AesDecrypt(message, key string) string {
	if message == "" {
		return ""
	}
	hash, err := GenerateHashKey(key, 16)
	if err != nil {
		return ""
	}
	plaintext, err := aes.AesEcbDecryptByBase64(message, hash)
	if err != nil {
		return ""
	}
	return string(plaintext)
}
