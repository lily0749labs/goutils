package crypto

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 返回 str 的十六进制 SHA-512 摘要。
func (crypto) SHA512(str string) string {
	sum := sha512.Sum512([]byte(str))
	return hex.EncodeToString(sum[:])
}

// SHA512 是 Crypto.SHA512 的包级便捷入口。
func SHA512(str string) string {
	return Crypto.SHA512(str)
}
