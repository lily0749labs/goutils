package stringutil

import "unicode"

// EscapeSpecial 在 Unicode 标点、符号和汉字前添加反斜杠。
// 该行为用于兼容旧查询字符串转义规则；它不是 JSON、SQL 或 URL 转义。
func (stringutil) EscapeSpecial(value string) string {
	result := make([]rune, 0, len([]rune(value))*2)
	for _, current := range value {
		if unicode.IsPunct(current) || unicode.IsSymbol(current) || unicode.Is(unicode.Han, current) {
			result = append(result, '\\')
		}
		result = append(result, current)
	}
	return string(result)
}

// EscapeSpecial 是 String.EscapeSpecial 的包级便捷入口。
func EscapeSpecial(value string) string {
	return String.EscapeSpecial(value)
}
