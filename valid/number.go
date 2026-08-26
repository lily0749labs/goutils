package valid

import (
	"regexp"
)

// Number 验证是否全部为数字。
func (facade) Number(input string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(input)
}

// IsNumber 验证是否全部为数字。
// Deprecated: 使用 Valid.Number。
func IsNumber(input string) bool { return Valid.Number(input) }
