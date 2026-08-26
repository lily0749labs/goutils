package valid

import "regexp"

// Mobile 验证是否为手机号码。
func (valid) Mobile(phone string) bool {
	match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, phone)
	return match
}

// Telephone 验证是否为座机号码。
func (valid) Telephone(telephone string) bool {
	match, _ := regexp.MatchString(`^0\d{2,3}-?\d{7,8}$`, telephone)
	return match
}

// Deprecated: 使用 Valid.Mobile。
func IsMobile(phone string) bool { return Valid.Mobile(phone) }

// Deprecated: 使用 Valid.Telephone。
func IsTelephone(telephone string) bool { return Valid.Telephone(telephone) }
