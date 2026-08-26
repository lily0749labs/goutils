package strto

import (
	"strconv"
	"strings"
)

// BoolE 将字符串严格转换为 bool。
// 支持 true、false 及数值形式的 1、0；其他值返回错误。
func (strto) BoolE(v string) (bool, error) {
	if strings.EqualFold(v, "true") {
		return true, nil
	}
	if strings.EqualFold(v, "false") {
		return false, nil
	}

	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return false, ErrInvalidBool
	}
	switch i {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, ErrInvalidBool
	}
}

// Bool 将字符串转换为 bool。
// 转换失败时返回 false；需要错误信息时请使用 BoolE。
func (s strto) Bool(v string) bool {
	value, _ := s.BoolE(v)
	return value
}

// StrToBool 将字符串转换为 bool。
// Deprecated: 使用 StrTo.Bool。
func StrToBool(v string) bool { return StrTo.Bool(v) }
