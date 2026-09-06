package strto

import (
	"strconv"
	"strings"
)

// ParseBool 将字符串解析为 bool。
// 它接受不区分大小写的 true、false，以及值为 0 或 1 的十进制整数形式。
func (strto) ParseBool(value string) (bool, error) {
	switch {
	case strings.EqualFold(value, "true"):
		return true, nil
	case strings.EqualFold(value, "false"):
		return false, nil
	}

	numeric, err := strconv.ParseUint(value, 10, 64)
	if err != nil || numeric > 1 {
		return false, ErrInvalidBool
	}
	return numeric == 1, nil
}

// BoolOrFalse 将字符串解析为 bool；解析失败时返回 false。
func (s strto) BoolOrFalse(value string) bool {
	parsed, _ := s.ParseBool(value)
	return parsed
}
