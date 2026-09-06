package strto

import "strconv"

// ParseUint 将十进制字符串解析为 uint，失败时返回语法或范围错误。
func (strto) ParseUint(value string) (uint, error) {
	parsed, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}

// UintOrZero 将十进制字符串解析为 uint；解析失败时返回 0。
func (s strto) UintOrZero(value string) uint {
	parsed, _ := s.ParseUint(value)
	return parsed
}

// ParseUint8 将十进制字符串解析为 uint8，失败时返回语法或范围错误。
func (strto) ParseUint8(value string) (uint8, error) {
	parsed, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(parsed), nil
}

// Uint8OrZero 将十进制字符串解析为 uint8；解析失败时返回 0。
func (s strto) Uint8OrZero(value string) uint8 {
	parsed, _ := s.ParseUint8(value)
	return parsed
}

// ParseUint16 将十进制字符串解析为 uint16，失败时返回语法或范围错误。
func (strto) ParseUint16(value string) (uint16, error) {
	parsed, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(parsed), nil
}

// Uint16OrZero 将十进制字符串解析为 uint16；解析失败时返回 0。
func (s strto) Uint16OrZero(value string) uint16 {
	parsed, _ := s.ParseUint16(value)
	return parsed
}

// ParseUint32 将十进制字符串解析为 uint32，失败时返回语法或范围错误。
func (strto) ParseUint32(value string) (uint32, error) {
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(parsed), nil
}

// Uint32OrZero 将十进制字符串解析为 uint32；解析失败时返回 0。
func (s strto) Uint32OrZero(value string) uint32 {
	parsed, _ := s.ParseUint32(value)
	return parsed
}

// ParseUint64 将十进制字符串解析为 uint64，失败时返回语法或范围错误。
func (strto) ParseUint64(value string) (uint64, error) {
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}
	return parsed, nil
}

// Uint64OrZero 将十进制字符串解析为 uint64；解析失败时返回 0。
func (s strto) Uint64OrZero(value string) uint64 {
	parsed, _ := s.ParseUint64(value)
	return parsed
}
