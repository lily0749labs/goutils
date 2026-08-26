package valid

import "regexp"

// Decimal 验证给定的字符串小数点后是否最多两位。
func (valid) Decimal(input string) bool {
	// 定义小数的正则表达式，小数点后最多两位
	pattern := `^[0-9]+(\.[0-9]{1,2})?$`
	match, err := regexp.MatchString(pattern, input)
	return match && err == nil
}

// IsDecimal 验证给定的字符串小数点后是否最多两位。
// Deprecated: 使用 Valid.Decimal。
func IsDecimal(input string) bool { return Valid.Decimal(input) }
