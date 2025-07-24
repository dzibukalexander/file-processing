// internal/encryption/rsa/rsa.go
package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type RSAEncryptor struct{}

func (e *RSAEncryptor) Encrypt(data []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

type RSADecryptor struct{}

func (d *RSADecryptor) Decrypt(data []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

func (e *RSAEncryptor) GenerateKey(path string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	// Save private key
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}
	privFile, err := os.Create(path + ".priv")
	if err != nil {
		return fmt.Errorf("failed to create private key file: %w", err)
	}
	defer privFile.Close()
	if err := pem.Encode(privFile, privBlock); err != nil {
		return fmt.Errorf("failed to write private key: %w", err)
	}

	// Save public key
	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %w", err)
	}
	pubBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}
	pubFile, err := os.Create(path + ".pub")
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer pubFile.Close()
	return pem.Encode(pubFile, pubBlock)
}
