package strto

import (
	"strconv"
)

// IntE 将字符串严格转换为 int，失败时返回解析错误。
func (strto) IntE(v string) (int, error) {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Int 将字符串转换为 int。
// 转换失败时返回零值；需要区分失败和数值零时请使用 IntE。
func (s strto) Int(v string) int {
	i, _ := s.IntE(v)
	return i
}

// Int8E 将字符串严格转换为 int8，失败时返回解析错误。
func (strto) Int8E(v string) (int8, error) {
	i, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

// Int8 将字符串转换为 int8。
// 转换失败时返回零值；需要错误信息时请使用 Int8E。
func (s strto) Int8(v string) int8 {
	i, _ := s.Int8E(v)
	return i
}

// Int16E 将字符串严格转换为 int16，失败时返回解析错误。
func (strto) Int16E(v string) (int16, error) {
	i, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

// Int16 将字符串转换为 int16。
// 转换失败时返回零值；需要错误信息时请使用 Int16E。
func (s strto) Int16(v string) int16 {
	i, _ := s.Int16E(v)
	return i
}

// Int32E 将字符串严格转换为 int32，失败时返回解析错误。
func (strto) Int32E(v string) (int32, error) {
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// Int32 将字符串转换为 int32。
// 转换失败时返回零值；需要错误信息时请使用 Int32E。
func (s strto) Int32(v string) int32 {
	i, _ := s.Int32E(v)
	return i
}

// Int64E 将字符串严格转换为 int64，失败时返回解析错误。
func (strto) Int64E(v string) (int64, error) {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Int64 将字符串转换为 int64。
// 转换失败时返回零值；需要错误信息时请使用 Int64E。
func (s strto) Int64(v string) int64 {
	i, _ := s.Int64E(v)
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
