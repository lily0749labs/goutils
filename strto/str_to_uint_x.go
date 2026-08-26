package strto

import "strconv"

// Uint 将字符串转换为 uint。
func (facade) Uint(v string) uint {
	i, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0
	}
	return uint(i)
}

// Uint8 将字符串转换为 uint8。
func (facade) Uint8(v string) uint8 {
	i, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(i)
}

// Uint16 将字符串转换为 uint16。
func (facade) Uint16(v string) uint16 {
	i, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(i)
}

// Uint32 将字符串转换为 uint32。
func (facade) Uint32(v string) uint32 {
	i, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(i)
}

// Uint64 将字符串转换为 uint64。
func (facade) Uint64(v string) uint64 {
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0
	}
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
