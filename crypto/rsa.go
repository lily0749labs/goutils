package crypto

import (
	stdcrypto "crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

// GenerateRSAKeys 生成 RSA 私钥和公钥。
func (f crypto) GenerateRSAKeys() (string, string, error) {
	return f.GenerateRSAKeysWithReader(cryptorand.Reader, 2048)
}

// GenerateRSAKeysWithReader 使用指定随机源和位数生成 RSA 密钥，便于故障注入和受控测试。
func (crypto) GenerateRSAKeysWithReader(reader io.Reader, bits int) (string, string, error) {
	if reader == nil {
		return "", "", ErrNilRandomReader
	}
	if bits < 2048 {
		return "", "", ErrRSAKeySize
	}
	privateKey, err := rsa.GenerateKey(reader, bits)
	if err != nil {
		return "", "", err
	}

	publicKey := &privateKey.PublicKey

	// 将私钥转换为 PEM 格式。
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 将公钥转换为 PEM 格式。
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

func (crypto) parseRSAPublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, ErrInvalidPublicKey
	}

	if parsed, err := x509.ParsePKIXPublicKey(block.Bytes); err == nil {
		publicKey, ok := parsed.(*rsa.PublicKey)
		if !ok {
			return nil, ErrNotRSAKey
		}
		return publicKey, nil
	}
	if publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes); err == nil {
		return publicKey, nil
	}
	return nil, ErrInvalidPublicKey
}

func (crypto) parseRSAPrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, ErrInvalidPrivateKey
	}

	if privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return privateKey, nil
	}
	if parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		privateKey, ok := parsed.(*rsa.PrivateKey)
		if !ok {
			return nil, ErrNotRSAKey
		}
		return privateKey, nil
	}
	return nil, ErrInvalidPrivateKey
}

// EncryptRSA 使用 RSA 公钥加密数据。
func (f crypto) EncryptRSA(publicKeyStr string, message []byte) ([]byte, error) {
	publicKey, err := f.parseRSAPublicKey(publicKeyStr)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(cryptorand.Reader, publicKey, message)
}

// DecryptRSA 使用 RSA 私钥解密数据。
func (f crypto) DecryptRSA(privateKeyStr string, ciphertext []byte) ([]byte, error) {
	privateKey, err := f.parseRSAPrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(cryptorand.Reader, privateKey, ciphertext)
}

// EncryptRSAOAEP 使用 RSA-OAEP 和 SHA-256 加密数据。
func (f crypto) EncryptRSAOAEP(publicKeyStr string, message []byte) ([]byte, error) {
	return f.EncryptRSAOAEPWithReader(cryptorand.Reader, publicKeyStr, message)
}

// EncryptRSAOAEPWithReader 使用指定随机源执行 RSA-OAEP 加密。
func (f crypto) EncryptRSAOAEPWithReader(reader io.Reader, publicKeyStr string, message []byte) ([]byte, error) {
	if reader == nil {
		return nil, ErrNilRandomReader
	}
	publicKey, err := f.parseRSAPublicKey(publicKeyStr)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptOAEP(sha256.New(), reader, publicKey, message, nil)
}

// DecryptRSAOAEP 使用 RSA-OAEP 和 SHA-256 解密数据。
func (f crypto) DecryptRSAOAEP(privateKeyStr string, ciphertext []byte) ([]byte, error) {
	return f.DecryptRSAOAEPWithReader(cryptorand.Reader, privateKeyStr, ciphertext)
}

// DecryptRSAOAEPWithReader 使用指定随机源执行 RSA-OAEP 解密。
func (f crypto) DecryptRSAOAEPWithReader(reader io.Reader, privateKeyStr string, ciphertext []byte) ([]byte, error) {
	if reader == nil {
		return nil, ErrNilRandomReader
	}
	privateKey, err := f.parseRSAPrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptOAEP(sha256.New(), reader, privateKey, ciphertext, nil)
}

// SignRSAPSS 使用 RSA-PSS 和 SHA-256 为消息签名。
func (f crypto) SignRSAPSS(privateKeyStr string, message []byte) ([]byte, error) {
	return f.SignRSAPSSWithReader(cryptorand.Reader, privateKeyStr, message)
}

// SignRSAPSSWithReader 使用指定随机源执行 RSA-PSS 签名。
func (f crypto) SignRSAPSSWithReader(reader io.Reader, privateKeyStr string, message []byte) ([]byte, error) {
	if reader == nil {
		return nil, ErrNilRandomReader
	}
	privateKey, err := f.parseRSAPrivateKey(privateKeyStr)
	if err != nil {
		return nil, err
	}
	digest := sha256.Sum256(message)
	return rsa.SignPSS(reader, privateKey, stdcrypto.SHA256, digest[:], nil)
}

// VerifyRSAPSS 使用 RSA-PSS 和 SHA-256 校验消息签名。
func (f crypto) VerifyRSAPSS(publicKeyStr string, message, signature []byte) error {
	publicKey, err := f.parseRSAPublicKey(publicKeyStr)
	if err != nil {
		return err
	}
	digest := sha256.Sum256(message)
	if err = rsa.VerifyPSS(publicKey, stdcrypto.SHA256, digest[:], signature, nil); err != nil {
		return fmt.Errorf("%w: %v", ErrRSASignature, err)
	}
	return nil
}

// Deprecated: 使用 Crypto.GenerateRSAKeys。
func GenerateRSAKeys() (string, string, error) { return Crypto.GenerateRSAKeys() }

// Deprecated: 使用 Crypto.EncryptRSA。
func EncryptRSA(publicKeyStr string, message []byte) ([]byte, error) {
	return Crypto.EncryptRSA(publicKeyStr, message)
}

// Deprecated: 使用 Crypto.DecryptRSA。
func DecryptRSA(privateKeyStr string, ciphertext []byte) ([]byte, error) {
	return Crypto.DecryptRSA(privateKeyStr, ciphertext)
}
