package strto

import "strconv"

// UintE 将字符串严格转换为 uint，失败时返回解析错误。
func (strto) UintE(v string) (uint, error) {
	i, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

// Uint 将字符串转换为 uint。
// 转换失败时返回零值；需要错误信息时请使用 UintE。
func (s strto) Uint(v string) uint {
	i, _ := s.UintE(v)
	return i
}

// Uint8E 将字符串严格转换为 uint8，失败时返回解析错误。
func (strto) Uint8E(v string) (uint8, error) {
	i, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

// Uint8 将字符串转换为 uint8。
// 转换失败时返回零值；需要错误信息时请使用 Uint8E。
func (s strto) Uint8(v string) uint8 {
	i, _ := s.Uint8E(v)
	return i
}

// Uint16E 将字符串严格转换为 uint16，失败时返回解析错误。
func (strto) Uint16E(v string) (uint16, error) {
	i, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(i), nil
}

// Uint16 将字符串转换为 uint16。
// 转换失败时返回零值；需要错误信息时请使用 Uint16E。
func (s strto) Uint16(v string) uint16 {
	i, _ := s.Uint16E(v)
	return i
}

// Uint32E 将字符串严格转换为 uint32，失败时返回解析错误。
func (strto) Uint32E(v string) (uint32, error) {
	i, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

// Uint32 将字符串转换为 uint32。
// 转换失败时返回零值；需要错误信息时请使用 Uint32E。
func (s strto) Uint32(v string) uint32 {
	i, _ := s.Uint32E(v)
	return i
}

// Uint64E 将字符串严格转换为 uint64，失败时返回解析错误。
func (strto) Uint64E(v string) (uint64, error) {
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Uint64 将字符串转换为 uint64。
// 转换失败时返回零值；需要错误信息时请使用 Uint64E。
func (s strto) Uint64(v string) uint64 {
	i, _ := s.Uint64E(v)
	return i
}

// StrToUint 将字符串转换为 uint。
// Deprecated: 使用 StrTo.Uint。
func StrToUint(v string) uint { return StrTo.Uint(v) }

// StrToUint8 将字符串转换为 uint8。
// Deprecated: 使用 StrTo.Uint8。
func StrToUint8(v string) uint8 { return StrTo.Uint8(v) }

// StrToUint16 将字符串转换为 uint16。
// Deprecated: 使用 StrTo.Uint16。
func StrToUint16(v string) uint16 { return StrTo.Uint16(v) }

// StrToUint32 将字符串转换为 uint32。
// Deprecated: 使用 StrTo.Uint32。
func StrToUint32(v string) uint32 { return StrTo.Uint32(v) }

// StrToUint64 将字符串转换为 uint64。
// Deprecated: 使用 StrTo.Uint64。
func StrToUint64(v string) uint64 { return StrTo.Uint64(v) }
