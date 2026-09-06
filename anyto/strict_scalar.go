package anyto

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

// StrictBool 将任意标量严格转换为 bool。
// 字符串接受不区分大小写的 true、false，以及十进制解析结果为 0 或 1 的无符号整数；
// nil、NaN、无穷大、非法字符串和不支持类型会返回对应的哨兵错误。
func (anyto) StrictBool(value any) (bool, error) {
	reflected, err := strictValue(value)
	if err != nil {
		return false, err
	}
	switch reflected.Kind() {
	case reflect.Bool:
		return reflected.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflected.Int() != 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflected.Uint() != 0, nil
	case reflect.Float32, reflect.Float64:
		floatValue := reflected.Float()
		if math.IsNaN(floatValue) || math.IsInf(floatValue, 0) {
			return false, ErrOutOfRange
		}
		return floatValue != 0, nil
	case reflect.Complex64, reflect.Complex128:
		complexValue := reflected.Complex()
		if math.IsNaN(real(complexValue)) || math.IsNaN(imag(complexValue)) || math.IsInf(real(complexValue), 0) || math.IsInf(imag(complexValue), 0) {
			return false, ErrOutOfRange
		}
		return complexValue != 0, nil
	case reflect.String:
		value := reflected.String()
		switch {
		case strings.EqualFold(value, "true"):
			return true, nil
		case strings.EqualFold(value, "false"):
			return false, nil
		}
		numeric, parseErr := strconv.ParseUint(value, 10, 64)
		if parseErr == nil && numeric <= 1 {
			return numeric == 1, nil
		}
		return false, ErrSyntax
	default:
		return false, ErrUnsupportedType
	}
}

// StrictString 将受支持的标量严格转换为 string。
// 转换格式与 String 一致，但 nil、nil 指针和不支持类型返回 ErrUnsupportedType。
func (f anyto) StrictString(value any) (string, error) {
	reflected, err := strictValue(value)
	if err != nil {
		return "", err
	}
	switch reflected.Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Bool:
		return f.String(reflected.Interface()), nil
	default:
		return "", ErrUnsupportedType
	}
}
