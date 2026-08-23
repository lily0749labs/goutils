package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 返回 str 的十六进制 MD5 摘要。
func MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

// Md5 返回 str 的十六进制 MD5 摘要。
// Deprecated: 请使用 MD5。
func Md5(str string) string {
	return MD5(str)
}
