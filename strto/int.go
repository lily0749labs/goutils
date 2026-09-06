package strto

import "strconv"

// ParseInt 将十进制字符串解析为 int，失败时返回语法或范围错误。
func (strto) ParseInt(value string) (int, error) {
	parsed, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(parsed), nil
}

// IntOrZero 将十进制字符串解析为 int；解析失败时返回 0。
func (s strto) IntOrZero(value string) int {
	parsed, _ := s.ParseInt(value)
	return parsed
}

// ParseInt8 将十进制字符串解析为 int8，失败时返回语法或范围错误。
func (strto) ParseInt8(value string) (int8, error) {
	parsed, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(parsed), nil
}

// Int8OrZero 将十进制字符串解析为 int8；解析失败时返回 0。
func (s strto) Int8OrZero(value string) int8 {
	parsed, _ := s.ParseInt8(value)
	return parsed
}

// ParseInt16 将十进制字符串解析为 int16，失败时返回语法或范围错误。
func (strto) ParseInt16(value string) (int16, error) {
	parsed, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(parsed), nil
}

// Int16OrZero 将十进制字符串解析为 int16；解析失败时返回 0。
func (s strto) Int16OrZero(value string) int16 {
	parsed, _ := s.ParseInt16(value)
	return parsed
}

// ParseInt32 将十进制字符串解析为 int32，失败时返回语法或范围错误。
func (strto) ParseInt32(value string) (int32, error) {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(parsed), nil
}

// Int32OrZero 将十进制字符串解析为 int32；解析失败时返回 0。
func (s strto) Int32OrZero(value string) int32 {
	parsed, _ := s.ParseInt32(value)
	return parsed
}

// ParseInt64 将十进制字符串解析为 int64，失败时返回语法或范围错误。
func (strto) ParseInt64(value string) (int64, error) {
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsed, nil
}

// Int64OrZero 将十进制字符串解析为 int64；解析失败时返回 0。
func (s strto) Int64OrZero(value string) int64 {
	parsed, _ := s.ParseInt64(value)
	return parsed
}
