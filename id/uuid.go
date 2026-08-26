package id

import uuid "github.com/satori/go.uuid"

// Token 生成 UUID 格式的令牌。
func (id) Token() string {
	return uuid.NewV4().String()
}

// GetToken 生成 UUID 格式的令牌。
// Deprecated: 使用 ID.Token。
func GetToken() string { return ID.Token() }
