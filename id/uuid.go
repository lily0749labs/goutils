package id

import uuid "github.com/satori/go.uuid"

// GetToken 生成 UUID 格式的令牌。
func GetToken() string {
	return uuid.NewV4().String()
}
