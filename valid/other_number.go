package valid

import "regexp"

// PostalCode 验证是否为邮编号码。
func (valid) PostalCode(str string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{5}$`)
	return reg.MatchString(str)
}

// IsPostalCode 验证是否为邮编号码。
// Deprecated: 使用 Valid.PostalCode。
func IsPostalCode(str string) bool { return Valid.PostalCode(str) }
