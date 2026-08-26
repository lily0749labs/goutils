package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Uint 将给定的值转换为 uint。
func (f anyto) Uint(i any) (uint, error) {
	v, err := f.Uint64(i)
	if err != nil {
		return 0, err
	}

	// uint 兼容32位和64位系统
	if uint64(uint(v)) != v {
		return 0, ErrValOut
	}

	return uint(v), nil
}

// Uint8 将给定的值转换为 uint8。
func (f anyto) Uint8(i any) (uint8, error) {
	value, err := f.Uint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint8 {
		return 0, ErrValOut
	}
	return uint8(value), nil
}

// Uint16 将给定的值转换为 uint16。
func (f anyto) Uint16(i any) (uint16, error) {
	value, err := f.Uint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint16 {
		return 0, ErrValOut
	}
	return uint16(value), nil
}

// Uint32 将给定的值转换为 uint32。
func (f anyto) Uint32(i any) (uint32, error) {
	value, err := f.Uint64(i)
	if err != nil {
		return 0, err
	}
	if value > math.MaxUint32 {
		return 0, ErrValOut
	}
	return uint32(value), nil
}

// Uint64 将给定的值转换为 uint64。
func (anyto) Uint64(i any) (uint64, error) {
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
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := v.Int()
		if intValue < 0 {
			return 0, ErrUnsignedInt
		}
		return uint64(intValue), nil
	case reflect.Float32, reflect.Float64:
		floatValue := v.Float()
		if floatValue < 0 {
			return 0, ErrUnsignedInt
		}
		if math.IsNaN(floatValue) || math.IsInf(floatValue, 0) || floatValue >= float64(math.MaxUint64) {
			return 0, ErrValOut
		}
		return uint64(floatValue), nil
	case reflect.Complex64, reflect.Complex128:
		realValue := real(v.Complex())
		if realValue < 0 {
			return 0, ErrUnsignedInt
		}
		if math.IsNaN(realValue) || math.IsInf(realValue, 0) || realValue >= float64(math.MaxUint64) {
			return 0, ErrValOut
		}
		return uint64(realValue), nil
	case reflect.String:
		strValue := v.String()
		uintValue, err := strconv.ParseUint(strValue, 10, 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return uintValue, nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrType
	}
}

// AnyToUint 将给定的值转换为 uint。
// Deprecated: 使用 AnyTo.Uint。
func AnyToUint(i any) (uint, error) { return AnyTo.Uint(i) }

// AnyToUint8 将给定的值转换为 uint8。
// Deprecated: 使用 AnyTo.Uint8。
func AnyToUint8(i any) (uint8, error) { return AnyTo.Uint8(i) }

// AnyToUint16 将给定的值转换为 uint16。
// Deprecated: 使用 AnyTo.Uint16。
func AnyToUint16(i any) (uint16, error) { return AnyTo.Uint16(i) }

// AnyToUint32 将给定的值转换为 uint32。
// Deprecated: 使用 AnyTo.Uint32。
func AnyToUint32(i any) (uint32, error) { return AnyTo.Uint32(i) }

// AnyToUint64 将给定的值转换为 uint64。
// Deprecated: 使用 AnyTo.Uint64。
func AnyToUint64(i any) (uint64, error) { return AnyTo.Uint64(i) }
