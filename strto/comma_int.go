package strto

import (
	"encoding/json"
	"strconv"
	"strings"
)

// CommaInts 将逗号分隔的字符串转换为 int 切片。
// 空字段和无法转换的字段会被忽略；空字符串返回 nil。
func (strto) CommaInts(value string) []int {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	var result []int
	for _, part := range parts {
		if part == "" {
			continue
		}
		item, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result
}

// ParseCommaInts 严格解析逗号分隔的 int 列表。
// 首尾各允许一个逗号；任何无效字段都会返回错误。
func (strto) ParseCommaInts(value string) ([]int, error) {
	if value == "" {
		return nil, nil
	}
	value = strings.TrimPrefix(value, ",")
	value = strings.TrimSuffix(value, ",")
	var result []int
	if err := json.Unmarshal([]byte("["+value+"]"), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CommaInts 是 StrTo.CommaInts 的包级便捷入口。
func CommaInts(value string) []int {
	return StrTo.CommaInts(value)
}

// ParseCommaInts 是 StrTo.ParseCommaInts 的包级便捷入口。
func ParseCommaInts(value string) ([]int, error) {
	return StrTo.ParseCommaInts(value)
}
