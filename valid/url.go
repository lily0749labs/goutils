package valid

import (
	"regexp"
)

// URL 验证是否为 URL 地址。
func (facade) URL(url string) bool {
	match, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$`, url)
	return match
}

// IsURL 验证是否为 URL 地址。
// Deprecated: 使用 Valid.URL。
func IsURL(url string) bool { return Valid.URL(url) }
