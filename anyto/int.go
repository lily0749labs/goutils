package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Int 将任意标量转换为 int，并按照当前平台位数检查取值范围。
// 输入与转换规则同 Int64。
func (f anyto) Int(value any) (int, error) {
	v, err := f.Int64(value)
	if err != nil {
		return 0, err
	}

	// 二次转换后再反向比较，可同时覆盖 32 位和 64 位平台。
	if int64(int(v)) != v {
		return 0, ErrOutOfRange
	}

	return int(v), nil
}

// Int8 将任意标量转换为 int8，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Int8(input any) (int8, error) {
	value, err := f.Int64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, ErrOutOfRange
	}
	return int8(value), nil
}

// Int16 将任意标量转换为 int16，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Int16(input any) (int16, error) {
	value, err := f.Int64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, ErrOutOfRange
	}
	return int16(value), nil
}

// Int32 将任意标量转换为 int32，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Int32(input any) (int32, error) {
	value, err := f.Int64(input)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, ErrOutOfRange
	}
	return int32(value), nil
}

// Int64 将任意标量转换为 int64。
// 支持整数、浮点数、复数、布尔值和十进制整数字符串；浮点数截去小数部分，
// 复数只使用实部，布尔值转换为 0 或 1。nil 和 nil 指针返回 (0, nil)。
// 语法错误、不支持类型和范围错误分别返回对应的哨兵错误。
func (anyto) Int64(value any) (int64, error) {
	v, ok := indirectValue(value)
	if !ok {
		return 0, nil
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		value := v.Float()
		if math.IsNaN(value) || math.IsInf(value, 0) || value < math.MinInt64 || value >= float64(math.MaxInt64) {
			return 0, ErrOutOfRange
		}
		return int64(value), nil
	case reflect.String:
		intValue, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return 0, classifyNumberError(err)
		}
		return intValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value := v.Uint()
		if value > math.MaxInt64 {
			return 0, ErrOutOfRange
		}
		return int64(value), nil
	case reflect.Complex64, reflect.Complex128:
		value := real(v.Complex())
		if math.IsNaN(value) || math.IsInf(value, 0) || value < math.MinInt64 || value >= float64(math.MaxInt64) {
			return 0, ErrOutOfRange
		}
		return int64(value), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrUnsupportedType
	}
}
