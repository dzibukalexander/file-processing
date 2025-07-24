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
	var encryptor Encryptor
	switch encType {
	case AES:
		encryptor = &aes.AESEncryptor{}
	case RSA:
		encryptor = &rsa.RSAEncryptor{}
	default:
		return nil
	}
	return NewLoggingEncryptor(encryptor)
}

func NewDecryptor(encType EncryptionType) Decryptor {
	var decryptor Decryptor
	switch encType {
	case AES:
		decryptor = &aes.AESDecryptor{}
	case RSA:
		decryptor = &rsa.RSADecryptor{}
	default:
		return nil
	}
	return NewLoggingDecryptor(decryptor)
}
