package strto

import (
	"strconv"
)

// Int 将字符串转换为 int。
func (facade) Int(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

// Int8 将字符串转换为 int8。
func (facade) Int8(v string) int8 {
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// Int16 将字符串转换为 int16。
func (facade) Int16(v string) int16 {
	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// Int32 将字符串转换为 int32。
func (facade) Int32(v string) int32 {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// Int64 将字符串转换为 int64。
func (facade) Int64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// StrToInt 将字符串转换为 int。
// Deprecated: 使用 StrTo.Int。
func StrToInt(v string) int { return StrTo.Int(v) }

// StrToInt8 将字符串转换为 int8。
// Deprecated: 使用 StrTo.Int8。
func StrToInt8(v string) int8 { return StrTo.Int8(v) }

// StrToInt16 将字符串转换为 int16。
// Deprecated: 使用 StrTo.Int16。
func StrToInt16(v string) int16 { return StrTo.Int16(v) }

// StrToInt32 将字符串转换为 int32。
// Deprecated: 使用 StrTo.Int32。
func StrToInt32(v string) int32 { return StrTo.Int32(v) }

// StrToInt64 将字符串转换为 int64。
// Deprecated: 使用 StrTo.Int64。
func StrToInt64(v string) int64 { return StrTo.Int64(v) }
