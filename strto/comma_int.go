package strto

import "strings"

// ParseCommaIntsLoose 从逗号分隔的字段中提取可解析的 int。
// 它会去除字段两侧空白并忽略空字段、非法整数和溢出值；没有有效整数时返回 nil。
func (s strto) ParseCommaIntsLoose(value string) []int {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		item, err := s.ParseInt(part)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// ParseCommaInts 解析逗号分隔的 int 列表。
// 它会去除字段两侧空白并允许首尾各有一个空字段；中间空字段、非法整数和溢出值返回
// *ElementError。空字符串返回 (nil, nil)。
func (s strto) ParseCommaInts(value string) ([]int, error) {
	if value == "" {
		return nil, nil
	}

	parts := strings.Split(value, ",")
	if len(parts) > 0 && strings.TrimSpace(parts[0]) == "" {
		parts = parts[1:]
	}
	if len(parts) > 0 && strings.TrimSpace(parts[len(parts)-1]) == "" {
		parts = parts[:len(parts)-1]
	}
	if len(parts) == 0 {
		return []int{}, nil
	}

	result := make([]int, len(parts))
	for index, part := range parts {
		rawPart := part
		part = strings.TrimSpace(rawPart)
		if part == "" {
			return nil, &ElementError{Index: index, Value: rawPart, Err: ErrSyntax}
		}
		parsed, err := s.ParseInt(part)
		if err != nil {
			return nil, &ElementError{Index: index, Value: rawPart, Err: err}
		}
		result[index] = parsed
	}
	return result, nil
}
