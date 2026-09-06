package strto

import "errors"

// StrTo 提供字符串转换的结构化入口。
var StrTo = strto{}

var (
	// ErrInvalidBool 表示字符串不是支持的布尔值。
	ErrInvalidBool = errors.New("value is not a supported boolean")
	// ErrNonFiniteFloat 表示浮点数是 NaN 或无穷大。
	ErrNonFiniteFloat = errors.New("float value must be finite")
	// ErrInvalidDuration 表示字符串不是有效的 Go 时长。
	ErrInvalidDuration = errors.New("invalid duration")
)

type strto struct{}
