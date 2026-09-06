package strto

import (
	"math"
	"strconv"
)

// ParseFiniteFloat32 将字符串解析为有限 float32；NaN、无穷大、语法错误和溢出均返回错误。
func (strto) ParseFiniteFloat32(value string) (float32, error) {
	parsed, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, ErrNonFiniteFloat
	}
	return float32(parsed), nil
}

// Float32OrZero 将字符串解析为有限 float32；解析失败时返回 0。
func (s strto) Float32OrZero(value string) float32 {
	parsed, _ := s.ParseFiniteFloat32(value)
	return parsed
}

// ParseFiniteFloat64 将字符串解析为有限 float64；NaN、无穷大、语法错误和溢出均返回错误。
func (strto) ParseFiniteFloat64(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, ErrNonFiniteFloat
	}
	return parsed, nil
}

// Float64OrZero 将字符串解析为有限 float64；解析失败时返回 0。
func (s strto) Float64OrZero(value string) float64 {
	parsed, _ := s.ParseFiniteFloat64(value)
	return parsed
}
