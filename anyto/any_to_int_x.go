package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Int 将给定的值转换为 int。
func (f anyto) Int(i any) (int, error) {
	v, err := f.Int64(i)
	if err != nil {
		return 0, err
	}

	// int 兼容32位和64位系统
	if int64(int(v)) != v {
		return 0, ErrValOut
	}

	return int(v), nil
}

// Int8 将给定的值转换为 int8。
func (f anyto) Int8(i any) (int8, error) {
	value, err := f.Int64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt8 || value > math.MaxInt8 {
		return 0, ErrValOut
	}
	return int8(value), nil
}

// Int16 将给定的值转换为 int16。
func (f anyto) Int16(i any) (int16, error) {
	value, err := f.Int64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt16 || value > math.MaxInt16 {
		return 0, ErrValOut
	}
	return int16(value), nil
}

// Int32 将给定的值转换为 int32。
func (f anyto) Int32(i any) (int32, error) {
	value, err := f.Int64(i)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, ErrValOut
	}
	return int32(value), nil
}

// Int64 将给定的值转换为 int64。
func (anyto) Int64(i any) (int64, error) {
	if i == nil {
		return 0, nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return 0, nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		value := v.Float()
		if math.IsNaN(value) || math.IsInf(value, 0) || value < math.MinInt64 || value >= float64(math.MaxInt64) {
			return 0, ErrValOut
		}
		return int64(value), nil
	case reflect.String:
		intValue, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return intValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value := v.Uint()
		if value > math.MaxInt64 {
			return 0, ErrValOut
		}
		return int64(value), nil
	case reflect.Complex64, reflect.Complex128:
		value := real(v.Complex())
		if math.IsNaN(value) || math.IsInf(value, 0) || value < math.MinInt64 || value >= float64(math.MaxInt64) {
			return 0, ErrValOut
		}
		return int64(value), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrType
	}
}

// AnyToInt 将给定的值转换为 int。
// Deprecated: 使用 AnyTo.Int。
func AnyToInt(i any) (int, error) { return AnyTo.Int(i) }

// AnyToInt8 将给定的值转换为 int8。
// Deprecated: 使用 AnyTo.Int8。
func AnyToInt8(i any) (int8, error) { return AnyTo.Int8(i) }

// AnyToInt16 将给定的值转换为 int16。
// Deprecated: 使用 AnyTo.Int16。
func AnyToInt16(i any) (int16, error) { return AnyTo.Int16(i) }

// AnyToInt32 将给定的值转换为 int32。
// Deprecated: 使用 AnyTo.Int32。
func AnyToInt32(i any) (int32, error) { return AnyTo.Int32(i) }

// AnyToInt64 将给定的值转换为 int64。
// Deprecated: 使用 AnyTo.Int64。
func AnyToInt64(i any) (int64, error) { return AnyTo.Int64(i) }
