package strto

import "errors"

// StrTo 提供字符串转换的结构化入口。
var StrTo = strto{}

// ErrInvalidBool 表示字符串不是支持的布尔值。
var ErrInvalidBool = errors.New("value is not a supported boolean")

type strto struct{}
