package anyto

import "reflect"

// indirectValue 持续解开指针和接口，返回最终承载的值。
// nil、nil 指针及循环引用均返回 false，避免调用方重复处理和无限解引用。
func indirectValue(value any) (reflect.Value, bool) {
	if value == nil {
		return reflect.Value{}, false
	}

	reflected := reflect.ValueOf(value)
	type pointerIdentity struct {
		typeOf  reflect.Type
		address uintptr
	}
	var visited map[pointerIdentity]struct{}

	for reflected.IsValid() && (reflected.Kind() == reflect.Ptr || reflected.Kind() == reflect.Interface) {
		if reflected.IsNil() {
			return reflect.Value{}, false
		}
		if reflected.Kind() == reflect.Ptr {
			if visited == nil {
				visited = make(map[pointerIdentity]struct{})
			}
			identity := pointerIdentity{typeOf: reflected.Type(), address: reflected.Pointer()}
			if _, exists := visited[identity]; exists {
				return reflect.Value{}, false
			}
			visited[identity] = struct{}{}
		}
		reflected = reflected.Elem()
	}
	return reflected, reflected.IsValid()
}

// strictValue 将 indirectValue 的无值结果转换为严格模式使用的类型错误。
func strictValue(value any) (reflect.Value, error) {
	reflected, ok := indirectValue(value)
	if !ok {
		return reflect.Value{}, ErrUnsupportedType
	}
	return reflected, nil
}
