package strto

import "strings"

// Bool 将字符串转换为 bool。
func (f facade) Bool(v string) bool {
	if strings.EqualFold(v, "true") {
		return true
	}
	i := f.Uint64(v)
	return i == 1
}

// StrToBool 将字符串转换为 bool。
// Deprecated: 使用 StrTo.Bool。
func StrToBool(v string) bool { return StrTo.Bool(v) }
