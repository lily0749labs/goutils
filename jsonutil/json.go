package jsonutil

import (
	"bytes"
	"encoding/json"
)

// JSON 提供 JSON 辅助函数的结构化入口。
var JSON = jsonutil{}

type jsonutil struct{}

// MarshalString 将 value 编码为 JSON 字符串。
func (jsonutil) MarshalString(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// MarshalStringNoEscapeHTML 将 value 编码为不转义 HTML 字符的 JSON 字符串。
// 为兼容 json.Encoder 的原始行为，返回值末尾包含换行符。
func (jsonutil) MarshalStringNoEscapeHTML(value any) (string, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// MarshalString 是 JSON.MarshalString 的包级便捷入口。
func MarshalString(value any) (string, error) {
	return JSON.MarshalString(value)
}

// MarshalStringNoEscapeHTML 是 JSON.MarshalStringNoEscapeHTML 的包级便捷入口。
func MarshalStringNoEscapeHTML(value any) (string, error) {
	return JSON.MarshalStringNoEscapeHTML(value)
}
