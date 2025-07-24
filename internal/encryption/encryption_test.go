package encryption

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	"github.com/dzibukalexander/file-processing/internal/encryption/rsa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "encryption-test")
	require.NoError(t, err)
	cleanup := func() { os.RemoveAll(tempDir) }
	return tempDir, cleanup
}

func TestAESEncryptDecrypt(t *testing.T) {
	encryptor := &aes.AESEncryptor{}
	decryptor := &aes.AESDecryptor{}
	originalData := []byte("hello aes")
	key := make([]byte, 32)

	encrypted, err := encryptor.Encrypt(originalData, key)
	require.NoError(t, err)

	decrypted, err := decryptor.Decrypt(encrypted, key)
	require.NoError(t, err)

	assert.Equal(t, originalData, decrypted)
}

func TestRSAEncryptDecrypt(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	keyPath := filepath.Join(tempDir, "rsa_key")
	encryptor := &rsa.RSAEncryptor{}
	decryptor := &rsa.RSADecryptor{}

	require.NoError(t, encryptor.GenerateKey(keyPath))

	pubKey, err := ioutil.ReadFile(keyPath + ".pub")
	require.NoError(t, err)
	privKey, err := ioutil.ReadFile(keyPath + ".priv")
	require.NoError(t, err)

	originalData := []byte("hello rsa")
	encrypted, err := encryptor.Encrypt(originalData, pubKey)
	require.NoError(t, err)

	decrypted, err := decryptor.Decrypt(encrypted, privKey)
	require.NoError(t, err)

	assert.Equal(t, originalData, decrypted)
}

func TestKeyGeneration(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	// Test AES key generation
	aesKeyPath := filepath.Join(tempDir, "aes.key")
	aesGen := &aes.AESEncryptor{}
	err := aesGen.GenerateKey(aesKeyPath)
	require.NoError(t, err)
	_, err = os.Stat(aesKeyPath)
	assert.NoError(t, err)

	// Test RSA key generation
	rsaKeyPath := filepath.Join(tempDir, "rsa_key")
	rsaGen := &rsa.RSAEncryptor{}
	err = rsaGen.GenerateKey(rsaKeyPath)
	require.NoError(t, err)
	_, err = os.Stat(rsaKeyPath + ".pub")
	assert.NoError(t, err)
	_, err = os.Stat(rsaKeyPath + ".priv")
	assert.NoError(t, err)
}
