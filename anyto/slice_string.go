package anyto

// sliceToStrings 按输入顺序应用 String 的宽松转换规则。
// 为兼容旧行为，nil 和长度为零的切片都返回 nil。
func sliceToStrings[T any](values []T, convert func(any) string) []string {
	if len(values) == 0 {
		return nil
	}

	result := make([]string, len(values))
	for index, value := range values {
		result[index] = convert(value)
	}
	return result
}

// ValuesToStrings 将任意值切片转换为字符串切片。
// 元素顺序保持不变，不支持的元素按 String 的宽松规则转换为空字符串。
func (f anyto) ValuesToStrings(values []any) []string {
	return sliceToStrings(values, f.String)
}

// IntsToStrings 按顺序将 int 切片转换为字符串切片。
func (f anyto) IntsToStrings(values []int) []string {
	return sliceToStrings(values, f.String)
}

// Int8sToStrings 按顺序将 int8 切片转换为字符串切片。
func (f anyto) Int8sToStrings(values []int8) []string {
	return sliceToStrings(values, f.String)
}

// Int16sToStrings 按顺序将 int16 切片转换为字符串切片。
func (f anyto) Int16sToStrings(values []int16) []string {
	return sliceToStrings(values, f.String)
}

// Int32sToStrings 按顺序将 int32 切片转换为字符串切片。
func (f anyto) Int32sToStrings(values []int32) []string {
	return sliceToStrings(values, f.String)
}

// Int64sToStrings 按顺序将 int64 切片转换为字符串切片。
func (f anyto) Int64sToStrings(values []int64) []string {
	return sliceToStrings(values, f.String)
}

// UintsToStrings 按顺序将 uint 切片转换为字符串切片。
func (f anyto) UintsToStrings(values []uint) []string {
	return sliceToStrings(values, f.String)
}

// Uint8sToStrings 按顺序将 uint8 切片转换为字符串切片。
func (f anyto) Uint8sToStrings(values []uint8) []string {
	return sliceToStrings(values, f.String)
}

// Uint16sToStrings 按顺序将 uint16 切片转换为字符串切片。
func (f anyto) Uint16sToStrings(values []uint16) []string {
	return sliceToStrings(values, f.String)
}

// Uint32sToStrings 按顺序将 uint32 切片转换为字符串切片。
func (f anyto) Uint32sToStrings(values []uint32) []string {
	return sliceToStrings(values, f.String)
}

// Uint64sToStrings 按顺序将 uint64 切片转换为字符串切片。
func (f anyto) Uint64sToStrings(values []uint64) []string {
	return sliceToStrings(values, f.String)
}
