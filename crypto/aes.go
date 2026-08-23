package crypto

// AES 加解密工具。

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// ErrAESKeyTooShort 表示 AES 密钥不足 16 字节。
var ErrAESKeyTooShort = errors.New("aes key must be at least 16 bytes")

// AESEncryptor 提供 AES 字符串加密和解密能力。
type AESEncryptor struct{}

// AesEncrypt 是为兼容旧代码而保留的类型别名。
// Deprecated: 请使用 AESEncryptor。
type AesEncrypt = AESEncryptor

func (e *AESEncryptor) getKey(strKey string) ([]byte, error) {
	arrKey := []byte(strKey)
	keyLen := len(arrKey)
	if keyLen < 16 {
		return nil, ErrAESKeyTooShort
	}
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32], nil
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24], nil
	}
	//取前16个字节
	return arrKey[:16], nil
}

// Encrypt 使用 strKey 加密 strMesg。
func (e *AESEncryptor) Encrypt(strKey, strMesg string) ([]byte, error) {
	key, err := e.getKey(strKey)
	if err != nil {
		return nil, err
	}
	iv := key[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

// Decrypt 使用 strKey 解密 src。
func (e *AESEncryptor) Decrypt(strKey string, src []byte) ([]byte, error) {
	key, err := e.getKey(strKey)
	if err != nil {
		return nil, err
	}
	iv := key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	aesBlockDecrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}
