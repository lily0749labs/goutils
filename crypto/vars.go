package crypto

import "errors"

// Crypto 提供加密与哈希的结构化入口。
var Crypto = crypto{}

type crypto struct{}

var (
	// ErrInvalidPublicKey 表示无法解码或解析 RSA 公钥。
	ErrInvalidPublicKey = errors.New("invalid RSA public key")
	// ErrInvalidPrivateKey 表示无法解码或解析 RSA 私钥。
	ErrInvalidPrivateKey = errors.New("invalid RSA private key")
	// ErrNotRSAKey 表示 PEM 内容是其他类型的密钥。
	ErrNotRSAKey = errors.New("key is not an RSA key")
	// ErrNilRandomReader 表示调用方传入了空随机源。
	ErrNilRandomReader = errors.New("random reader must not be nil")
	// ErrRSAKeySize 表示 RSA 密钥位数低于安全下限。
	ErrRSAKeySize = errors.New("RSA key size must be at least 2048 bits")
	// ErrRSASignature 表示 RSA-PSS 签名校验失败。
	ErrRSASignature = errors.New("RSA-PSS signature verification failed")
)
