package anyto

import "reflect"

// Bool 将给定的值转换为 bool。
func (facade) Bool(i any) bool {
	if i == nil {
		return false
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0
	case reflect.String:
		val := v.String()
		if val == "true" {
			return true
		}
		if val == "false" {
			return false
		}
		return val != ""
	default:
		return false
	}
}

// AnyToBool 将给定的值转换为 bool。
// Deprecated: 使用 AnyTo.Bool。
func AnyToBool(i any) bool { return AnyTo.Bool(i) }
