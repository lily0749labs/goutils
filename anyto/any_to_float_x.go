package anyto

import (
	"math"
	"reflect"
	"strconv"
)

// Float32 将给定的值转换为 float32。
func (f anyto) Float32(i any) (float32, error) {
	f64, err := f.Float64(i)
	if err != nil {
		return 0, err
	}
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 {
		return 0, ErrValOut
	}
	return float32(f64), nil
}

// Float64 将给定的值转换为 float64。
func (anyto) Float64(i any) (float64, error) {
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
		return v.Float(), nil
	case reflect.String:
		floatValue, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return 0, ErrSyntax
		}
		return floatValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint()), nil
	case reflect.Complex64, reflect.Complex128:
		return real(v.Complex()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, ErrType
	}
}

// AnyToFloat32 将给定的值转换为 float32。
// Deprecated: 使用 AnyTo.Float32。
func AnyToFloat32(i any) (float32, error) { return AnyTo.Float32(i) }

// AnyToFloat64 将给定的值转换为 float64。
// Deprecated: 使用 AnyTo.Float64。
func AnyToFloat64(i any) (float64, error) { return AnyTo.Float64(i) }
