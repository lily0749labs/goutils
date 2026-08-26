package valid

import (
	"regexp"
	"unicode"
)

// AllChinese 验证给定的字符串全部为中文。
func (facade) AllChinese(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

// ContainsChinese 验证给定的字符串包含中文。
func (facade) ContainsChinese(input string) bool {
	for _, r := range input {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// ChineseName 验证是否为中文名。
func (facade) ChineseName(name string) bool {
	pattern := "^[\u4E00-\u9FA5]{2,6}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// EnglishName 验证是否为英文名。
func (facade) EnglishName(name string) bool {
	match, _ := regexp.MatchString(`^([a-zA-Z]+\s)*[a-zA-Z]+$`, name)
	return match
}

// Deprecated: 使用 Valid.AllChinese。
func IsAllChinese(input string) bool { return Valid.AllChinese(input) }

// Deprecated: 使用 Valid.ContainsChinese。
func IsContainChinese(input string) bool { return Valid.ContainsChinese(input) }

// Deprecated: 使用 Valid.ChineseName。
func IsChineseName(name string) bool { return Valid.ChineseName(name) }

// Deprecated: 使用 Valid.EnglishName。
func IsEnglishName(name string) bool { return Valid.EnglishName(name) }
