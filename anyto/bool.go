package anyto

import "reflect"

// Bool 将任意标量宽松转换为 bool，不返回转换错误。
// 数值零、空字符串、精确小写字符串 "false"、nil、nil 指针和不支持类型返回 false；
// 其他受支持的非零数值或非空字符串返回 true。业务输入校验应使用 StrictBool。
func (anyto) Bool(value any) bool {
	v, ok := indirectValue(value)
	if !ok {
		return false
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
