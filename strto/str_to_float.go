package strto

import (
	"math"
	"strconv"
)

// Float32E 将字符串严格转换为有限 float32，失败时返回解析错误。
func (strto) Float32E(value string) (float32, error) {
	parsed, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, ErrNonFiniteFloat
	}
	return float32(parsed), nil
}

// Float32 将字符串转换为 float32，失败时返回零值。
func (s strto) Float32(value string) float32 {
	parsed, _ := s.Float32E(value)
	return parsed
}

// Float64E 将字符串严格转换为有限 float64，失败时返回解析错误。
func (strto) Float64E(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, ErrNonFiniteFloat
	}
	return parsed, nil
}

// Float64 将字符串转换为 float64，失败时返回零值。
func (s strto) Float64(value string) float64 {
	parsed, _ := s.Float64E(value)
	return parsed
}

// Deprecated: 使用 StrTo.Float32。
func StrToFloat32(value string) float32 { return StrTo.Float32(value) }

// Deprecated: 使用 StrTo.Float64。
func StrToFloat64(value string) float64 { return StrTo.Float64(value) }
