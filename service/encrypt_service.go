package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/ahmadsalafudin/simple-encrypt-generator/config"
)

// RSA key pair
var rsaPrivateKey *rsa.PrivateKey
var rsaPublicKey *rsa.PublicKey

// Initialize RSA keys
func InitRSA() {
	var err error
	rsaPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generating RSA key pair: %v", err)
	}
	rsaPublicKey = &rsaPrivateKey.PublicKey
}

// Encrypt text using the specified method
func Encrypt(text, method string) string {
	var result string
	switch method {
	case config.Base64:
		result = base64.StdEncoding.EncodeToString([]byte(text))
	case config.Hash:
		result = hashText(text)
	case config.RSA:
		encryptedText, err := rsaEncrypt(text)
		if err != nil {
			log.Fatalf("Error encrypting with RSA: %v", err)
		}
		result = encryptedText
	default:
		return ""
	}
	return "Basic " + result
}

// Decrypt text using the specified method
func Decrypt(cipherText, method, originalText string) string {
	// Remove "Basic " prefix if present
	cipherText = strings.TrimPrefix(cipherText, "Basic ")

	switch method {
	case config.Base64:
		decoded, err := base64.StdEncoding.DecodeString(cipherText)
		if err != nil {
			log.Fatalf("Error decoding Base64: %v", err)
		}
		return string(decoded)
	case config.RSA:
		decryptedText, err := rsaDecrypt(cipherText)
		if err != nil {
			log.Fatalf("Error decrypting with RSA: %v", err)
		}
		return decryptedText
	case config.Hash:
		match := CompareHash(originalText, cipherText)
		if match {
			return "Match"
		}
		return "No Match"
	default:
		return ""
	}
}

// CompareHash checks if the hash of the plainText matches the given hash.
func CompareHash(plainText, hash string) bool {
	hashedText := hashText(plainText)
	return hashedText == hash
}

// Hashing function (SHA-256)
func hashText(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// RSA encryption
func rsaEncrypt(text string) (string, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, []byte(text), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// RSA decryption
func rsaDecrypt(cipherText string) (string, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivateKey, encryptedBytes, nil)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}
