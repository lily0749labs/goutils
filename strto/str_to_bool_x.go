package strto

import "strings"

// StrToBool string转bool
func StrToBool(v string) bool {
	if strings.EqualFold(v, "true") {
		return true
	}
	i := StrToUint64(v)
	return i == 1
}
