package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 返回 str 的十六进制 MD5 摘要。
func (crypto) MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

// Md5 返回 str 的十六进制 MD5 摘要。
// Deprecated: 使用 Crypto.MD5。
func (f crypto) Md5(str string) string { return f.MD5(str) }

// Deprecated: 使用 Crypto.MD5。
func MD5(str string) string { return Crypto.MD5(str) }

// Deprecated: 使用 Crypto.MD5。
func Md5(str string) string { return Crypto.MD5(str) }
