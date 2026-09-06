package anyto

import (
	"reflect"
	"strconv"
)

// String 将任意标量宽松转换为 string，不返回转换错误。
// 支持字符串、整数、浮点数、复数、布尔值以及它们的多级指针；
// nil、nil 指针和不支持类型返回空字符串。需要区分失败和空字符串时请使用 StrictString。
func (anyto) String(value any) string {
	v, ok := indirectValue(value)
	if !ok {
		return ""
	}

	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		return strconv.FormatComplex(v.Complex(), 'g', -1, 64)
	case reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'g', -1, 128)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		return ""
	}
}
