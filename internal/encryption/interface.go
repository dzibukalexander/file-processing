// internal/encryption/interface.go
package encryption

import (
	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	. "github.com/dzibukalexander/file-processing/internal/encryption/constants"
	"github.com/dzibukalexander/file-processing/internal/encryption/rsa"
)

type Encryptor interface {
	Encrypt(data []byte, key []byte) ([]byte, error)
}

type Decryptor interface {
	Decrypt(data []byte, key []byte) ([]byte, error)
}

func NewEncryptor(encType EncryptionType) Encryptor {
	switch encType {

	case AES:
		return &aes.AESEncryptor{}
	case RSA:
		return &rsa.RSAEncryptor{}
	default:
		return nil
	}
}

func NewDecryptor(encType EncryptionType) Decryptor {
	switch encType {
	case AES:
		return &aes.AESDecryptor{}
	case RSA:
		return &rsa.RSADecryptor{}
	default:
		return nil
	}
}
