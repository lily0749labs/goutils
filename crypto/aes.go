package crypto

// AES 加解密工具。

import (
	"crypto/aes"
	"crypto/cipher"
	cryptorand "crypto/rand"
	"errors"
)

var (
	// ErrAESKeyTooShort 表示兼容 AES 方法收到不足 16 字节的密钥。
	ErrAESKeyTooShort = errors.New("aes key must be at least 16 bytes")
	// ErrAESInvalidKeySize 表示 AES-GCM 密钥长度不是 16、24 或 32 字节。
	ErrAESInvalidKeySize = errors.New("aes key must be exactly 16, 24, or 32 bytes")
	// ErrAESGCMCiphertextTooShort 表示 AES-GCM 数据中缺少完整的 Nonce 或认证标签。
	ErrAESGCMCiphertextTooShort = errors.New("aes-gcm ciphertext is too short")
)

// AESEncryptor 提供 AES 字符串加密和解密能力。
type AESEncryptor struct{}

// AesEncrypt 是为兼容旧代码而保留的类型别名。
// Deprecated: 请使用 AESEncryptor。
type AesEncrypt = AESEncryptor

func (crypto) getAESKey(strKey string) ([]byte, error) {
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

func (crypto) newGCM(strKey string) (cipher.AEAD, error) {
	key := []byte(strKey)
	switch len(key) {
	case 16, 24, 32:
	default:
		return nil, ErrAESInvalidKeySize
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// EncryptGCM 使用 AES-GCM 加密明文。
// 返回值依次包含随机 Nonce、密文和认证标签，可直接交给 DecryptGCM 解密。
func (f crypto) EncryptGCM(strKey, plaintext string) ([]byte, error) {
	gcm, err := f.newGCM(strKey)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = cryptorand.Read(nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

// DecryptGCM 解密 EncryptGCM 生成的数据，并校验数据是否被篡改。
func (f crypto) DecryptGCM(strKey string, ciphertext []byte) ([]byte, error) {
	gcm, err := f.newGCM(strKey)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize+gcm.Overhead() {
		return nil, ErrAESGCMCiphertextTooShort
	}
	nonce, encrypted := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, encrypted, nil)
}

// Encrypt 使用 AES 加密 strMesg。
func (f crypto) Encrypt(strKey, strMesg string) ([]byte, error) {
	key, err := f.getAESKey(strKey)
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

// Decrypt 使用 AES 解密 src。
func (f crypto) Decrypt(strKey string, src []byte) ([]byte, error) {
	key, err := f.getAESKey(strKey)
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

// Encrypt 使用 strKey 加密 strMesg。
// Deprecated: 使用 Crypto.Encrypt。
func (*AESEncryptor) Encrypt(strKey, strMesg string) ([]byte, error) {
	return Crypto.Encrypt(strKey, strMesg)
}

// Decrypt 使用 strKey 解密 src。
// Deprecated: 使用 Crypto.Decrypt。
func (*AESEncryptor) Decrypt(strKey string, src []byte) ([]byte, error) {
	return Crypto.Decrypt(strKey, src)
}
