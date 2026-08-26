package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// HashSHA256 返回 str 的 SHA-256 摘要。
func (facade) HashSHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// BcryptHash 使用 bcrypt 对密码进行加密。
func (f facade) BcryptHash(password string) string {
	hash, _ := f.BcryptHashWithError(password)
	return hash
}

// BcryptHashWithError 使用 bcrypt 对密码进行加密，并返回输入或加密错误。
func (facade) BcryptHashWithError(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// BcryptCheck 对比明文密码和数据库的哈希值。
func (facade) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5V 返回字节切片的十六进制 MD5 摘要，并允许在摘要前附加字节。
func (facade) MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// Deprecated: 使用 Crypto.HashSHA256。
func HashSHA256(str string) string { return Crypto.HashSHA256(str) }

// Deprecated: 使用 Crypto.BcryptHash。
func BcryptHash(password string) string { return Crypto.BcryptHash(password) }

// Deprecated: 使用 Crypto.BcryptHashWithError。
func BcryptHashWithError(password string) (string, error) {
	return Crypto.BcryptHashWithError(password)
}

// Deprecated: 使用 Crypto.BcryptCheck。
func BcryptCheck(password, hash string) bool { return Crypto.BcryptCheck(password, hash) }

// Deprecated: 使用 Crypto.MD5V。
func MD5V(str []byte, b ...byte) string { return Crypto.MD5V(str, b...) }
