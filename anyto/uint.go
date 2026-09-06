package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Uint 将任意标量转换为 uint，并按照当前平台位数检查取值范围。
// 输入与转换规则同 Uint64。
func (f anyto) Uint(value any) (uint, error) {
	v, err := f.Uint64(value)
	if err != nil {
		return 0, err
	}

	// 二次转换后再反向比较，可同时覆盖 32 位和 64 位平台。
	if uint64(uint(v)) != v {
		return 0, ErrOutOfRange
	}

	return uint(v), nil
}

// Uint8 将任意标量转换为 uint8，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Uint8(input any) (uint8, error) {
	value, err := f.Uint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint8 {
		return 0, ErrOutOfRange
	}
	return uint8(value), nil
}

// Uint16 将任意标量转换为 uint16，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Uint16(input any) (uint16, error) {
	value, err := f.Uint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint16 {
		return 0, ErrOutOfRange
	}
	return uint16(value), nil
}

// Uint32 将任意标量转换为 uint32，超出取值范围时返回 ErrOutOfRange。
func (f anyto) Uint32(input any) (uint32, error) {
	value, err := f.Uint64(input)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint32 {
		return 0, ErrOutOfRange
	}
	return uint32(value), nil
}

// Uint64 将任意标量转换为 uint64。
// 支持整数、浮点数、复数、布尔值和十进制整数字符串；浮点数截去小数部分，
// 复数只使用实部，布尔值转换为 0 或 1。nil 和 nil 指针返回 (0, nil)。
// 负数返回 ErrNegativeToUnsigned，其他失败按语法、范围或不支持类型分类。
func (anyto) Uint64(value any) (uint64, error) {
	v, ok := indirectValue(value)
	if !ok {
		return 0, nil
	}

	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := v.Int()
		if intValue < 0 {
			return 0, ErrNegativeToUnsigned
		}
		return uint64(intValue), nil
	case reflect.Float32, reflect.Float64:
		floatValue := v.Float()
		if floatValue < 0 {
			return 0, ErrNegativeToUnsigned
		}
		if math.IsNaN(floatValue) || math.IsInf(floatValue, 0) || floatValue >= float64(math.MaxUint64) {
			return 0, ErrOutOfRange
		}
		return uint64(floatValue), nil
	case reflect.Complex64, reflect.Complex128:
		realValue := real(v.Complex())
		if realValue < 0 {
			return 0, ErrNegativeToUnsigned
		}
		if math.IsNaN(realValue) || math.IsInf(realValue, 0) || realValue >= float64(math.MaxUint64) {
			return 0, ErrOutOfRange
		}
		return uint64(realValue), nil
	case reflect.String:
		strValue := v.String()
		uintValue, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			return 0, classifyNumberError(err)
		}
		return uintValue, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrUnsupportedType
	}
}
