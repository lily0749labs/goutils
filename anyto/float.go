package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Float32 将任意标量转换为有限的 float32。
// 输入与转换规则同 Float64；超出 float32 取值范围时返回 ErrOutOfRange。
func (f anyto) Float32(value any) (float32, error) {
	f64, err := f.Float64(value)
	if err != nil {
		return 0, err
	}
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 {
		return 0, ErrOutOfRange
	}
	return float32(f64), nil
}

// Float64 将任意标量转换为有限的 float64。
// 支持整数、浮点数、复数、布尔值和浮点数字符串；复数只使用实部，布尔值转换为
// 0 或 1。nil 和 nil 指针返回 (0, nil)，NaN、无穷大和解析溢出返回 ErrOutOfRange。
func (anyto) Float64(input any) (float64, error) {
	v, ok := indirectValue(input)
	if !ok {
		return 0, nil
	}

	var value float64
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		value = v.Float()
	case reflect.String:
		parsed, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, classifyNumberError(err)
		}
		value = parsed
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = float64(v.Uint())
	case reflect.Complex64, reflect.Complex128:
		value = real(v.Complex())
	case reflect.Bool:
		if v.Bool() {
			value = 1
		}
	default:
		return 0, ErrUnsupportedType
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrOutOfRange
	}
	return value, nil
}
