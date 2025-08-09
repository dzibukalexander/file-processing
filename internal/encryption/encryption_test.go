package encryption

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	"github.com/dzibukalexander/file-processing/internal/encryption/rsa"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "encryption-test")
	if err != nil {
		t.Fatalf("temp dir: %v", err)
	}
	cleanup := func() { os.RemoveAll(tempDir) }
	return tempDir, cleanup
}

func TestAESEncryptDecrypt(t *testing.T) {
	runner.Run(t, "AES encrypt/decrypt", func(t provider.T) {
		t.WithNewStep("roundtrip", func(s provider.StepCtx) {
			encryptor := &aes.AESEncryptor{}
			decryptor := &aes.AESDecryptor{}
			originalData := []byte("hello aes")
			key := make([]byte, 32)

			encrypted, err := encryptor.Encrypt(originalData, key)
			s.Require().NoError(err)

			decrypted, err := decryptor.Decrypt(encrypted, key)
			s.Require().NoError(err)
			s.Assert().Equal(originalData, decrypted)
		})
	})
}

func TestRSAEncryptDecrypt(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	runner.Run(t, "RSA encrypt/decrypt", func(t provider.T) {
		t.WithNewStep("roundtrip", func(s provider.StepCtx) {
			keyPath := filepath.Join(tempDir, "rsa_key")
			encryptor := &rsa.RSAEncryptor{}
			decryptor := &rsa.RSADecryptor{}

			s.Require().NoError(encryptor.GenerateKey(keyPath))

			pubKey, err := ioutil.ReadFile(keyPath + ".pub")
			s.Require().NoError(err)
			privKey, err := ioutil.ReadFile(keyPath + ".priv")
			s.Require().NoError(err)

			originalData := []byte("hello rsa")
			encrypted, err := encryptor.Encrypt(originalData, pubKey)
			s.Require().NoError(err)

			decrypted, err := decryptor.Decrypt(encrypted, privKey)
			s.Require().NoError(err)
			s.Assert().Equal(originalData, decrypted)
		})
	})
}

func TestKeyGeneration(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	runner.Run(t, "Key generation", func(t provider.T) {
		t.WithNewStep("AES and RSA", func(s provider.StepCtx) {
			// AES
			aesKeyPath := filepath.Join(tempDir, "aes.key")
			aesGen := &aes.AESEncryptor{}
			s.Require().NoError(aesGen.GenerateKey(aesKeyPath))
			_, err := os.Stat(aesKeyPath)
			s.Assert().NoError(err)

			// RSA
			rsaKeyPath := filepath.Join(tempDir, "rsa_key")
			rsaGen := &rsa.RSAEncryptor{}
			s.Require().NoError(rsaGen.GenerateKey(rsaKeyPath))
			_, err = os.Stat(rsaKeyPath + ".pub")
			s.Assert().NoError(err)
			_, err = os.Stat(rsaKeyPath + ".priv")
			s.Assert().NoError(err)
		})
	})
}
