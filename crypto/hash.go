package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// HashSHA256 hash加密
func HashSHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	hash, _ := BcryptHashWithError(password)
	return hash
}

// BcryptHashWithError 使用 bcrypt 对密码进行加密，并返回输入或加密错误。
func BcryptHashWithError(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5V 返回字节切片的十六进制 MD5 摘要，并允许在摘要前附加字节。
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
